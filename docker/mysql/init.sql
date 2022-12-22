CREATE DATABASE IF NOT EXISTS `gobase-backend_local`;
GRANT ALL ON `gobase-backend_local`.* TO 'testuser'@'%';

CREATE DATABASE IF NOT EXISTS `gobase-backend_test`;
GRANT ALL ON `gobase-backend_test`.* TO 'testuser'@'%';
