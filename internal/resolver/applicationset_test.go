package resolver_test

import (
	"testing"

	"github.com/celeguim/emp-cli/internal/resolver"
	"github.com/celeguim/emp-cli/internal/testutil"
	"github.com/stretchr/testify/require"
)

func TestResolveApplicationSet(t *testing.T) {

	cat := testutil.NewCatalog()

	testutil.AddCluster(
		cat,
		"dev",
		"https://kubernetes.default.svc",
	)

	testutil.AddEnvironment(
		cat,
		"dev",
		"default",
		"dev",
		"default",
		"HEAD",
	)

	testutil.AddApplication(
		cat,
		"payments",
		"platform",
		"dev",
		"https://github.com/company/gitops.git",
		"services/payments",
	)

	testutil.AddApplication(
		cat,
		"orders",
		"platform",
		"dev",
		"https://github.com/company/gitops.git",
		"services/orders",
	)

	rc, err := resolver.Resolve(cat)
	require.NoError(t, err)

	appset := rc.ApplicationSets[0]

	require.Equal(t, "platform", appset.Name)
	require.Len(t, appset.Applications, 2)
}
