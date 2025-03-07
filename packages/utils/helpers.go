package utils

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/tree"
	"github.com/enescakir/emoji"
	"github.com/sharan-cj/lsd/packages/styles"
)

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

	filesCount, dirsCount := 0, 0
	for index, entry := range entries {

		if entry.IsDir() {
			dirsCount++
			newEntry := tree.New().Root(styles.DirStyle.Render(entry.Name()))
			t.Child(newEntry)
			BuildTree(newEntry, filepath.Join(path, entry.Name()), depth-1, verbose)
		} else {
			filesCount++
			fileInfo, err := entry.Info()

			if err != nil {
				println("Error: ", err.Error())
				return
			}

			timestamp := styles.TimestampStyle.Render(fileInfo.ModTime().Format("2006-01-02 03:04:05 PM"))
			fileSize := styles.FileSizeStyle.Render(FormatFileSize(fileInfo.Size()))
			fileName := styles.FileNameStyle.Render(entry.Name())
			str := lipgloss.JoinHorizontal(lipgloss.Top, fileName, timestamp, fileSize)

			t.Child(str)

		}

		if index == len(entries)-1 {
			rootDirName := t.Value()
			t.Root(lipgloss.JoinHorizontal(lipgloss.Top, rootDirName, styles.DirCountStyle.Render(string(emoji.PageFacingUp), strconv.Itoa(filesCount)), styles.FileCountStyle.Render(string(emoji.OpenFileFolder), strconv.Itoa(dirsCount))))
		}
	}
}
