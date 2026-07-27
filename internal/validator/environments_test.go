package validator

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/celeguim/emp-cli/internal/catalog"
)

func env(path, name, project, revision, namespace string) catalog.Document[catalog.Environment] {
	return catalog.Document[catalog.Environment]{
		Path: path,
		Object: catalog.Environment{
			Name:           name,
			Project:        project,
			TargetRevision: revision,
			Namespace:      namespace,
		},
	}
}

func TestValidEnvironment(t *testing.T) {
	cat := &catalog.Catalog{
		Environments: []catalog.Document[catalog.Environment]{
			env("env/dev.yaml", "dev", "payments", "main", "payments"),
		},
	}

	report := New().Validate(cat)

	if report.HasErrors() {
		t.Fatalf("expected no validation errors, got %v", report.Errors)
	}
}

func TestMissingEnvironmentName(t *testing.T) {
	fmt.Println(filepath.Abs("."))

	loader := catalog.NewFilesystemLoader("../..")
	cat, err := loader.Load()
	if err != nil {
		t.Fatalf("failed to load catalog: %v", err)
	}

	// cat := &catalog.Catalog{
	// 	Environments: []catalog.Document[catalog.Environment]{
	// 		env("env/dev.yaml", "", "payments", "main", "payments"),
	// 	},
	// }

	report := New().Validate(cat)
	fmt.Println("REPORT ", report)

	if !report.HasErrors() {
		t.Fatal("expected validation error")
	}
}

func TestDuplicateEnvironmentName(t *testing.T) {
	cat := &catalog.Catalog{
		Environments: []catalog.Document[catalog.Environment]{
			env("env/dev.yaml", "dev", "payments", "main", "payments"),
			env("env/dev2.yaml", "dev", "payments", "main", "payments"),
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
			env("env/dev.yaml", "dev", "", "main", "payments"),
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
			env("env/dev.yaml", "dev", "payments", "", "payments"),
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
			env("env/dev.yaml", "dev", "payments", "main", ""),
		},
	}

	report := New().Validate(cat)

	if !report.HasErrors() {
		t.Fatal("expected validation error")
	}
}
