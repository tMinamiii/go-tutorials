package bun_model

import "github.com/uptrace/bun"

type Content struct {
	bun.BaseModel `bun:"table:contents"`
	ID            int64  `bun:"id,pk"`
	UserID        int64  `bun:"user_id,notnull"`
	Content       string `bun:"content,notnull"`
	Timestamp
}
