package style

import (
	"charm.land/lipgloss/v2"
	"charm.land/tone"
)

func Dir(s string) string {
	return lipgloss.NewStyle().
		Foreground(tone.Lichen).
		Background(tone.Mauve).Faint(false).Render(s)
}
