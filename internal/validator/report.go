package validator

import (
	"fmt"
	"strings"
)

type Report struct {
	Errors []Error
	// Warnings []Warning
}

func (r *Report) Add(err Error) {
	r.Errors = append(r.Errors, err)
}

func (r Report) HasErrors() bool {
	return len(r.Errors) > 0
}

func (r Report) Error() string {
	var b strings.Builder

	fmt.Fprintf(&b, "%d validation error(s):\n\n", len(r.Errors))

	for _, err := range r.Errors {
		fmt.Fprintf(&b,
			"%s: %s (%s)\n",
			err.File,
			err.Message,
			err.Field,
		)
	}

	return b.String()
}
