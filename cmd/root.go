package cmd

import (
	"fmt"
	"os"

	"github.com/sharan-cj/lsd/utils"

	"github.com/charmbracelet/lipgloss/tree"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "lsd",
	Short: "lsd is a CLI utility for visualizing directory structures in a tree format, enriched with metadata for files and folders.",
	Long:  "lsd is a command-line utility designed for developers and system administrators to visualize directory structures in a tree-like format.",
	Run:   execFunc,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func execFunc(cmd *cobra.Command, args []string) {
	path := "."

	if len(args) > 0 {
		path = args[0]
	}

	t := tree.New().Root(path)
	utils.BuildTree(t, path)
	fmt.Println(t)
}
