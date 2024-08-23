package objectgraph

import (
	"github.com/olyop/objectgraph/objectgraph/internal/objectcache"
	"github.com/vektah/gqlparser/ast"
)

type Engine struct {
	schemaAst   *ast.Schema
	schema      *Schema
	objectcache *objectcache.ObjectCache
	retreivers  retreivers
}
