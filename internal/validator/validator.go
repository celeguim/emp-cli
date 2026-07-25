package validator

type Validator struct {
}

func New() *Validator {
	return &Validator{}
}

// func (v *Validator) Validate(cat *catalog.Catalog) Report {
// 	var report Report

// 	v.validateApplications(rt, &report)
// 	v.validateEnvironments(rt, &report)
// 	v.validateClusters(rt, &report)

// 	return report
// }
