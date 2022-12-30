package sqlc_multi

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type databaseConfig struct {
	database string
	user     string
	passwd   string
	address  string
}

var webappConfig = &databaseConfig{
	database: databaseWithEnv("webapp"),
	user:     "apiuser",
	passwd:   "WebappLocal",
	address:  "localhost:13306",
}

func (d *databaseConfig) source() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", d.user, d.passwd, d.address, d.database)
}

func databaseWithEnv(database string) string {
	goenv := os.Getenv("GO_ENV")
	return fmt.Sprintf("%s_%s", database, goenv)
}

var (
	lock sync.Mutex
	sess *sql.DB
)

func GetSession(database string) *sql.DB {
	lock.Lock()
	defer lock.Unlock()
	if sess != nil {
		return sess
	}

	sess, err := sql.Open("mysql", webappConfig.source())
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
