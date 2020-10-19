package models

import DB "gin-web/database"
import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name          string `gorm:"unique;size:100" form:"name"`
	Password      string `gorm:"index:idx_password;size:255" form:"password"`
	PasswordAgain string `gorm:"-" form:"password_again" binding:"eqfield=Password"`
	Email         string `gorm:"index:idx_email;size:255" form:"email" binding:"required"`
}

func init() {
	tableExt := DB.Conn.Migrator().HasTable(&User{})
	if !tableExt {
		DB.Conn.AutoMigrate(&User{})
	}
}

func (u *User) Create() (uint, error) {
	res := DB.Conn.Create(&u)
	return uint(u.ID), res.Error
}

func (User) GetUserByName(username string) User {
	user := User{}
	DB.Conn.Where("name = ?", username).First(&user)
	return user
}

func (U *User) LoginVerify(user, password string) {

}
