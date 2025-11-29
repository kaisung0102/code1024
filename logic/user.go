package logic

import (
	"jachow/code1024/dao/mysql"
	"jachow/code1024/model"
	"jachow/code1024/pkg"
)

func SignUp(params *model.ParamsSignUp) (err error) {
	// 检查用户名是否存在
	if mysql.QueryUser(params.Username) {
		return mysql.ErrUserExist
	}

	// 创建用户
	uid := pkg.GetID()
	newUser := &model.User{
		UserID:   uid,
		Username: params.Username,
		Password: params.Password,
	}

	return mysql.CreateUser(newUser)
}

func Login(params *model.ParamsLogin) (token string, err error) {
	// 检查用户名是否存在
	if !mysql.QueryUser(params.Username) {
		return "", mysql.ErrUserNotExist
	}

	// 检查密码是否正确
	if !mysql.CheckPassword(params.Username, params.Password) {
		return "", mysql.ErrInvalidPassword
	}

	user := &model.User{
		Username: params.Username,
		Password: params.Password,
	}

	if err = mysql.Login(user); err != nil {
		return "", mysql.ErrInvalidPassword
	}
	// 怎么获得userid

	token, err = pkg.GenToken(user.UserID, user.Username)
	if err != nil {
		return "", err
	}

	return token, nil
}



