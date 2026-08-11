package validator

import (
	"testing"

	"github.com/celeguim/emp-cli/internal/catalog"
)

func env(path, name, project, cluster, repo, namespace, revision string) catalog.Document[catalog.Environment] {
	return catalog.Document[catalog.Environment]{
		Path: path,
		Object: catalog.Environment{
			Name:           name,
			Project:        project,
			Cluster:        cluster,
			RepoURL:        repo,
			Namespace:      namespace,
			TargetRevision: revision,
		},
	}
}

func TestValidEnvironment(t *testing.T) {
	cat := &catalog.Catalog{
		Environments: []catalog.Document[catalog.Environment]{
			env("env/env1.yaml", "env1", "project2", "cluster1", "repo:gitops", "payments", "main"),
		},
	}

	report := New().Validate(cat)

	if report.HasErrors() {
		t.Fatalf("expected no validation errors, got %v", report.Errors)
	}
}

func TestMissingEnvironmentName(t *testing.T) {
	cat := &catalog.Catalog{
		Environments: []catalog.Document[catalog.Environment]{
			env("env/env1.yaml", "", "project2", "cluster1", "repo:gitops", "payments", "main"),
		},
	}

	report := New().Validate(cat)

	if !report.HasErrors() {
		t.Fatal("expected validation error")
	}
}

func TestDuplicateEnvironmentName(t *testing.T) {
	cat := &catalog.Catalog{
		Environments: []catalog.Document[catalog.Environment]{
			env("env/env1.yaml", "env1", "project2", "cluster1", "repo:gitops", "payments", "main"),
			env("env/env1.yaml", "env1", "project2", "cluster1", "repo:gitops", "payments", "main"),
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
			env("env/env1.yaml", "env1", "", "cluster1", "repo:gitops", "payments", "main"),
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
			env("env/env1.yaml", "env1", "project2", "cluster1", "repo:gitops", "payments", ""),
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
			env("env/env1.yaml", "env1", "project2", "cluster1", "repo:gitops", "", "main"),
		},
	}

	report := New().Validate(cat)

	if !report.HasErrors() {
		t.Fatal("expected validation error")
	}
}
