package model

import (
	"errors"

	"github.com/zjwdmlmx/ssensor/global"
)

type User struct {
	Uid string `gorm:"type:char(64);primary_key"`
}

type userModel struct{}

var (
	ErrUserExists error = errors.New("User exists")
)

var (
	UserModel userModel
)

func (userModel) CreateOne(row *User) (err error) {
	var u User
	global.DB.Where("uid = ?", row.Uid).First(&u)
	if u.Uid != "" {
		err = ErrUserExists
		return
	}

	global.DB.Create(row)
	return
}
