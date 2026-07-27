package validator

import (
	"fmt"

	"github.com/celeguim/emp-cli/internal/catalog"
)

type Validator struct {
}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) Validate(cat *catalog.Catalog) Report {
	var report Report

	println("Validating catalog...")
	fmt.Println("Apps:", len(cat.Applications))
	fmt.Println("Envs:", len(cat.Environments))
	fmt.Println("Clusters:", len(cat.Clusters))

	v.validateApplications(cat, &report)
	v.validateEnvironments(cat, &report)
	v.validateClusters(cat, &report)

	return report
}
