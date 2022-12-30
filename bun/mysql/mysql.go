package mysql

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func Init() *bun.DB {
	sqldb, err := sql.Open("mysql", "root@localhost:13306")
	if err != nil {
		panic(err)
	}

	return bun.NewDB(sqldb, mysqldialect.New())
}
