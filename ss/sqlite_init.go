//go:build sqliteBackend
// +build sqliteBackend

package ss

import (
	"github.com/she-protocol/she-db/common/utils"
	"github.com/she-protocol/she-db/config"
	"github.com/she-protocol/she-db/ss/sqlite"
	"github.com/she-protocol/she-db/ss/types"
)

func init() {
	initializer := func(dir string, configs config.StateStoreConfig) (types.StateStore, error) {
		dbHome := utils.GetStateStorePath(dir, configs.Backend)
		if configs.DBDirectory != "" {
			dbHome = configs.DBDirectory
		}
		return sqlite.New(dbHome, configs)
	}
	RegisterBackend(SQLiteBackend, initializer)
}
