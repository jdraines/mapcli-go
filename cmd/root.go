package cmd

import (
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use: "mapcli",
    Short: "A mapper that lets you reword CLI tools",
    Long: `Use this tool to remap any key words used in any CLI tool by creating a yaml string:string mapping to be used for word replacement. Then run this command to generate an executable that can be used to execute your mapped function.`,
    Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
}

func Execute() {
    cobra.CheckErr(rootCmd.Execute())
}


