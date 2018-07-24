package cli

import (
	"strings"

	"github.com/j-frost/isac-cli-prototype/pkg/add"
	"github.com/j-frost/isac-cli-prototype/pkg/rm"
	"github.com/spf13/cobra"
)

func Execute() {
	cmd := &cobra.Command{
		Use:     "isac-cli [draft] [options...] VERB SUBJECT OPERAND",
		Version: "0.1.0",
		Short:   "isac-cli is a prototype cli for our ISAC repository",
		Long: `isac-cli is a prototype cli for our ISAC repository at 
//SEIBERT/MEDIA. It tries to unify a number of other tools that also 
manipulate ISAC. This is an experiment in usability. `,
	}

	cmd.PersistentFlags().BoolP("verbose", "v", false, "verbose mode with additional debugging information")
	cmd.PersistentFlags().BoolP("no-act", "n", false, "dry run mode; no actual changes will be made")

	customUsage := strings.Replace(cmd.UsageTemplate(), `{{if .Runnable}}
  {{.UseLine}}{{end}}{{if .HasAvailableSubCommands}}
  {{.CommandPath}} [command]{{end}}`, `
  {{.UseLine}}`, -1)
	customUsage = strings.Replace(customUsage, "Available Commands", "Available Verbs", -1)
	cmd.SetUsageTemplate(customUsage)

	cmd.AddCommand(add.NewAddCmd(), rm.NewRmCmd())

	for _, subCmd := range cmd.Commands() {
		cmd.ValidArgs = append(cmd.ValidArgs, subCmd.Name())
	}
	cmd.Args = cobra.OnlyValidArgs

	cmd.Execute()
}
