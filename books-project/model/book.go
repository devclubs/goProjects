package model

type Book struct {
	Id   int64  `json:"id" gorm:"primaryKey"`
	Name string `json:"name" binding:"required" gorm:"uniq" `
	Desc string `json:"desc"`
	User []User `gorm:"many2many:book_user"`
}

func (Book) TableName() string {
	return "b_book"
}
