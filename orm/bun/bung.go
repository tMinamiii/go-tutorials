package bun_mysql

import (
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tMinamiii/tutorials/conf"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

var (
	lock sync.Mutex
	db   *bun.DB
)

func GetDB() *bun.DB {
	lock.Lock()
	defer lock.Unlock()
	if db != nil {
		return db
	}

	config := conf.NewDBConfig("webapp")
	sqldb, err := sql.Open("mysql", config.Source())
	if err != nil {
		panic(err)
	}

	db = bun.NewDB(sqldb, mysqldialect.New())
	return db
}
