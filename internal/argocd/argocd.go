package argocd

import "github.com/celeguim/emp-cli/internal/runner"

func Version() (*runner.Result, error) {
	return runner.Run("argocd", "version", "--client")
}

// func AppList() (*runner.Result, error) {
// 	return runner.Run("argocd", "app", "list")
// }
