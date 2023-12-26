package orm

import "fmt"

type ConnectOptions struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func (config ConnectOptions) GetDSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.DBName)
}

type Database interface {
	Connect() error
}
