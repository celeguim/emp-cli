package validator

type Error struct {
	Kind    string
	Name    string
	File    string
	Field   string
	Message string
}
