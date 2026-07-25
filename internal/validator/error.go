package validator

type Error struct {
	File    string
	Name    string
	Field   string
	Message string
}
