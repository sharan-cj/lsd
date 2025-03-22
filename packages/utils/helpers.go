package utils

import (
	"os"
	"path/filepath"
	"strconv"
	"sync"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/tree"
	"github.com/enescakir/emoji"
	"github.com/sharan-cj/lsd/packages/styles"
)

func BuildTree(t *tree.Tree, path string, depth uint8, verbose bool, all bool) {

	if depth == 0 {
		return
	}

	entries, err := os.ReadDir(path)
	if err != nil {
		println("Error: ", err.Error())
		return
	}

	if verbose {
		println("Reading dir: ", path, " | Found: ", len(entries), "entries")
	}

	var filteredEntries []os.DirEntry
	if !all {
		filteredEntries = make([]os.DirEntry, 0, len(entries))
		for _, entry := range entries {
			if entry.Name()[0] != '.' {
				filteredEntries = append(filteredEntries, entry)
			}
		}
	} else {
		filteredEntries = entries
	}

	filesCount, dirsCount := 0, 0
	var wg sync.WaitGroup

	children := make([]*tree.Tree, 0, len(filteredEntries))

	for _, entry := range filteredEntries {
		if entry.IsDir() {
			dirsCount++
			dirName := entry.Name()
			dirPath := filepath.Join(path, dirName)
			newEntry := tree.New().Root(styles.DirStyle.Render(dirName))
			children = append(children, newEntry)

			// Only use goroutines for deeper levels to avoid overhead for small directories
			if depth > 1 {
				wg.Add(1)
				go func(subTree *tree.Tree, subPath string, remainingDepth uint8) {
					defer wg.Done()
					BuildTree(subTree, subPath, remainingDepth, verbose, all)
				}(newEntry, dirPath, depth-1)
			} else {
				BuildTree(newEntry, dirPath, depth-1, verbose, all)
			}
		} else {
			filesCount++
			children = append(children, createFileNode(entry))
		}
	}

	// Wait for all directory processing goroutines to complete
	wg.Wait()

	// Add all children to the tree
	for _, child := range children {
		t.Child(child)
	}

	// Update the root node with counts
	updateRootWithCounts(t, filesCount, dirsCount)
}

// Helper function to create a file node
func createFileNode(entry os.DirEntry) *tree.Tree {
	fileInfo, err := entry.Info()
	if err != nil {
		return tree.New().Root(styles.FileNameStyle.Render(entry.Name() + " [error]"))
	}

	timestamp := styles.TimestampStyle.Render(fileInfo.ModTime().Format("2006-01-02 03:04:05 PM"))
	fileSize := styles.FileSizeStyle.Render(FormatFileSize(fileInfo.Size()))
	fileName := styles.FileNameStyle.Render(entry.Name())

	return tree.New().Root(lipgloss.JoinHorizontal(lipgloss.Top, fileName, timestamp, fileSize))
}

// Helper function to update root node with counts
func updateRootWithCounts(t *tree.Tree, filesCount, dirsCount int) {
	rootDirName := t.Value()
	fileCountDisplay := styles.FileCountStyle.Render(string(emoji.OpenFileFolder), strconv.Itoa(dirsCount))
	dirCountDisplay := styles.DirCountStyle.Render(string(emoji.PageFacingUp), strconv.Itoa(filesCount))

	t.Root(lipgloss.JoinHorizontal(lipgloss.Top, rootDirName, dirCountDisplay, fileCountDisplay))
}
