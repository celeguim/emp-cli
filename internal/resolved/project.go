package resolved

type Project struct {
	Name         string
	Destinations []Destination
	SourceRepos  []string
}

type Destination struct {
	Server    string
	Namespace string
}
