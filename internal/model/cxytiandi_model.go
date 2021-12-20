// @Title 请填写文件名称（需要改）
// @Description 请填写文件描述（需要改）
// @Author shigx 2021/11/26 6:36 下午
package model

import (
	"github.com/kuaileshi1/dbable"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gspider/internal/model/entity"
)

// 猿天地文章表模型
type CxytiandiModel struct {
	db *gorm.DB
}

// @Description 实例化TaskModel
// @Auth shigx
// @Date 2021/11/26 6:45 下午
// @param
// @return
func NewCxytiandiModel(dbName string) *CxytiandiModel {
	db, err := dbable.GetMysql(dbName)
	if err != nil {
		log.Errorf("CxytiandiModel: get %s mysql connection failed: %s", dbName, err.Error())
	}

	return &CxytiandiModel{
		db: db,
	}
}

// @Description 批量插入
// @Auth shigx
// @Date 2021/12/1 10:16 上午
// @param
// @return
func (m *CxytiandiModel) BatchInsert(cxytiandi []entity.Cxytiandi) error {
	return m.db.Create(cxytiandi).Error
}
