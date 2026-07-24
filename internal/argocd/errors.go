package argocd

import (
	"fmt"
	"strings"
)

func HandleError(stderr string, err error) error {

	switch {

	case strings.Contains(stderr, "Unauthenticated"):
		return fmt.Errorf("You are not logged into Argo CD. Run 'argocd login'")

	case strings.Contains(stderr, "connection refused"),
		strings.Contains(stderr, "Unavailable"),
		strings.Contains(stderr, "Error while dialing"),
		strings.Contains(stderr, "connect: connection refused"):
		return fmt.Errorf("Unable to connect to Argo CD server.\n")

	default:
		return err
	}
}
