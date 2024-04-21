package models

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Blog struct {
	Id      int `PK` // 设置Id为主键
	Title   string
	Content string
	Created time.Time
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
