// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package mysql_content

import (
	"context"
	"database/sql"
)

type Querier interface {
	DeleteContent(ctx context.Context, db DBTX, id int64) error
	InsertContent(ctx context.Context, db DBTX, arg InsertContentParams) (sql.Result, error)
	SelectContentByID(ctx context.Context, db DBTX, id int64) (Content, error)
	UpdateContent(ctx context.Context, db DBTX, arg UpdateContentParams) (sql.Result, error)
}

var _ Querier = (*Queries)(nil)