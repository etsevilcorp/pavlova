package style

import (
	"charm.land/lipgloss/v2"
	"charm.land/tone"
)

func Name(s string) string {
	return lipgloss.NewStyle().
		Foreground(tone.Zest).
		Background(tone.Blush).Faint(false).Render(s)
}
