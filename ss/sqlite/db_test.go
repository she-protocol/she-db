//go:build sqliteBackend
// +build sqliteBackend

package sqlite

import (
	"testing"

	"github.com/she-protocol/she-db/config"
	sstest "github.com/she-protocol/she-db/ss/test"
	"github.com/she-protocol/she-db/ss/types"
	"github.com/stretchr/testify/suite"
)

// TODO: Update Sqlite to latest
func TestStorageTestSuite(t *testing.T) {
	s := &sstest.StorageTestSuite{
		NewDB: func(dir string) (types.StateStore, error) {
			return New(dir, config.DefaultStateStoreConfig())
		},
		EmptyBatchSize: 0,
	}

	suite.Run(t, s)
}
