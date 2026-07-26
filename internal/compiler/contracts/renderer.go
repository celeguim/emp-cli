package contracts

import "github.com/celeguim/emp-cli/internal/resolved"

type Renderer interface {
	Render(ctx *Context, cat *resolved.Catalog) error
}
