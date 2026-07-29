package resolver_test

import (
	"testing"

	"github.com/celeguim/emp-cli/internal/resolver"
	"github.com/celeguim/emp-cli/internal/testutil"
	"github.com/stretchr/testify/require"
)

func TestResolveProjects(t *testing.T) {

	cat := testutil.DefaultCatalog()

	testutil.AddApplication(
		cat,
		"payments",
		"",
		"dev",
		"https://github.com/celeguim/gitops.git",
		"services/payments",
	)

	testutil.AddApplication(
		cat,
		"orders",
		"",
		"dev",
		"https://github.com/celeguim/gitops.git",
		"services/orders",
	)

	rc, err := resolver.Resolve(cat)

	require.NoError(t, err)

	require.Len(t, rc.Projects, 1)

	project := rc.Projects[0]

	require.Equal(t, "default", project.Name)

	require.Len(t, project.Destinations, 1)

	require.Len(t, project.SourceRepos, 1)

	require.Equal(
		t,
		"https://github.com/celeguim/gitops.git",
		project.SourceRepos[0],
	)
}
