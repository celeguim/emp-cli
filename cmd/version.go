package cmd

import (
	"fmt"
	"runtime"

	"github.com/celeguim/emp-cli/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show EMP CLI version",
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println("Enterprise Microservice Platform CLI")
		fmt.Println("")
		fmt.Printf("Version    : %s\n", version.Version)
		fmt.Printf("Git Commit : %s\n", version.GitCommit)
		fmt.Printf("Build Date : %s\n", version.BuildDate)
		fmt.Printf("Go Version : %s\n", runtime.Version())
		fmt.Printf("OS/Arch    : %s/%s\n", runtime.GOOS, runtime.GOARCH)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
