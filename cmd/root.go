package cmd

import (
	"github.com/spf13/cobra"
	"github.com/omabena/api_earn_spots/cmd/svc"
)

// NewRootCmd start main command functions
func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "Challenge",
		Short: "CLI interface for challenge",
	}

	cmd.AddCommand(svc.NewCmd())

	return cmd
}
