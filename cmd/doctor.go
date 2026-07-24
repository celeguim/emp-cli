package cmd

import (
	"fmt"

	"github.com/celeguim/emp-cli/internal/doctor"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check EMP environment",

	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("EMP Doctor")
		fmt.Println()

		for _, check := range doctor.Run() {

			status := "✓"
			if !check.OK {
				status = "✗"
			}

			fmt.Printf("%s %-10s %s\n",
				status,
				check.Name,
				check.Message,
			)
		}
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
