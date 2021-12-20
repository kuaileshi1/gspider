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

// 任务表模型
type TaskModel struct {
	db *gorm.DB
}

// @Description 实例化TaskModel
// @Auth shigx
// @Date 2021/11/26 6:45 下午
// @param
// @return
func NewTaskModel(dbName string) *TaskModel {
	db, err := dbable.GetMysql(dbName)
	if err != nil {
		log.Errorf("TaskModel: get %s mysql connection failed: %s", dbName, err.Error())
	}

	return &TaskModel{
		db: db,
	}
}

// @Description 查询所有任务
// @Auth shigx
// @Date 2021/11/26 11:33 下午
// @param
// @return
func (m *TaskModel) GetAll() ([]entity.Task, error) {
	var res []entity.Task
	if err := m.db.Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

// @Description 分页查询任务
// @Auth shigx
// @Date 2021/12/13 11:53 下午
// @param
// @return
func (m *TaskModel) GetList(size int, offset int) ([]entity.Task, error) {
	var res []entity.Task
	if err := m.db.Order("id desc").Limit(size).Offset(offset).Find(&res).Error; err != nil {
		return nil, err
	}

	return res, nil
}

// @Description 查询任务数量
// @Auth shigx
// @Date 2021/12/13 11:50 下午
// @param
// @return
func (m *TaskModel) Count() (total int64, err error) {
	err = m.db.Model(&entity.Task{}).Count(&total).Error
	return
}

// @Description 根据ID返回一条记录
// @Auth shigx
// @Date 2021/11/30 10:57 下午
// @param
// @return
func (m *TaskModel) GetOneById(id int) (ret *entity.Task, err error) {
	err = m.db.Where("id = ?", id).Take(&ret).Error
	return
}

// @Description 更新任务
// @Auth shigx
// @Date 2021/11/26 11:50 下午
// @param
// @return
func (m *TaskModel) Updates(query string, queryVal []interface{}, update map[string]interface{}) error {
	return m.db.Model(&entity.Task{}).Where(query, queryVal...).Updates(update).Error
}

// @Description 保存数据
// @Auth shigx
// @Date 2021/12/1 10:16 上午
// @param
// @return
func (m *TaskModel) Save(task *entity.Task) error {
	return m.db.Save(task).Error
}

// @Description 删除任务
// @Auth shigx
// @Date 2021/12/20 10:39 上午
// @param
// @return
func (m *TaskModel) Delete(id int) error {
	return m.db.Delete(&entity.Task{ID: id}).Error
}
