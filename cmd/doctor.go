package cmd

import (
	"fmt"
	"strings"

	"github.com/celeguim/emp-cli/internal/doctor"
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use:   "doctor",
	Short: "Check EMP environment",

	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("EMP Doctor")
		fmt.Println()

		checks := doctor.Run()
		currentCategory := ""

		for _, check := range checks {

			if check.Category != currentCategory {

				if currentCategory != "" {
					fmt.Println()
				}

				fmt.Println(check.Category)
				fmt.Println(strings.Repeat("-", len(check.Category)))

				currentCategory = check.Category
			}

			status := "✓"
			if !check.OK {
				status = "✗"
			}

			fmt.Printf("%s %-12s %s\n",
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
