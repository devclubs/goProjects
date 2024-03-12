package gorm_project

import (
	"fmt"
	"testing"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 连接到数据库
func CreateDBLink() *gorm.DB {
	dsn := "root:rootpwd@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败：", err.Error())
	}

	return db
}

// 声明模型
type User struct {
	gorm.Model
	Name     string    `json:"name" gorm:"size=32"`
	Age      uint8     `json:"age`
	Email    string    `json:"email" gorm:"unique"`
	Birthday time.Time `json:"birthday" gorm:"column:birthday"`
}

func (User) TableName() string {
	return "t_user"
}

// 迁移数据库
func TestMigrateDB(*testing.T) {
	db := CreateDBLink()
	db.AutoMigrate(&User{})
}

// 插入一条记录
func TestInsertOne(t *testing.T) {
	var u = User{Name: "亚瑟", Age: 5, Email: "yase@qq.com", Birthday: time.Now()}

	db := CreateDBLink()
	db.Create(&u)
}

func TestInsertMany(t *testing.T) {
	var mu []User = []User{
		{Name: "豹女", Age: 3, Email: "baonv@qq.com", Birthday: time.Now()},
		{Name: "奶妈", Age: 4, Email: "naima@qq.com", Birthday: time.Now()},
	}

	db := CreateDBLink()
	db.Create(&mu)
}

func TestSelectOne(t *testing.T) {
	var u User
	db := CreateDBLink()
	//  获取第一条记录（主键升序）
	// db.First(&u)

	// 获取一条记录，没有指定排序字段
	db.Take(&u)
	fmt.Println(u)

	// 获取最后一条记录（主键降序）
	db.Last(&u)
	fmt.Println(u)

	result := db.First(&u)
	fmt.Println(result.RowsAffected, result.Error)
}
