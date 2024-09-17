package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title    string
	Content  string
	UserID   uint
	User     User
	Comments []Comment
	Tags     []Tag `gorm:"many2many:post_tags"`
}
