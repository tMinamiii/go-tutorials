package mysql

import (
	"database/sql"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/tMinamiii/tutorials/conf"
)

var (
	lock sync.Mutex
	sess *sql.DB
)

func GetDB(database string) *sql.DB {
	lock.Lock()
	defer lock.Unlock()
	if sess != nil {
		return sess
	}

	config := conf.NewDBConfig("webapp")
	sess, err := sql.Open("mysql", config.Source())
	if err != nil {
		panic(err)
	}

	sess.SetMaxOpenConns(15)
	sess.SetConnMaxIdleTime(10 * time.Second)
	sess.SetConnMaxLifetime(1 * time.Minute)
	return sess
}

// RollbackUnlessCommit トランザクションがcommitされていなかった場合 rollback する
func RollbackUnlessCommit(tx *sql.Tx) {
	_ = tx.Rollback()
}
