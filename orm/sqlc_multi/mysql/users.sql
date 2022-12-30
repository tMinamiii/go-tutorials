CREATE TABLE users (
  `id` BIGINT NOT NULL PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL COMMENT '名前',
  `nickname` VARCHAR(255) DEFAULT NULL COMMENT 'ニックネーム',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- name: SelectByID :one
SELECT * FROM users WHERE id = ?;

-- name: Insert :execresult
INSERT INTO users (`id`, `name`, `nickname`)
    VALUES(?, ?, ?);

-- name: Update :execresult
UPDATE users SET `name` = ?, `nickname` = ?  WHERE id = ?;

-- name: Delete :exec
DELETE FROM users WHERE id = ?;