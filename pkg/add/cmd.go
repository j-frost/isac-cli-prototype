package add

import (
	"github.com/spf13/cobra"
	"strings"
)

func NewAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add SUBJECT OPERAND",
		Short: "Add an entity to isac",
		Long: `Add a TYPE type entity SUBJECT to isac, where TYPE may be one of 
application, node, virtual-machine, ip, or test-system. The value of SUBJECT 
varies with TYPE. `,
		Aliases: []string{"a", "create", "c"},
	}

	cmd.SetUsageFunc(func(command *cobra.Command) error {
		if !command.HasParent() {
			return command.UsageFunc()(command)
		}
		command.SetUsageTemplate(strings.Replace(command.Parent().UsageTemplate(), "Available Verbs", "Available Subjects", -1))
		return command.Parent().UsageFunc()(command)
	})

	cmd.AddCommand(newAddAppCmd())
	cmd.AddCommand(newAddNodeCmd())
	cmd.AddCommand(newAddVmCmd())
	cmd.AddCommand(newAddIpCmd())
	cmd.AddCommand(newAddTestSystemCmd())

	for _, subCmd := range cmd.Commands() {
		cmd.ValidArgs = append(cmd.ValidArgs, subCmd.Name())
	}
	cmd.Args = cobra.OnlyValidArgs

	return cmd
}

func newAddAppCmd() *cobra.Command {
	return &cobra.Command{
		Use: "application TYPE",
	}
}

func newAddNodeCmd() *cobra.Command {
	return &cobra.Command{
		Use: "node BW_NODE_NAME",
		Short: "Adds a new bundlewrap managed node to isac. ",
		Long: ``,
		Example: "isac-cli add node pb-0.smedia.customer-product --draft=6877b0ea",
		Aliases: []string{"n"},
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
}

func newAddVmCmd() *cobra.Command {
	return &cobra.Command{
		Use: "virtual-machine",
	}
}

func newAddIpCmd() *cobra.Command {
	return &cobra.Command{
		Use: "ip",
	}
}

func newAddTestSystemCmd() *cobra.Command {
	return &cobra.Command{
		Use: "test-system",
	}
}
