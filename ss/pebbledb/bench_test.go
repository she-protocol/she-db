package pebbledb

import (
	"testing"

	"github.com/she-protocol/she-db/config"
	sstest "github.com/she-protocol/she-db/ss/test"
	"github.com/she-protocol/she-db/ss/types"
)

func BenchmarkDBBackend(b *testing.B) {
	s := &sstest.StorageBenchSuite{
		NewDB: func(dir string) (types.StateStore, error) {
			return New(dir, config.DefaultStateStoreConfig())
		},
		BenchBackendName: "PebbleDB",
	}

	s.BenchmarkGet(b)
	s.BenchmarkApplyChangeset(b)
	s.BenchmarkIterate(b)
}
