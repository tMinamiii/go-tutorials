package sqlc_single

import (
	"database/sql"
	"log"
	"sync"
	"time"
	"tutorials/conf"

	_ "github.com/go-sql-driver/mysql"
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
	err := tx.Rollback()
	if err != nil {
		log.Fatalf("failed to MySQL Rollback: %s", err.Error())
	}
}
