package svc

import "github.com/spf13/cobra"

func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "svc",
		Short: "Manage services",
	}

	cmd.AddCommand(startCmd)

	return cmd
}
