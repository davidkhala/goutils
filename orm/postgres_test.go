package orm

import (
	"github.com/davidkhala/goutils"
	"github.com/davidkhala/goutils/orm/testdata"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConnect(t *testing.T) {
	var postgres = Postgres{
		ConnectOptions: ConnectOptions{
			Host:     "localhost",
			Port:     5432,
			User:     "postgres",
			Password: "example",
			DBName:   "postgres",
		},
	}
	err := postgres.Connect()
	goutils.PanicError(err)

	t.Run("use table", func(t *testing.T) {
		var data = testdata.User{
			Name:  "david",
			Email: "david-khala@hotmail.com",
		}

		err := postgres.Init(&testdata.User{})
		goutils.PanicError(err)
		var newCreate = postgres.CreateIfNotExist(&data)
		assert.True(t, newCreate)
		assert.True(t, postgres.Count(&testdata.User{}) > 0)
		postgres.Truncate(&testdata.User{})

	})
}
