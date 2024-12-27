package models

type User struct {
	Id       int
	Username string
	Pwd      string
	Age      int
	Email    string
	AddTime  int
}

/*
CREATE TABLE users (

	id INT AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(50) NOT NULL,
	pwd varchar(50) NOT NULL,
	age INT,
	email VARCHAR(100) UNIQUE,
	add_time int

)
*/
func (User) TableName() string {
	return "user"
}
