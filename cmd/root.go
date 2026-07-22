package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "emp",
	Short: "Enterprise Microservice Platform CLI",
}

func Execute() error {
	return rootCmd.Execute()
}
