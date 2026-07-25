package validator

import (
	"testing"

	"github.com/celeguim/emp-cli/internal/catalog"
)

func app(path, name, chart string) catalog.Document[catalog.Application] {
	return catalog.Document[catalog.Application]{
		Path: path,
		Object: catalog.Application{
			AppName: name,
			Chart:   chart,
		},
	}
}

func TestApplicationNameRequired(t *testing.T) {

	cat := &catalog.Catalog{
		Applications: []catalog.Document[catalog.Application]{
			app("api.yaml", "", "charts/api"),
		},
	}

	report := New().Validate(cat)

	if !report.HasErrors() {
		t.Fatal("expected validation error")
	}
}

func TestDuplicateApplicationName(t *testing.T) {

	cat := &catalog.Catalog{
		Applications: []catalog.Document[catalog.Application]{
			app("api.yaml", "payment", "charts/payment"),
			app("web.yaml", "payment", "charts/web"),
		},
	}

	report := New().Validate(cat)

	if !report.HasErrors() {
		t.Fatal("expected duplicate validation error")
	}
}

func TestChartRequired(t *testing.T) {

	cat := &catalog.Catalog{
		Applications: []catalog.Document[catalog.Application]{
			app("app.yaml", "app1", ""),
		},
	}

	report := New().Validate(cat)

	if !report.HasErrors() {
		t.Fatal("expected validation error")
	}
}

func TestValidApplication(t *testing.T) {

	cat := &catalog.Catalog{
		Applications: []catalog.Document[catalog.Application]{
			app("api.yaml", "payment", "charts/payment"),
			app("web.yaml", "frontend", "charts/frontend"),
		},
	}

	report := New().Validate(cat)

	if report.HasErrors() {
		t.Fatalf("unexpected validation errors: %v", report.Errors)
	}
}
