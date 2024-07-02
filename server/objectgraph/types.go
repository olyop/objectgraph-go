package objectgraph

type RetrieverMap map[string]*RetrieverType
type RetrieverType struct {
	PrimaryKey string
	Value      any
}
