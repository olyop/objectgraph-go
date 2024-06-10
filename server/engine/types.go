package engine

import "context"

type ResolverRequestMutexContextKey struct{}
type ResolverRetrieveMutexContextKey struct{}

type RetrieverMap map[string]Retriever
type Retriever func(ctx context.Context, args RetrieverArgs) (any, error)
type RetrieverArgs map[string]string
