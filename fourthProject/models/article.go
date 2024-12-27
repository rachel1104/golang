package models

type Article struct {
	Id      int
	Title   string
	CateId  int
	status  int
	AddTime int
	//ArticleCate ArticleCate `gorm:"foreignkey:CateId"`
}

/*
CREATE TABLE ARTICLE (

	id INT AUTO_INCREMENT PRIMARY KEY,
	title VARCHAR(50) NOT NULL,
	cateid varchar(50) NOT NULL,
	status bool,
	add_time int

)
*/
func (Article) TableName() string {
	return "article"
}
