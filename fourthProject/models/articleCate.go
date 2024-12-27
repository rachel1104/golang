package models

// foreignKey外鍵
// references主鍵 ，默認id
type ArticleCate struct {
	Id      int
	Title   string
	Status  int
	Article []Article `gorm:"foreignkey:CateId"`
	//Article []Article `gorm:"foreignkey:ArticleCateId"`
}

func (ArticleCate) TableName() string {
	return "article_cate"
}
