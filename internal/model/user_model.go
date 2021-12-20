// @Title 请填写文件名称（需要改）
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/11/26 6:36 下午
package model

import (
	"github.com/kuaileshi1/dbable"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gspider/internal/model/entity"
)

// 用户表模型
type UserModel struct {
	db *gorm.DB
}

// @Description 实例化UserModel
// @Auth shigx
// @Date 2021/11/26 6:45 下午
// @param
// @return
func NewUserModel(dbName string) *UserModel {
	db, err := dbable.GetMysql(dbName)
	if err != nil {
		log.Errorf("UserModel: get %s mysql connection failed: %s", dbName, err.Error())
	}

	return &UserModel{
		db: db,
	}
}

// @Description 校验用户账号密码是否正确
// @Auth shigx
// @Date 2021/12/13 9:50 上午
// @param
// @return
func (m *UserModel) IsValidUser(username, password string) (*entity.User, error) {
	user := &entity.User{}
	if err := m.db.Where("username = ?", username).Take(user).Error; err != nil {
		return nil, err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, err
	}

	return user, nil
}
