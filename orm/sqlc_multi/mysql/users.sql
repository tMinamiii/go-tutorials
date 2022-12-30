CREATE TABLE users (
  `id` BIGINT NOT NULL PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL COMMENT '名前',
  `nickname` VARCHAR(255) DEFAULT NULL COMMENT 'ニックネーム',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- name: SelectByID :one
SELECT * FROM users WHERE id = ?;

-- name: InsertUser :execresult
INSERT INTO users (`id`, `name`, `nickname`)
    VALUES(?, ?, ?);

-- name: UpdateUser :execresult
UPDATE users SET `name` = ?, `nickname` = ?  WHERE id = ?;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?;