package catalog

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {

	loader := NewFilesystemLoader("testdata/catalog")

	cat, err := loader.Load()

	require.NoError(t, err)

	require.Len(t, cat.Applications, 2)
	require.Len(t, cat.Environments, 1)
	require.Len(t, cat.Clusters, 1)
}
