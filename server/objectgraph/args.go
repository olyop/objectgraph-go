package objectgraph

type RetrieverArgs struct {
	primary any
	args    map[string]any
}

func (ra *RetrieverArgs) GetPrimary() any {
	return ra.primary
}

func (ra *RetrieverArgs) GetArg(key string) any {
	return ra.args[key]
}
