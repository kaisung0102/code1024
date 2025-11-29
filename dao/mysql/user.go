package mysql

import (
	"errors"
	"fmt"
	"jachow/code1024/config"
	"jachow/code1024/model"

	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	ErrUserNotExist      = errors.New("用户名不存在")
	ErrInvalidPassword   = errors.New("密码错误")
	ErrUserExist         = errors.New("用户名已存在")
	ErrCommunity         = errors.New("查询社区失败")
	ErrCommunityNotExist = errors.New("社区不存在")
	ErrCreatePostFailed  = errors.New("创建帖子失败")
	ErrPostNotExist      = errors.New("帖子不存在")
)

var DB *gorm.DB

// initDB 初始化数据库连接
func InitMysql() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Conf.Mysql.Username,
		config.Conf.Mysql.Password,
		config.Conf.Mysql.Host,
		config.Conf.Mysql.Port,
		config.Conf.Mysql.Database,
	)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.L().Error("mysql connect failed", zap.Error(err))
		panic(fmt.Sprintf("数据库连接失败: %v", err))
	}
	// TODO: 自动迁移数据库表结构
	// err = DB.AutoMigrate(&model.User{})
	// if err != nil {
	// 	zap.L().Error("mysql auto migrate failed", zap.Error(err))
	// 	panic("mysql自动迁移表结构失败")
	// }
	// TODO: 打印数据库连接成功日志
	zap.L().Info("mysql connect success")
	return err
}

func QueryUser(username string) (b bool) {
	var cnt int64
	err := DB.Model(&model.User{}).Where("username = ?", username).Count(&cnt).Error
	if err != nil {
		zap.L().Error("mysql query user failed", zap.Error(err))
		return false
	}
	return cnt == 1
}

func GetUserByID(userID int64) (user *model.User, err error) {
	err = DB.Model(&model.User{}).Where("user_id = ?", userID).First(&user).Error
	if err != nil {
		zap.L().Error("mysql get user failed", zap.Error(err))
		return nil, ErrUserNotExist
	}
	return user, nil
}

func CreateUser(user *model.User) (err error) {
	user.Password = EncryptPassword(user.Password)
	err = DB.Create(user).Error
	if err != nil {
		zap.L().Error("mysql create user failed", zap.Error(err))
		return ErrUserExist
	}
	return err
}

func EncryptPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		zap.L().Error("mysql encrypt password failed", zap.Error(err))
		panic("passwd加密失败")
	}
	return string(hash)
}

func CheckPassword(username, password string) bool {
	var user model.User
	err := DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		zap.L().Error("mysql check password failed", zap.Error(err))
		return false
	}
	return bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) == nil
}

func Login(user *model.User) (err error) {
	err = DB.Model(&model.User{}).Where("username = ?", user.Username).First(&user).Error
	if err != nil {
		zap.L().Error("mysql login failed", zap.Error(err))
		return ErrUserNotExist
	}

	return nil
}
