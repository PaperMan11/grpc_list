CREATE DATABASE IF NOT EXISTS grpc_list;

CREATE TAble IF NOT EXISTS `user` (
    `user_id` VARCHAR(50) PRIMARY KEY,
    `user_name` VARCHAR(100) NOT NULL,
    `user_password` VARCHAR(30) NOT NULL 
)ENGINE=InnoDB DEFAULT CHARSET=utf8;