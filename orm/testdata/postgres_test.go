package testdata

import (
	"fmt"
	"github.com/davidkhala/goutils"
	"github.com/davidkhala/goutils/orm"
	"testing"
)

func TestConnect(t *testing.T) {
	var postgres = orm.Postgres{
		ConnectOptions: orm.ConnectOptions{
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
		var data = User{
			Name:  "david",
			Email: "david-khala@hotmail.com",
		}

		err := postgres.Init(&User{})
		goutils.PanicError(err)
		var newCreate = postgres.CreateIfNotExist(&data)
		fmt.Println(newCreate)
		fmt.Println(postgres.Count(&User{}))
		postgres.Truncate(&User{})

	})
}
