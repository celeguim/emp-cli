package resolver

import (
	"testing"

	"github.com/celeguim/emp-cli/internal/catalog"
)

func TestResolveClusterEnvironment(t *testing.T) {
	cat := &catalog.Catalog{
		Environments: []catalog.Document[catalog.Environment]{
			{
				Path: "environments/dev.yaml",
				Object: catalog.Environment{
					EnvName: "dev",
				},
			},
		},
		Clusters: []catalog.Document[catalog.Cluster]{
			{
				Path: "clusters/dev-eks.yaml",
				Object: catalog.Cluster{
					ClusterName: "dev-eks",
					Environment: "dev",
				},
			},
		},
	}

	err := New().resolveClusters(cat)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestResolveUnknownEnvironment(t *testing.T) {
	cat := &catalog.Catalog{
		Clusters: []catalog.Document[catalog.Cluster]{
			{
				Path: "clusters/dev-eks.yaml",
				Object: catalog.Cluster{
					ClusterName: "dev-eks",
					Environment: "dev",
				},
			},
		},
	}

	err := New().resolveClusters(cat)
	if err == nil {
		t.Fatal("expected resolver error")
	}
}
