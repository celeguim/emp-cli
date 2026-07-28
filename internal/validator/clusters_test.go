package validator

import (
	"testing"

	"github.com/celeguim/emp-cli/internal/catalog"
)

func cluster(path, name, server, environment string) catalog.Document[catalog.Cluster] {
	return catalog.Document[catalog.Cluster]{
		Path: path,
		Object: catalog.Cluster{
			Name:   name,
			Server: server,
			// Environment: environment,
		},
	}
}

func TestValidCluster(t *testing.T) {
	cat := &catalog.Catalog{
		Clusters: []catalog.Document[catalog.Cluster]{
			cluster(
				"clusters/dev.yaml",
				"dev-eks",
				"https://kubernetes.default.svc",
				"dev",
			),
		},
	}

	report := New().Validate(cat)

	if report.HasErrors() {
		t.Fatalf("expected no validation errors, got %v", report.Errors)
	}
}

func TestMissingClusterName(t *testing.T) {
	cat := &catalog.Catalog{
		Clusters: []catalog.Document[catalog.Cluster]{
			cluster(
				"clusters/dev.yaml",
				"",
				"https://kubernetes.default.svc",
				"dev",
			),
		},
	}

	report := New().Validate(cat)

	if !report.HasErrors() {
		t.Fatal("expected validation error")
	}
}

func TestDuplicateClusterName(t *testing.T) {
	cat := &catalog.Catalog{
		Clusters: []catalog.Document[catalog.Cluster]{
			cluster(
				"clusters/dev.yaml",
				"dev-eks",
				"https://kubernetes.default.svc",
				"dev",
			),
			cluster(
				"clusters/dev2.yaml",
				"dev-eks",
				"https://10.0.0.1",
				"dev",
			),
		},
	}

	report := New().Validate(cat)

	if !report.HasErrors() {
		t.Fatal("expected duplicate validation error")
	}
}

func TestMissingServer(t *testing.T) {
	cat := &catalog.Catalog{
		Clusters: []catalog.Document[catalog.Cluster]{
			cluster(
				"clusters/dev.yaml",
				"dev-eks",
				"",
				"dev",
			),
		},
	}

	report := New().Validate(cat)

	if !report.HasErrors() {
		t.Fatal("expected validation error")
	}
}

func TestMissingEnvironment(t *testing.T) {
	cat := &catalog.Catalog{
		Clusters: []catalog.Document[catalog.Cluster]{
			cluster(
				"clusters/dev.yaml",
				"dev-eks",
				"https://kubernetes.default.svc",
				"",
			),
		},
	}

	report := New().Validate(cat)

	if !report.HasErrors() {
		t.Fatal("expected validation error")
	}
}
