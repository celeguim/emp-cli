package resolved

type Project struct {
	Name         string
	Destinations []Destination
	SourceRepos  []string
}

type Destination struct {
	Name      string
	Server    string
	Namespace string
}
