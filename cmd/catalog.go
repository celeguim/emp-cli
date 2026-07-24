package cmd

import "github.com/spf13/cobra"

var catalogCmd = &cobra.Command{
	Use:   "catalog",
	Short: "Manage EMP catalogs",
}

func init() {
	rootCmd.AddCommand(catalogCmd)
}
