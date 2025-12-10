package cmd

import (
    "github.com/spf13/cobra"
    "github.com/jdraines/mapcli-go/internal/mapcli"
)


var makeCopy bool = false

var createCmd = &cobra.Command{
		Use: "create <command> <mappingPath>",
		Short: "Create a cli mapping",
		Args: cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
        	command := args[0]
			mappingPath := args[1]
			err := mapcli.CreateMappedCli(mappingPath, command, makeCopy)
			cobra.CheckErr(err)
		},
	}

func init() {
	createCmd.Flags().BoolVar(&makeCopy, "copy", false, 
		"Whether to create a copy of the mapping rather than a symlink to the file you provide")
    rootCmd.AddCommand(createCmd)
}
