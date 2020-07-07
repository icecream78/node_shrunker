package cmd

import (
	"fmt"
	"os"

	"github.com/icecream78/node_shrinker/shrunk"

	"github.com/spf13/cobra"
)

var checkPath string
var verbose bool
var excludeNames []string
var includeNames []string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "node_shrinker",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: move Shrinker configuring with builder
		err := shrunk.NewShrinker(&shrunk.Config{
			CheckPath:       checkPath,
			RemoveDirNames:  []string{},
			RemoveFileNames: []string{},
			VerboseOutput:   verbose,
			ExcludeNames:    excludeNames,
			IncludeNames:    includeNames,
		}).Start()
		if err != nil {
			fmt.Printf("Someghing broken=) %v\n", err)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&checkPath, "dir", "d", "", "path to directory where need cleanup")
	rootCmd.PersistentFlags().StringSliceVarP(&excludeNames, "exclude", "e", []string{}, "List of files/directories that should not be removed. Flag can be specified multiple times")
	rootCmd.PersistentFlags().StringSliceVarP(&includeNames, "include", "i", []string{}, "List of files/directories that should be included in remove list. Flag can be specified multiple times")

	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "more detailed output")
}
