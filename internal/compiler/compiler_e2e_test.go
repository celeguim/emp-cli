package compiler

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/celeguim/emp-cli/internal/catalog"
	"github.com/celeguim/emp-cli/internal/compiler/contracts"
	"github.com/celeguim/emp-cli/internal/resolver"
	"github.com/stretchr/testify/require"
)

func TestCompiler(t *testing.T) {
	out := filepath.Join("testdata", "output")

	_ = os.RemoveAll(out)
	require.NoError(t, os.MkdirAll(out, 0755))

	ctx := &contracts.Context{
		Root: out,
	}

	t.Logf("output dir: %s", out)

	loader := catalog.NewFilesystemLoader("testdata")

	cat, err := loader.Load()
	require.NoError(t, err)

	resolved, err := resolver.Resolve(cat)
	require.NoError(t, err)

	// compiler := compiler.New(ctx)
	c := NewCompiler(ctx.Root)

	require.NoError(t, c.Compile(resolved))

	entries, err := os.ReadDir(out)
	// require.NoError(t, err)
	t.Logf("entries: %+v", entries)

	for _, e := range entries {
		t.Log(e.Name())
	}

	// require.NotEmpty(t, entries)
	// info, errs := os.Stat(emp)
	// println(info)
	// require.NoError(t, errs)

	emp := filepath.Join(out, ".emp")
	apps := filepath.Join(emp, "applications")
	entries, errx := os.ReadDir(apps)
	require.NoError(t, errx)
	require.Len(t, entries, 1)
	require.Equal(t, "gitops1.yaml", entries[0].Name())
	t.Logf("app gitops1.yaml criado")

	projects := filepath.Join(emp, "projects")
	entriesp, errp := os.ReadDir(projects)
	require.NoError(t, errp)
	require.Len(t, entriesp, 1)
	require.Equal(t, "project1.yaml", entriesp[0].Name())
	t.Logf("project project1.yaml criado")

	appsets := filepath.Join(emp, "applicationsets")
	entriesa, erra := os.ReadDir(appsets)
	require.NoError(t, erra)
	require.Len(t, entriesa, 1)
	require.Equal(t, "dev1.yaml", entriesa[0].Name())
	t.Logf("appset dev1.yaml criado")

	// application validation
	got, err := os.ReadFile(filepath.Join(
		out,
		".emp",
		"applications",
		"gitops1.yaml",
	))
	require.NoError(t, err)

	expected, err := os.ReadFile(filepath.Join(
		"testdata",
		"expected",
		"applications",
		"gitops1.yaml",
	))
	require.NoError(t, err)

	require.YAMLEq(t, string(expected), string(got))

	// project validation
	gotp, errp := os.ReadFile(filepath.Join(
		out,
		".emp",
		"projects",
		"project1.yaml",
	))
	require.NoError(t, errp)

	expectedp, errp := os.ReadFile(filepath.Join(
		"testdata",
		"expected",
		"projects",
		"project1.yaml",
	))
	require.NoError(t, errp)

	require.YAMLEq(t, string(expectedp), string(gotp))

	// applicationset validation
	gota, erra := os.ReadFile(filepath.Join(
		out,
		".emp",
		"applicationsets",
		"dev1.yaml",
	))
	require.NoError(t, erra)

	expecteda, erra := os.ReadFile(filepath.Join(
		"testdata",
		"expected",
		"applicationsets",
		"dev1.yaml",
	))
	require.NoError(t, erra)

	require.YAMLEq(t, string(expecteda), string(gota))
	println("OK\n")
}
