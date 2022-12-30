package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/tMinamiii/tutorials/orm/sqlc_multi/db/mysql"
	"github.com/tMinamiii/tutorials/orm/sqlc_multi/model/m_content"
	"github.com/tMinamiii/tutorials/orm/sqlc_multi/model/m_user"
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

	mu := m_user.New()
	u, err := mu.SelectByID(ctx, tx, 1)
	if err != nil && err.Error() != "sql: no rows in result set" {
		log.Fatalf("type %T, msg = %s", err, err.Error())
	}
	fmt.Printf("%+v\n", u)

	mc := m_content.New()
	contents, err := mc.SelectByUserID(ctx, tx, u.ID)
	if err != nil {
		log.Fatalf("type %T, msg = %s", err, err.Error())
	}

	for _, v := range contents {
		fmt.Println(v.Content)
	}

	err = tx.Commit()
	if err != nil {
		log.Fatalf("type %T, msg = %s", err, err.Error())
	}
}
