package style

import (
	"charm.land/lipgloss/v2"
	"charm.land/tone"
)

func Error(err error) string {
	return lipgloss.NewStyle().
		Foreground(tone.Hazy).
		Background(tone.Pom).Render(err.Error())
}
