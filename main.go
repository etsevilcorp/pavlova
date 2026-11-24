package main

import (
	"bufio"
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
		lipgloss.Print(starter.String())
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			parts := strings.Fields(scanner.Text())
			cmd := exec.Command(parts[0], parts[1:]...)
			out, err := cmd.Output()
			lipgloss.Fprint(os.Stdout, style.String(string(out)))
			lipgloss.Fprint(os.Stderr, style.Error(err))
		}
	}

	os.Exit(0)
}
