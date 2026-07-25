package runtime

import "fmt"

func (r *Runtime) Write() error {

	fmt.Printf(
		"apps=%d envs=%d clusters=%d\n",
		len(r.Applications),
		len(r.Environments),
		len(r.Clusters),
	)

	if err := r.RenderApplications(); err != nil {
		return err
	}

	if err := r.RenderEnvironments(); err != nil {
		return err
	}

	if err := r.RenderClusters(); err != nil {
		return err
	}

	return nil
}
