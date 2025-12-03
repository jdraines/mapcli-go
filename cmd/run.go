package cmd

import (
    "fmt"
    "github.com/spf13/cobra"
    "github.com/jdraines/mapcli-go/internal/mapcli"
)


var runCmd = &cobra.Command{
    Use: "run <command> [additional args]",
    Short: "Run a mapped cli",
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		commandName := args[0]
		m, err := mapcli.ReadMapping(commandName)
		if err != nil { panic(err) }
    	// var m = mapcli.ReadMap("/home/john/src/okprograms/mapcli-go/examples/example2.yaml")
		mapArgs(args, m)
    },
}


func init() {
	rootCmd.AddCommand(runCmd)
}


func mapArgs(args []string, mapping map[string]string) {
    for k, v := range mapping {
        fmt.Printf("%s = %s\n", k, v)
    }
    var newArgs = mapcli.MapArgs(mapping, args)
    for _, a := range newArgs {
        fmt.Printf("%s ", a)
    }
    fmt.Print("\n")
}
