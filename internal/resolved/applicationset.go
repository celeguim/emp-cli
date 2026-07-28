package resolved

type ApplicationSet struct {
	Name         string
	Applications []Application
	Signature    ApplicationSetSignature
}

type ApplicationSetSignature struct {
	Project        string
	RepoURL        string
	TargetRevision string
	Server         string
	Namespace      string
}
