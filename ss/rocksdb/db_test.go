//go:build rocksdbBackend
// +build rocksdbBackend

package rocksdb

import (
	"testing"

	"github.com/she-protocol/she-db/config"
	sstest "github.com/she-protocol/she-db/ss/test"
	"github.com/she-protocol/she-db/ss/types"
	"github.com/stretchr/testify/suite"
)

func TestStorageTestSuite(t *testing.T) {
	s := &sstest.StorageTestSuite{
		NewDB: func(dir string) (types.StateStore, error) {
			return New(dir, config.DefaultStateStoreConfig())
		},
		EmptyBatchSize: 12,
	}

	suite.Run(t, s)
}
