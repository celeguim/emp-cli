package validator

import "github.com/celeguim/emp-cli/internal/catalog"

type Validator struct {
}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) Validate(cat *catalog.Catalog) Report {

	var report Report

	v.validateApplications(cat, &report)

	return report
}
