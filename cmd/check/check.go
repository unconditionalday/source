package check

import "github.com/spf13/cobra"

func NewCheckCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "check",
		Short: "Check source",
		Long:  `Check source`,
	}

	cmd.AddCommand(NewRSSCheckCmd())

	return cmd
}
