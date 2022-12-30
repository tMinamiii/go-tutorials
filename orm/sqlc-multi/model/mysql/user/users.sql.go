// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: users.sql

package sqlc_multi_content

import (
	"context"
	"database/sql"
)

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?
`

func (q *Queries) DeleteUser(ctx context.Context, db DBTX, id int64) error {
	_, err := db.ExecContext(ctx, deleteUser, id)
	return err
}

const insertUser = `-- name: InsertUser :execresult
INSERT INTO users (` + "`" + `id` + "`" + `, ` + "`" + `name` + "`" + `, ` + "`" + `nickname` + "`" + `)
    VALUES(?, ?, ?)
`

type InsertUserParams struct {
	ID       int64          `db:"id"`
	Name     string         `db:"name"`
	Nickname sql.NullString `db:"nickname"`
}

func (q *Queries) InsertUser(ctx context.Context, db DBTX, arg InsertUserParams) (sql.Result, error) {
	return db.ExecContext(ctx, insertUser, arg.ID, arg.Name, arg.Nickname)
}

const selectByID = `-- name: SelectByID :one
SELECT id, name, nickname, created_at, updated_at FROM users WHERE id = ?
`

func (q *Queries) SelectByID(ctx context.Context, db DBTX, id int64) (User, error) {
	row := db.QueryRowContext(ctx, selectByID, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Nickname,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const updateUser = `-- name: UpdateUser :execresult
UPDATE users SET ` + "`" + `name` + "`" + ` = ?, ` + "`" + `nickname` + "`" + ` = ?  WHERE id = ?
`

type UpdateUserParams struct {
	Name     string         `db:"name"`
	Nickname sql.NullString `db:"nickname"`
	ID       int64          `db:"id"`
}

func (q *Queries) UpdateUser(ctx context.Context, db DBTX, arg UpdateUserParams) (sql.Result, error) {
	return db.ExecContext(ctx, updateUser, arg.Name, arg.Nickname, arg.ID)
}
