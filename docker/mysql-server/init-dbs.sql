CREATE USER 'apiuser'@'%' IDENTIFIED BY 'WebappLocal';

CREATE DATABASE IF NOT EXISTS webapp_local CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
GRANT ALL ON `webapp_local`.* TO 'apiuser'@'%';

CREATE DATABASE IF NOT EXISTS webapp_test CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
GRANT ALL ON `webapp_test_1`.* TO 'apiuser'@'%';

FLUSH PRIVILEGES;