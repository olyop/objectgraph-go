package engine

type ResolverRequestMutexContextKey struct{}
type ResolverRetrieveMutexContextKey struct{}

type RetrieverMap map[string]Retriever
type Retriever func(args RetrieverArgs) (any, error)
type RetrieverArgs map[string]string
