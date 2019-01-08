package leveldb

import (
	. "github.com/davidkhala/goutils"
	"github.com/syndtr/goleveldb/leveldb"
)

func Connect(path string) *leveldb.DB {
	db, err := leveldb.OpenFile(path, nil)
	PanicError(err)
	return db
}
func Close(db *leveldb.DB) {
	var err = db.Close()
	PanicError(err)
}
