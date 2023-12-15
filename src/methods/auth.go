package methods

import (
	"W-chat/pkg/encrypt"
	"W-chat/src/repository/database"
	"errors"

	"gorm.io/gorm"
)

type AuthMethods struct {
	user *database.Users
}

func NewAuthMethodsObj(user *database.Users) *AuthMethods {
	return &AuthMethods{user: user}
}

// Login 登录处理
func (a *AuthMethods) Login(mobile string, password string) (*database.UsersModel, error) {

	user, err := a.user.FindByMobile(mobile)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("登录账号不存在! ")
		}

		return nil, err
	}

	if !encrypt.VerifyPassword(user.Password, password) {
		return nil, errors.New("登录密码填写错误! ")
	}

	return user, nil
}
