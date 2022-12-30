// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package sqlc_multi_content

import (
	"context"
	"database/sql"
)

type Querier interface {
	DeleteUser(ctx context.Context, db DBTX, id int64) error
	InsertUser(ctx context.Context, db DBTX, arg InsertUserParams) (sql.Result, error)
	SelectByID(ctx context.Context, db DBTX, id int64) (User, error)
	UpdateUser(ctx context.Context, db DBTX, arg UpdateUserParams) (sql.Result, error)
}

var _ Querier = (*Queries)(nil)
