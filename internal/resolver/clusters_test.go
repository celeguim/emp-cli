package resolver_test

// func TestResolveClusterEnvironment(t *testing.T) {
// 	cat := &catalog.Catalog{
// 		Environments: []catalog.Document[catalog.Environment]{
// 			{
// 				Path: "environments/dev.yaml",
// 				Object: catalog.Environment{
// 					Name: "dev",
// 				},
// 			},
// 		},
// 		Clusters: []catalog.Document[catalog.Cluster]{
// 			{
// 				Path: "clusters/dev-eks.yaml",
// 				Object: catalog.Cluster{
// 					Name: "dev-eks",
// 					// Environment: "dev",
// 				},
// 			},
// 		},
// 	}

// 	err := New().resolveClusters(cat)
// 	if err != nil {
// 		t.Fatalf("unexpected error: %v", err)
// 	}
// }

// func TestResolveUnknownEnvironment(t *testing.T) {
// 	cat := &catalog.Catalog{
// 		Clusters: []catalog.Document[catalog.Cluster]{
// 			{
// 				Path: "clusters/dev-eks.yaml",
// 				Object: catalog.Cluster{
// 					Name: "dev-eks",
// 					// Environment: "dev",
// 				},
// 			},
// 		},
// 	}

// 	err := New().resolveClusters(cat)
// 	if err == nil {
// 		t.Fatal("expected resolver error")
// 	}
// }
