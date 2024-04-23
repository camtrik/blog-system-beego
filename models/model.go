package models

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var validate = validator.New()

type Blog struct {
	Id      int `PK` // 设置Id为主键
	Title   string
	Content string
	Created time.Time
}

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100);uniqueIndex"` // 设置为 VARCHAR 并且唯一
	Password string
	Email    string
}

func (user *User) Validate() error {
	return validate.Struct(user)
}

var DB *gorm.DB

func Init() {
	var err error
	// db, err := gorm.Open("mysql", "root:root@/test?charset=utf8&parseTime=True&loc=Local")
	dsn := "ebbi:ht1234@tcp(localhost:3306)/blog_test?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	DB.AutoMigrate(&Blog{})
	DB.AutoMigrate(&User{})
}

func GetAll() ([]Blog, error) {
	var blogs []Blog
	result := DB.Find(&blogs)
	return blogs, result.Error
}

func GetBlog(id int) (Blog, error) {
	var blog Blog
	result := DB.First(&blog, id) // 查找第一个符合条件的数据
	return blog, result.Error
}

func SaveBlog(blog *Blog) error {
	if blog.Id == 0 {
		result := DB.Create(blog)
		return result.Error
	}

	// 否则，更新现有记录
	result := DB.Save(blog)
	return result.Error
}

func DelBlog(blog *Blog) error {
	result := DB.Delete(blog)
	return result.Error
}

func UserInsert(user *User) error {
	result := DB.Create(user)
	return result.Error
}

func CreateUser(user *User) error {
	result := DB.Create(&user)
	return result.Error
}

func GetUser(username string) (User, error) {
	var user User
	result := DB.Where("username=?", username).First(&user)
	return user, result.Error
}
