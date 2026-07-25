package validator

type Report struct {
	Errors []Error
}

func (r *Report) Add(err Error) {
	r.Errors = append(r.Errors, err)
}

func (r Report) HasErrors() bool {
	return len(r.Errors) > 0
}
