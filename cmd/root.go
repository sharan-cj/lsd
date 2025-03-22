package cmd

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/sharan-cj/lsd/packages/styles"
	"github.com/sharan-cj/lsd/packages/utils"

	"github.com/charmbracelet/lipgloss/tree"
	"github.com/spf13/cobra"
)

// flags
var depth uint8
var verbose bool
var maxDepth bool
var all bool

var rootCmd = &cobra.Command{
	Use:     "lsd",
	Short:   "lsd is a CLI utility for visualizing directory structures in a tree format, enriched with metadata for files and folders.",
	Long:    "lsd is a command-line utility designed for developers and system administrators to visualize directory structures in a tree-like format.",
	Run:     execFunc,
	Version: "0.1.0",
	Args:    cobra.MaximumNArgs(1),
}

func init() {
	rootCmd.Flags().Uint8VarP(&depth, "depth", "d", 2, "Depth of the tree.")
	rootCmd.Flags().BoolVarP(&maxDepth, "max", "M", false, "Print the max depth of the tree. Default is false.")
	rootCmd.Flags().BoolVar(&verbose, "verbose", false, "Verbose mode. Default is false.")
	rootCmd.Flags().BoolVarP(&all, "all", "a", false, "views all files and directories, including hidden ones.")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func execFunc(cmd *cobra.Command, args []string) {
	dir := "."

	var dirName string

	if len(args) > 0 {
		dir = args[0]
		dirName = dir
	} else {
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
			return
		}
		dirName = wd
	}

	t := tree.New().Root(styles.DirStyle.Render(dirName))

	if maxDepth {
		depth = math.MaxUint8
	}

	start := time.Now()
	utils.BuildTree(t, dir, depth, verbose, all)
	fmt.Println("")
	fmt.Println(t)

	elapsed := time.Since(start)
	fmt.Printf("Time taken: %s\n", elapsed)
}
