package resolver

import (
	"fmt"

	"github.com/celeguim/emp-cli/internal/catalog"
)

func (r *Resolver) resolveClusters(cat *catalog.Catalog) error {
	envs := make(map[string]struct{})

	for _, env := range cat.Environments {
		envs[env.Object.Name] = struct{}{}
	}

	for _, cluster := range cat.Clusters {
		if _, ok := envs[cluster.Object.Environment]; !ok {
			return fmt.Errorf(
				"cluster %q references unknown environment %q",
				cluster.Object.ClusterName,
				cluster.Object.Environment,
			)
		}
	}

	return nil
}
