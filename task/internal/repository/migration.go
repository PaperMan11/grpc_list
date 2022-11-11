package repository

import "log"

const taskSchema = `CREATE TABLE IF NOT EXISTS task (
    task_id BIGINT PRIMARY KEY,
    user_id BIGINT UNIQUE NOT NULL,
    status INTEGER UNSIGNED NOT NULL,
    title VARCHAR(255) NOT NULL,
    content LONGTEXT NOT NULL,
    start_time INTEGER UNSIGNED,
    end_time INTEGER UNSIGNED,
    CONSTRAINT task_fk FOREIGN KEY (user_id) REFERENCES user(user_id) 
    ON DELETE CASCADE           
)ENGINE=InnoDB DEFAULT CHARSET=utf8;`

func migration() error {
	if _, err := Db.Exec(taskSchema); err != nil {
		log.Println("create table task failed")
		return err
	}
	return nil
}
