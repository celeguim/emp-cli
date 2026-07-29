[ ] AppProject Renderer
[ ] ApplicationSet Renderer
[ ] Repositories (caso necessário)
[ ] Doctor
[ ] Diff
[ ] Export Graph

1. Testes do Loader (internal/catalog) ← pega exatamente esse tipo de regressão.
2. Testes do Resolver.
3. Testes dos Manifestos.
4. Golden tests dos Renderers.

✅ catalog/loader_test.go (protege o carregamento dos dados)
✅ resolver/*_test.go (protege a lógica de negócio)
✅ compiler/manifests/*_test.go (protege a transformação para manifestos)
✅ Golden tests (protegem o YAML final)