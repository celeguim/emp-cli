package renderers_test

// import (
// 	"os"
// 	"testing"

// 	"github.com/celeguim/emp-cli/internal/catalog"
// 	"github.com/celeguim/emp-cli/internal/compiler/renderers"
// 	"github.com/celeguim/emp-cli/internal/resolved"
// 	"github.com/stretchr/testify/require"
// )

// func TestApplicationRenderer(t *testing.T) {
// 	app := resolved.Application{
// 		Application: catalog.Application{
// 			Name:        "payments",
// 			Environment: "dev",
// 			RepoURL:     "https://github.com/company/gitops.git",
// 			Path:        "services/payments",
// 		},

// 		Environment: catalog.Environment{
// 			Name:           "dev",
// 			Project:        "default",
// 			Cluster:        "dev",
// 			Namespace:      "default",
// 			TargetRevision: "HEAD",
// 		},

// 		Cluster: catalog.Cluster{
// 			Name:   "dev",
// 			Server: "https://kubernetes.default.svc",
// 		},
// 	}

// 	renderer := renderers.Application{}
// 	got, err := renderer.Render(app)
// 	require.NoError(t, err)

// 	expected, err := os.ReadFile("testdata/application.yaml")
// 	require.NoError(t, err)

// 	require.YAMLEq(
// 		t,
// 		string(expected),
// 		string(got),
// 	)
// }
