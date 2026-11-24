package main

import (
	"bufio"
	"io"
	"os"
	"os/exec"
	"os/user"
	"strings"

	"charm.land/lipgloss/v2"
	"charm.land/tone"
	"git.etsevilcorp.com/pavlova/style"
)

func main() {
	sub := os.Args[1:][0]

	starter := strings.Builder{}

	u, err := user.Current()
	if err != nil {
		lipgloss.Printf("pavlova: unable to get current user: %v", style.Error(err))
		os.Exit(1)
	}
	if u != nil {
		starter.WriteString(style.Name(u.Name))
		starter.WriteString(lipgloss.NewStyle().Background(tone.Lilac).Foreground(tone.Bok).Render("@"))
	}

	host, err := os.Hostname()
	if err != nil {
		lipgloss.Printf("pavlova: couldn't get hostname: %v", err.Error())
		os.Exit(1)
	}
	starter.WriteString(style.Host(host))

	wd, err := os.Getwd()
	if err != nil {
		lipgloss.Printf("pavlova: can't get current working directory: %v", err.Error())
		os.Exit(1)
	}
	starter.WriteString(style.Dir(wd))

	if sub == "prompt" {
		lipgloss.Print(starter.String(), "> ")
		scanner := bufio.NewScanner(os.Stdin)
		if !scanner.Scan() {
			return
		}

		parts := strings.Fields(scanner.Text())
		if len(parts) == 0 {
			os.Exit(0)
		}

		if parts[0] == "exit" {
			return
		}
		if parts[0] == "cd" {
			target := "."
			if len(parts) > 1 {
				target = parts[1]
			}
			if err := os.Chdir(target); err != nil {
				lipgloss.Fprintln(os.Stderr, style.Error(err))
			}
			return
		}

		cmd := exec.Command(parts[0], parts[1:]...)

		out, err := cmd.StdoutPipe()
		if err != nil {
			lipgloss.Fprint(os.Stderr, style.Error(err))
			os.Exit(127)
		}

		ebuf, err := cmd.StderrPipe()
		if err != nil {
			lipgloss.Fprint(os.Stderr, style.Error(err))
			os.Exit(127)
		}

		if err := cmd.Start(); err != nil {
			lipgloss.Fprint(os.Stderr, style.Error(err))
			os.Exit(127)
		}

		go io.Copy(os.Stderr, ebuf)
		io.Copy(os.Stdout, out)

		if err := cmd.Wait(); err != nil {
			lipgloss.Fprint(os.Stderr, style.Error(err))
			os.Exit(1)
		}

	}

	os.Exit(0)
}
