package validator

import (
	"testing"

	"github.com/celeguim/emp-cli/internal/catalog"
)

func TestValidSyncPolicyEnabled(t *testing.T) {
	cat := &catalog.Catalog{
		Environments: []catalog.Document[catalog.Environment]{
			env("env/dev.yaml", "dev", "project1", "cluster1", "namespace1", "revision1", "enabled"),
		},
	}

	report := New().Validate(cat)

	if report.HasErrors() {
		t.Fatalf("unexpected validation errors: %v", report.Errors)
	}
}

func TestInvalidSyncPolicy(t *testing.T) {
	cat := &catalog.Catalog{
		Environments: []catalog.Document[catalog.Environment]{
			env("env/dev.yaml", "dev", "project1", "cluster1", "namespace1", "revision1", "invalid"),
		},
	}

	report := New().Validate(cat)

	if !report.HasErrors() {
		t.Fatal("expected validation error")
	}
}
