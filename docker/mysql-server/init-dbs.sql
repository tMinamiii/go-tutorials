CREATE USER 'apiuser'@'%' IDENTIFIED BY 'WebappLocal';

CREATE DATABASE IF NOT EXISTS webapp_local CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
GRANT ALL ON `webapp_local`.* TO 'apiuser'@'%';

CREATE DATABASE IF NOT EXISTS webapp_test CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
GRANT ALL ON `webapp_test_1`.* TO 'apiuser'@'%';

USE webapp_local

CREATE TABLE IF NOT EXISTS users (
  `id` BIGINT NOT NULL PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL COMMENT '名前',
  `nickname` VARCHAR(255) DEFAULT NULL COMMENT 'ニックネーム',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  Index `index_id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE IF NOT EXISTS contents (
  `id` BIGINT NOT NULL PRIMARY KEY,
  `user_id` BIGINT NOT NULL,
  `content` VARCHAR(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
  Index `index_id` (`id`)
  Index `index_user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

USE webapp_test

CREATE TABLE IF NOT EXISTS users (
  `id` BIGINT NOT NULL PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL COMMENT '名前',
  `nickname` VARCHAR(255) DEFAULT NULL COMMENT 'ニックネーム',
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE IF NOT EXISTS contents (
  `id` BIGINT NOT NULL PRIMARY KEY,
  `user_id` BIGINT NOT NULL,
  `content` VARCHAR(255) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


FLUSH PRIVILEGES;