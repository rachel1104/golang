package models

type Student struct {
	Id          int
	Num         string
	Pwd         string
	ClassId     int
	Studentname string
	Lesson      []Lesson `gorm:"many2many:lesson_student;"`
}

func (Student) TableName() string {
	return "student"
}
