package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DB *gorm.DB

type Blog struct {
	Id      int `PK`
	Title   string
	Content string
	Created time.Time
}

func GetLink() (err error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/beeblog?charset=utf8&parseTime=True&loc=Local"

	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	return err
}

func GetAll() (blogs []Blog) {
	DB.Find(&blogs)
	return
}

func GetBlog(id int) (blog Blog) {
	DB.Where("id=?", id).Find(&blog)
	return
}

func SaveBlog(blog Blog) (bg Blog) {
	DB.Save(&blog)
	return bg
}

func DelBlog(blog Blog) {
	DB.Delete(&blog)
	return
}