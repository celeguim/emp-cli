package validator

import (
	"fmt"
	"strings"

	"github.com/celeguim/emp-cli/internal/catalog"
)

func (v *Validator) validateClusters(
	cat *catalog.Catalog,
	report *Report,
) {
	v.validateClusterNames(cat, report)
	v.validateClusterServers(cat, report)
	v.validateClusterEnvironments(cat, report)
}

func (v *Validator) validateClusterNames(
	cat *catalog.Catalog,
	report *Report,
) {
	seen := make(map[string]string)

	for _, doc := range cat.Clusters {

		name := strings.TrimSpace(doc.Object.ClusterName)

		if name == "" {
			report.Add(Error{
				File:    doc.Path,
				Name:    doc.Object.ClusterName,
				Field:   "clusterName",
				Message: "is required",
			})
			continue
		}

		if previous, ok := seen[name]; ok {
			report.Add(Error{
				File:  doc.Path,
				Name:  doc.Object.ClusterName,
				Field: "clusterName",
				Message: fmt.Sprintf(
					"duplicate cluster %q (already declared in %s)",
					name,
					previous,
				),
			})
			continue
		}

		seen[name] = doc.Path
	}
}

func (v *Validator) validateClusterServers(
	cat *catalog.Catalog,
	report *Report,
) {
	for _, doc := range cat.Clusters {

		if strings.TrimSpace(doc.Object.Server) == "" {
			report.Add(Error{
				File:    doc.Path,
				Name:    doc.Object.ClusterName,
				Field:   "server",
				Message: "is required",
			})
		}
	}
}

func (v *Validator) validateClusterEnvironments(
	cat *catalog.Catalog,
	report *Report,
) {
	for _, doc := range cat.Clusters {

		if strings.TrimSpace(doc.Object.Environment) == "" {
			report.Add(Error{
				File:    doc.Path,
				Name:    doc.Object.ClusterName,
				Field:   "environment",
				Message: "is required",
			})
		}
	}
}
