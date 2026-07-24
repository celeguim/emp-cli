package runtime

func Render(root string) error {
	rt := New(root)

	if err := rt.CreateWorkspace(); err != nil {
		return err
	}

	if err := rt.LoadCatalogs(); err != nil {
		return err
	}

	if err := rt.Write(); err != nil {
		return err
	}

	return rt.WriteManifest()
}
