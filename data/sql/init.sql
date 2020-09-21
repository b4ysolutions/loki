CREATE DATABASE lokiapidb;
USE lokiapidb;
CREATE USER 'lokiapiuser'@'%' IDENTIFIED BY 'lokiapipass';
CREATE TABLE `users` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(100) NOT NULL,
  `user_email` varchar(100) NOT NULL,
  `user_password` varchar(100) NOT NULL,
  `user_status` int(11) NOT NULL,
  `user_islogged` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `user_email_UNIQUE` (`user_email`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;
GRANT ALL PRIVILEGES ON `lokiapidb`.* TO `lokiapiuser`@`%`;
FLUSH PRIVILEGES;