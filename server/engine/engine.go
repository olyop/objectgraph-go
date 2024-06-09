package engine

import (
	"time"

	"github.com/olyop/graphql-go/server/engine/distributedcache"
)

var retrievers RetrieverMap
var cacheDurations map[string]time.Duration

func Initialize() {
	distributedcache.Connect()
}

func RegisterRetrievers(r RetrieverMap) {
	retrievers = r
}

func RegisterCacheDurations(d map[string]time.Duration) {
	cacheDurations = d
}

func Close() {
	distributedcache.Close()
}
