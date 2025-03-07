package styles

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/sharan-cj/lsd/packages/colors"
)

var DirStyle = lipgloss.NewStyle().Background(lipgloss.Color(colors.Blue)).Foreground(lipgloss.Color(colors.White)).Padding(0, 2).MarginBottom(1).Bold(true)

var TimestampStyle = lipgloss.NewStyle().Background(lipgloss.Color(colors.Brown)).Foreground(lipgloss.Color(colors.White)).Padding(0, 1)

var FileSizeStyle = lipgloss.NewStyle().Background(lipgloss.Color(colors.Green)).Foreground(lipgloss.Color(colors.White)).Padding(0, 1)

var FileNameStyle = lipgloss.NewStyle().Background(lipgloss.Color(colors.White)).Foreground(lipgloss.Color(colors.Black)).MarginBottom(1).Padding(0, 2)

var FileCountStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(colors.Blue)).MarginBottom(1).Padding(0, 1).Bold(true)

var DirCountStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(colors.Green)).MarginBottom(1).Padding(0, 1).Bold(true)
