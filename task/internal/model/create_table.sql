

CREATE TABLE IF NOT EXISTS `task` (
    `task_id` BIGINT PRIMARY KEY,
    `user_id` BIGINT UNIQUE NOT NULL,
    `status` INTEGER UNSIGNED NOT NULL,
    `title` VARCHAR(255) NOT NULL,
    `content` LONGTEXT NOT NULL,
    `start_time` INTEGER UNSIGNED,
    `end_time` INTEGER UNSIGNED,
    CONSTRAINT task_fk FOREIGN KEY (user_id) REFERENCES `user`(user_id) 
    ON DELETE CASCADE           
)ENGINE=InnoDB DEFAULT CHARSET=utf8;

// ON DELETE CASCADE: 当父表删除时, 子表对应的记录也删除