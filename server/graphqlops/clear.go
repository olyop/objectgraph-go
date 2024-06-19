package graphqlops

import (
	"github.com/olyop/graphqlops-go/graphqlops/distributedcache"
	"github.com/olyop/graphqlops-go/graphqlops/inmemorycache"
)

func ClearCache() error {
	err := distributedcache.Clear()
	if err != nil {
		return err
	}

	inmemorycache.Clear()

	return nil
}
