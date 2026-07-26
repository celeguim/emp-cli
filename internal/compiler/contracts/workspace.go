package contracts

type Workspace struct {
	Root string
}

func (w *Workspace) CompilerDir() string
func (w *Workspace) ApplicationsDir() string
func (w *Workspace) EnvironmentsDir() string
func (w *Workspace) ClustersDir() string
func (w *Workspace) Create() error
