package cmd

import (
    "fmt"
	"log"
	"os"
	"os/exec"
    "github.com/spf13/cobra"
    "github.com/jdraines/mapcli-go/internal/mapcli"
)


var runCmd = &cobra.Command{
    Use: "run <command> [additional args]",
    Short: "Run a mapped cli",
	Args: cobra.MinimumNArgs(1),
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		commandName := args[0]
		m, err := mapcli.ReadMapping(commandName)
		if err != nil { cobra.CheckErr(err) }
		newArgs := mapcli.MapArgs(m, args)
		execute(newArgs)
    },
}


func init() {
	rootCmd.AddCommand(runCmd)
	runCmd.PersistentFlags().BoolP("help", "h", false, "Print usage")
	runCmd.PersistentFlags().Lookup("help").Hidden = true
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


func execute(args []string) {
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
