-- Active: 1741074812216@@127.0.0.1@3306@gorm
/* type Users struct {
	Id       int
	UserName string
	Age      int
	Email    string
	AddTime  int
}
 */

 /* create table */
CREATE TABLE IF NOT EXISTS users (
    id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255) NOT NULL,
    age TINYINT,
    email VARCHAR(255),
    add_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

/* insert test value */
INSERT into users (username,age,email)
VALUES ('Alice', 25, 'alice@example.com'),
        ('Bob', 30, 'bob@example.com'),
        ('Charlie', 22, 'charlie@example.com');