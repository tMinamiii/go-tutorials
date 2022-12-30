package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tMinamiii/tutorials/orm/sqlc_multi/db/mysql"
	"github.com/tMinamiii/tutorials/orm/sqlc_multi/model/muser"
)

func main() {
	os.Setenv("GO_ENV", "local")
	ctx := context.Background()

	db := mysql.GetDB("webapp")
	tx, err := db.Begin()
	if err != nil {
		log.Fatalf("type %T, msg = %s", err, err.Error())
	}
	defer mysql.RollbackUnlessCommit(tx)

	mUserQ := muser.New()
	u, err := mUserQ.SelectByID(ctx, tx, 1)
	if err != nil && err.Error() != "sql: no rows in result set" {
		log.Fatalf("type %T, msg = %s", err, err.Error())
	}
	fmt.Printf("%+v\n", u)

	err = tx.Commit()
	if err != nil {
		log.Fatalf("type %T, msg = %s", err, err.Error())
	}
}
