package utils

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/tree"
)

var style = lipgloss.NewStyle().Border(lipgloss.RoundedBorder())

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
			newEntry := tree.New().Root(style.Render(entry.Name()))
			t.Child(newEntry)
			BuildTree(newEntry, filepath.Join(path, entry.Name()), depth-1, verbose)
		} else {
			t.Child(entry.Name())
		}
	}
}
