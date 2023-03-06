/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

// gatewayCmd represents the gateway command
var gatewayCmd = &cobra.Command{
	Use:   "gateway",
	Short: "run gateway service",
	Long:  `Run gateway service`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		return gracefulShutdown(ctx, gateway)
	},
}

func init() {
	rootCmd.AddCommand(gatewayCmd)
}

func gateway(ctx context.Context) error {
	return nil
}
