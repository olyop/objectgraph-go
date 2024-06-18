package graphqlops

import (
	"github.com/olyop/graphql-go/server/graphqlops/distributedcache"
	"github.com/olyop/graphql-go/server/graphqlops/inmemorycache"
)

func ClearCache() error {
	err := distributedcache.Clear()
	if err != nil {
		return err
	}

	inmemorycache.Clear()

	return nil
}
