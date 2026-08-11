package validator

import (
	"testing"

	"github.com/celeguim/emp-cli/internal/catalog"
)

// Name           string `yaml:"name"`
// Project        string `yaml:"project"`
// Cluster        string `yaml:"cluster"`
// Namespace      string `yaml:"namespace"`
// TargetRevision string `yaml:"targetRevision"`
// SyncPolicy     string `yaml:"syncPolicy,omitempty"`

func env(path, name, project, cluster, namespace, revision, syncpolicy string) catalog.Document[catalog.Environment] {
	return catalog.Document[catalog.Environment]{
		Path: path,
		Object: catalog.Environment{
			Name:           name,
			Project:        project,
			Cluster:        cluster,
			Namespace:      namespace,
			TargetRevision: revision,
			SyncPolicy:     syncpolicy,
		},
	}
}

func TestValidEnvironment(t *testing.T) {
	cat := &catalog.Catalog{
		Environments: []catalog.Document[catalog.Environment]{
			env("env/dev.yaml", "dev", "project1", "cluster1", "namespace1", "revision1", "enabled"),
		},
	}

	report := New().Validate(cat)

	if report.HasErrors() {
		t.Fatalf("expected no validation errors, got %v", report.Errors)
	}
}

// func TestMissingEnvironmentName(t *testing.T) {
// 	fmt.Println(filepath.Abs("."))

// 	loader := catalog.NewFilesystemLoader("../..")
// 	cat, err := loader.Load()
// 	if err != nil {
// 		t.Fatalf("failed to load catalog: %v", err)
// 	}

// 	report := New().Validate(cat)
// 	fmt.Println("REPORT ", report)

// 	if !report.HasErrors() {
// 		t.Fatal("expected validation error")
// 	}
// }

func TestDuplicateEnvironmentName(t *testing.T) {
	cat := &catalog.Catalog{
		Environments: []catalog.Document[catalog.Environment]{
			env("env/dev.yaml", "dev", "project1", "cluster1", "namespace1", "revision1", "enabled"),
			env("env/dev.yaml", "dev", "project1", "cluster1", "namespace1", "revision1", "enabled"),
		},
	}

	report := New().Validate(cat)

	if !report.HasErrors() {
		t.Fatal("expected duplicate validation error")
	}
}

func TestMissingProject(t *testing.T) {
	cat := &catalog.Catalog{
		Environments: []catalog.Document[catalog.Environment]{
			env("env/dev.yaml", "dev", "", "cluster1", "namespace1", "revision1", "enabled"),
		},
	}

	report := New().Validate(cat)

	if !report.HasErrors() {
		t.Fatal("expected validation error")
	}
}

func TestMissingTargetRevision(t *testing.T) {
	cat := &catalog.Catalog{
		Environments: []catalog.Document[catalog.Environment]{
			env("env/dev.yaml", "dev", "project1", "cluster1", "namespace1", "", "enabled"),
		},
	}

	report := New().Validate(cat)

	if !report.HasErrors() {
		t.Fatal("expected validation error")
	}
}

func TestMissingNamespace(t *testing.T) {
	cat := &catalog.Catalog{
		Environments: []catalog.Document[catalog.Environment]{
			env("env/dev.yaml", "dev", "project1", "cluster1", "", "revision1", "enabled")},
	}

	report := New().Validate(cat)

	if !report.HasErrors() {
		t.Fatal("expected validation error")
	}
}
