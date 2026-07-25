package runtime

import (
	"fmt"

	"github.com/celeguim/emp-cli/internal/validator"
)

func (r *Runtime) Render() error {
	fmt.Println("CreateWorkspace")
	if err := r.CreateWorkspace(); err != nil {
		return err
	}

	fmt.Println("LoadCatalogs")
	if err := r.LoadCatalogs(); err != nil {
		return err
	}

	report := validator.New().Validate(r)

	if report.HasErrors() {
		// por enquanto só imprime
		// depois melhoramos a formatação
	}

	fmt.Printf("Apps=%d Envs=%d Clusters=%d\n",
		len(r.Applications),
		len(r.Environments),
		len(r.Clusters),
	)

	fmt.Println("Write")
	if err := r.Write(); err != nil {
		return err
	}

	fmt.Println("WriteManifest")
	return r.WriteManifest()
}
