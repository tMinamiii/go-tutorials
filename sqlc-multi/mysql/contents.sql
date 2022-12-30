CREATE TABLE contents (
  `id` BIGINT NOT NULL PRIMARY KEY,
  `user_id` BIGINT NOT NULL,
  `content` VARCHAR(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

-- name: SelectContentByID :one
SELECT * FROM contents WHERE id = ?;

-- name: InsertContent :execresult
INSERT INTO contents (id, user_id, content)
    VALUES(?, ?, ?);

-- name: UpdateContent :execresult
UPDATE contents SET content = ? WHERE id = ?;

-- name: DeleteContent :exec
DELETE FROM contents WHERE id = ?;