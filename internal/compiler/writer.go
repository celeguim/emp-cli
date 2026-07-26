package compiler

import (
	"fmt"
	"os"

	"github.com/celeguim/emp-cli/internal/catalog"
	"sigs.k8s.io/yaml"
)

func writeYAML(path string, obj any) error {
	data, err := yaml.Marshal(obj)

	fmt.Printf("writing %s\n", path)

	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0o644)
}

func (c *Compiler) Write(cat *catalog.Catalog) error {

	if err := c.enderApplications(cat.Applications); err != nil {
		return err
	}

	if err := c.RenderEnvironments(cat.Environments); err != nil {
		return err
	}

	if err := c.RenderClusters(cat.Clusters); err != nil {
		return err
	}

	return nil
}
