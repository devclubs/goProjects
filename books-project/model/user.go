package model

type User struct {
	Id       int64  `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"not null" binding:"required"`
	Password string `json:"password" gorm:"not null" binding:"required"`
	Token    string `json:"token"`
}

func (User) TableName() string {
	return "b_user"
}
