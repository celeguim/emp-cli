package catalog

type Document[T any] struct {
	Path   string
	Object T
}
