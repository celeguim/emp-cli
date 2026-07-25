package validator

import "github.com/celeguim/emp-cli/internal/runtime"

type Validator struct {
}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) Validate(rt *runtime.Runtime) Report {

	var report Report

	v.validateApplications(rt, &report)
	v.validateEnvironments(rt, &report)
	v.validateClusters(rt, &report)

	return report
}
