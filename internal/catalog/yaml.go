package catalog

import (
	"fmt"
	"os"
	"path/filepath"

	"sigs.k8s.io/yaml"
)

func loadYAMLFiles[T any](dir string) ([]Document[T], error) {
	pattern := filepath.Join(dir, "*.yaml")
	fmt.Printf("pattern: %s \n", pattern)

	files, err := filepath.Glob(pattern)
	if err != nil {
		return nil, err
	}

	documents := make([]Document[T], 0, len(files))

	for _, file := range files {
		// fmt.Printf("file: %s", file)

		data, err := os.ReadFile(file)
		if err != nil {
			return nil, fmt.Errorf("reading %s: %w", file, err)
		}

		var obj T

		if err := yaml.Unmarshal(data, &obj); err != nil {
			return nil, fmt.Errorf("parsing %s: %w", file, err)
		}

		documents = append(documents, Document[T]{
			Path:   file,
			Object: obj,
		})
	}

	return documents, nil
}
