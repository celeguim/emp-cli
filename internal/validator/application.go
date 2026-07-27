package validator

import (
	"fmt"
	"strings"

	"github.com/celeguim/emp-cli/internal/catalog"
)

func (v *Validator) validateApplications(
	cat *catalog.Catalog,
	report *Report,
) {
	v.validateApplicationNames(cat, report)
	v.validateApplicationCharts(cat, report)
}

func (v *Validator) validateApplicationCharts(
	cat *catalog.Catalog,
	report *Report,
) {
	for _, doc := range cat.Applications {
		if strings.TrimSpace(doc.Object.Chart) == "" {
			report.Add(Error{
				File:    doc.Path,
				Field:   "chart",
				Message: "is required",
			})
		}
	}
}

func (v *Validator) validateApplicationNames(
	cat *catalog.Catalog,
	report *Report,
) {

	seen := make(map[string]string)

	for _, doc := range cat.Applications {

		// name := doc.Object.AppName
		name := strings.TrimSpace(doc.Object.Name)
		if name == "" {
			report.Add(Error{
				File:    doc.Path,
				Name:    doc.Object.Name,
				Field:   "appName",
				Message: "is required",
			})
			continue
		}

		if previous, ok := seen[name]; ok {
			report.Add(Error{
				File:    doc.Path,
				Name:    doc.Object.Name,
				Field:   "appName",
				Message: fmt.Sprintf("duplicate application %q (already declared in %s)", name, previous),
			})
			continue
		}

		seen[name] = doc.Path

	}
}
