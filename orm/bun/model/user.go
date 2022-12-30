package bun_model

import (
	"database/sql"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`
	ID            int64          `bun:"id,pk"`
	Name          string         `bun:"name,notnull"`
	Nickname      sql.NullString `bun:"nickname"`
	Timestamp
}
