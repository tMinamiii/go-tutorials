package conf

import (
	"fmt"
	"os"
)

type DatabaseConfig struct {
	database string
	user     string
	passwd   string
	address  string
}

func NewDBConfig(dbname string) *DatabaseConfig {
	e := os.Getenv("GO_ENV")
	database := fmt.Sprintf("%s_%s", dbname, e)
	return &DatabaseConfig{
		database: database,
		user:     "apiuser",
		passwd:   "WebappLocal",
		address:  "localhost:13306",
	}
}

func (d *DatabaseConfig) Source() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", d.user, d.passwd, d.address, d.database)
}
