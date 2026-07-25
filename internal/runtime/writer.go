package runtime

import (
	"fmt"
	"os"

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
