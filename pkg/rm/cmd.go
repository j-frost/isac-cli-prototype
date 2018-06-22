package rm

import (
	"github.com/spf13/cobra"
	"strings"
)

func NewRmCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove TYPE SUBJECT",
		Short: "Remove an entity from isac",
		Long: `Remove a TYPE type entity SUBJECT from isac, where TYPE may be one of 
application, node, virtual-machine, ip, or test-system. The value of SUBJECT 
varies with TYPE. `,
		Aliases: []string{"r", "rm", "delete", "del", "d"},
	}

	cmd.SetUsageFunc(func(command *cobra.Command) error {
		if !command.HasParent() {
			return command.UsageFunc()(command)
		}
		command.SetUsageTemplate(strings.Replace(command.Parent().UsageTemplate(), "Available Verbs", "Available Subjects", -1))
		return command.Parent().UsageFunc()(command)
	})

	cmd.AddCommand(newRmAppCmd())
	cmd.AddCommand(newRmNodeCmd())
	cmd.AddCommand(newRmVmCmd())
	cmd.AddCommand(newRmIpCmd())
	cmd.AddCommand(newRmTestSystemCmd())

	for _, subCmd := range cmd.Commands() {
		cmd.ValidArgs = append(cmd.ValidArgs, subCmd.Name())
	}
	cmd.Args = cobra.OnlyValidArgs

	return cmd
}

func newRmAppCmd() *cobra.Command {
	return &cobra.Command{
		Use: "application",
	}
}

func newRmNodeCmd() *cobra.Command {
	return &cobra.Command{
		Use: "node",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
}

func newRmVmCmd() *cobra.Command {
	return &cobra.Command{
		Use: "virtual-machine",
	}
}

func newRmIpCmd() *cobra.Command {
	return &cobra.Command{
		Use: "ip",
	}
}

func newRmTestSystemCmd() *cobra.Command {
	return &cobra.Command{
		Use: "test-system",
	}
}
