package runtime

type Runtime struct {
	Root string
}

func New(root string) *Runtime {
	return &Runtime{
		Root: root,
	}
}

func (r *Runtime) CreateWorkspace() error {
	return nil
}

func (r *Runtime) LoadCatalogs() error {
	return nil
}

func (r *Runtime) Write() error {
	return nil
}

func (r *Runtime) WriteManifest() error {
	return nil
}
