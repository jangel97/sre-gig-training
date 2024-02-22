package command

import (
	"fmt"

	"github.com/spf13/cobra"
	// Update the import path according to your project structure
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "chatcli",
		Short: "Chat CLI is a command line interface for entretainment purposes",
		Long:  `This application is a simple CLI tool built with Cobra for entretainment purposes.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Welcome to the Chat CLI! Run help command to see available options")
			fmt.Println(`Examples on how to run this:
			- chatcli joke --category Miscellaneous
			- chatcli joke
			- chatcli philosophy`)
		},
	}

	initCommands(rootCmd)
	return rootCmd
}

func initCommands(rootCmd *cobra.Command) {
	rootCmd.AddCommand(newJokeCmd())
	rootCmd.AddCommand(newPhilosophicalQuoteCmd())
	// Add more subcommands here
}
