package cmd

import (
	"fmt"
	"math"
	"os"

	"github.com/sharan-cj/lsd/utils"

	"github.com/charmbracelet/lipgloss/tree"
	"github.com/spf13/cobra"
)

// flags
var depth uint8
var verbose bool
var maxDepth bool

var rootCmd = &cobra.Command{
	Use:     "lsd",
	Short:   "lsd is a CLI utility for visualizing directory structures in a tree format, enriched with metadata for files and folders.",
	Long:    "lsd is a command-line utility designed for developers and system administrators to visualize directory structures in a tree-like format.",
	Run:     execFunc,
	Version: "0.1.0",
	Args:    cobra.MaximumNArgs(1),
}

func init() {
	rootCmd.Flags().Uint8VarP(&depth, "depth", "d", 2, "Depth of the tree. Default is 2.")
	rootCmd.Flags().BoolVarP(&maxDepth, "max", "M", false, "Print the max depth of the tree. Default is false.")
	rootCmd.Flags().BoolVar(&verbose, "verbose", false, "Verbose mode. Default is false.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func execFunc(cmd *cobra.Command, args []string) {
	dir := "."

	if len(args) > 0 {
		dir = args[0]
	}

	t := tree.New().Root(dir)

	if maxDepth {
		depth = math.MaxUint8
	}

	utils.BuildTree(t, dir, depth, verbose)
	fmt.Println(t)
}
