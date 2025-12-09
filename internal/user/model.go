package user

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"index"` // for search
	Password string
	Name     string
}
