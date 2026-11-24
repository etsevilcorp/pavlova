package style

import (
	"charm.land/lipgloss/v2"
	"charm.land/tone"
)

func String(s string) string {
	return lipgloss.NewStyle().
		Foreground(tone.Citron).
		Background(tone.Dolly).Render(s)
}
