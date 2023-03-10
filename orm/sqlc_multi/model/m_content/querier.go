// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package m_content

import (
	"context"
	"database/sql"
)

type Querier interface {
	Delete(ctx context.Context, db DBTX, id int64) error
	Insert(ctx context.Context, db DBTX, arg InsertParams) (sql.Result, error)
	SelectByID(ctx context.Context, db DBTX, id int64) (Content, error)
	SelectByUserID(ctx context.Context, db DBTX, userID int64) ([]Content, error)
	Update(ctx context.Context, db DBTX, arg UpdateParams) (sql.Result, error)
}

var _ Querier = (*Queries)(nil)
