package tools

type Tool interface {
	Name() string
	Exists() bool
	Version() string
}
