package validator

import (
	"fmt"
	"strings"

	"github.com/celeguim/emp-cli/internal/catalog"
)

func (v *Validator) validateEnvironments(
	cat *catalog.Catalog,
	report *Report,
) {
	v.validateEnvironmentNames(cat, report)
	v.validateEnvironmentProjects(cat, report)
	v.validateEnvironmentTargetRevisions(cat, report)
	v.validateEnvironmentNamespaces(cat, report)
}

func (v *Validator) validateEnvironmentNames(
	cat *catalog.Catalog,
	report *Report,
) {
	seen := make(map[string]string)

	for _, doc := range cat.Environments {

		name := strings.TrimSpace(doc.Object.Name)

		if name == "" {
			report.Add(Error{
				File:    doc.Path,
				Name:    doc.Object.Name,
				Field:   "envName",
				Message: "is required",
			})
			continue
		}

		if previous, ok := seen[name]; ok {
			report.Add(Error{
				File:  doc.Path,
				Name:  doc.Object.Name,
				Field: "envName",
				Message: fmt.Sprintf(
					"duplicate environment %q (already declared in %s)",
					name,
					previous,
				),
			})
			continue
		}

		seen[name] = doc.Path
	}
}

func (v *Validator) validateEnvironmentProjects(
	cat *catalog.Catalog,
	report *Report,
) {
	for _, doc := range cat.Environments {

		if strings.TrimSpace(doc.Object.Project) == "" {
			report.Add(Error{
				File:    doc.Path,
				Name:    doc.Object.Name,
				Field:   "project",
				Message: "is required",
			})
		}
	}
}

func (v *Validator) validateEnvironmentTargetRevisions(
	cat *catalog.Catalog,
	report *Report,
) {
	for _, doc := range cat.Environments {

		if strings.TrimSpace(doc.Object.TargetRevision) == "" {
			report.Add(Error{
				File:    doc.Path,
				Name:    doc.Object.Name,
				Field:   "targetRevision",
				Message: "is required",
			})
		}
	}
}

func (v *Validator) validateEnvironmentNamespaces(
	cat *catalog.Catalog,
	report *Report,
) {
	for _, doc := range cat.Environments {

		if strings.TrimSpace(doc.Object.Namespace) == "" {
			report.Add(Error{
				File:    doc.Path,
				Name:    doc.Object.Name,
				Field:   "namespace",
				Message: "is required",
			})
		}
	}
}
