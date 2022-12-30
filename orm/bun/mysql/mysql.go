package bun_mysql

import (
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
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

	sqldb, err := sql.Open("mysql", "apiuser:WebappLocal@localhost:13306")
	if err != nil {
		panic(err)
	}

	db = bun.NewDB(sqldb, mysqldialect.New())
	return db
}
