package validator

import (
	"fmt"

	"github.com/celeguim/emp-cli/internal/runtime"
)

func (v *Validator) validateApplications(
	rt *runtime.Runtime,
	report *Report,
) {

	seen := map[string]string{}

	for _, doc := range rt.Applications {

		name := doc.Object.Metadata.Name

		if name == "" {

			report.Add(Error{
				File:    doc.Path,
				Field:   "metadata.name",
				Message: "is required",
			})

			continue
		}

		if previous, ok := seen[name]; ok {

			report.Add(Error{
				File:  doc.Path,
				Field: "metadata.name",
				Message: fmt.Sprintf(
					"duplicated (already declared in %s)",
					previous,
				),
			})

			continue
		}

		seen[name] = doc.Path
	}
}
