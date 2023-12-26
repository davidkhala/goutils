package orm

import (
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)
import "gorm.io/driver/postgres"

type Postgres struct {
	ConnectOptions
	*gorm.DB
}

func (p *Postgres) Connect() (err error) {

	var dbConfig = gorm.Config{
		NamingStrategy:    schema.NamingStrategy{SingularTable: true},
		AllowGlobalUpdate: true,
	}
	db, err := gorm.Open(postgres.Open(p.ConnectOptions.GetDSN()),
		&dbConfig)
	p.DB = db
	return

}
func (p Postgres) Init(models ...interface{}) error {
	return p.DB.AutoMigrate(models...)
}
func (p Postgres) Truncate(model interface{}) {
	p.DB.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(model)
}
func (p Postgres) Count(model interface{}) (count int64) {
	p.DB.Model(model).Count(&count)
	return
}
func (p Postgres) CreateIfNotExist(model interface{}) bool {
	result := p.DB.FirstOrCreate(model)
	return result.RowsAffected == 1
}
