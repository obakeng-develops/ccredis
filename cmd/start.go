/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/obakeng-develops/redis-server/server"
	"github.com/spf13/cobra"
)

type startOptions struct{}

var startCmdOptions = &startOptions{}

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the server",
	Run: func(cmd *cobra.Command, args []string) {
		startCmdOptions.run()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}

func (options *startOptions) run() error {
	server := &server.Server{}
	server.NewServer(":6379")

	server.StartServer()
	return nil
}
