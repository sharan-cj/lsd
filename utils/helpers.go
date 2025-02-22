package utils

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/tree"
	"github.com/sharan-cj/lsd/utils/colors"
)

var DirStyle = lipgloss.NewStyle().Background(lipgloss.Color(colors.Blue)).Foreground(lipgloss.Color(colors.White)).Padding(0, 2).MarginBottom(1).Bold(true)

var TimestampStyle = lipgloss.NewStyle().Background(lipgloss.Color(colors.Orange)).Foreground(lipgloss.Color(colors.White)).Padding(0, 1)

var FileSizeStyle = lipgloss.NewStyle().Background(lipgloss.Color(colors.Green)).Foreground(lipgloss.Color(colors.White)).Padding(0, 1)

var FileNameStyle = lipgloss.NewStyle().Background(lipgloss.Color(colors.White)).Foreground(lipgloss.Color(colors.Black)).MarginBottom(1).Padding(0, 2)

var FileCountStyle = lipgloss.NewStyle().Background(lipgloss.Color(colors.White)).Foreground(lipgloss.Color(colors.Black)).MarginBottom(1).Padding(0, 2)

func BuildTree(t *tree.Tree, path string, depth uint8, verbose bool) {

	if depth == 0 {
		return
	}

	entries, err := os.ReadDir(path)

	if verbose {
		println("Reading dir: ", path, " | Found: ", len(entries), "entries")
	}

	if err != nil {
		println("Error: ", err.Error())
		return
	}

	for _, entry := range entries {

		if entry.IsDir() {

			newEntry := tree.New().Root(DirStyle.Render(entry.Name()))
			t.Child(newEntry)
			BuildTree(newEntry, filepath.Join(path, entry.Name()), depth-1, verbose)
		} else {

			fileInfo, err := entry.Info()

			if err != nil {
				println("Error: ", err.Error())
				return
			}

			timestamp := TimestampStyle.Render(fileInfo.ModTime().Format("2006-01-02 03:04:05 PM"))
			fileSize := FileSizeStyle.Render(FormatFileSize(fileInfo.Size()))
			fileName := FileNameStyle.Render(entry.Name())
			str := lipgloss.JoinHorizontal(lipgloss.Top, fileName, timestamp, fileSize)

			t.Child(str)

		}
	}
}
