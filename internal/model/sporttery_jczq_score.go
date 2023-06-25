package model

import (
	"github.com/kuaileshi1/dbable"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gspider/internal/model/entity"
)

type SportteryJczqScoreModel struct {
	db *gorm.DB
}

func NewSportteryJczqScoreModel(dbName string) *SportteryJczqScoreModel {
	db, err := dbable.GetMysql(dbName)
	if err != nil {
		log.Errorf("SportteryJczqScoreModel: get %s mysql connection failed: %s", dbName, err.Error())
	}

	return &SportteryJczqScoreModel{
		db: db,
	}
}

// BatchInsert
// @Description: 批量插入
// @Auth shigx 2023-04-27 09:29:38
// @param sportteryJczqScore
// @return error
func (m *SportteryJczqScoreModel) BatchInsertOnUpdate(sportteryJczqScore []entity.SportteryJczqScore, updateField []string) error {
	return m.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "match_id"}},
		DoUpdates: clause.AssignmentColumns(updateField),
	}).Create(sportteryJczqScore).Error
}

func (m *SportteryJczqScoreModel) GetJczqScoreList() ([]entity.SportteryJczqScore, error) {
	var list []entity.SportteryJczqScore
	err := m.db.Order("id desc, match_num desc").Limit(50).Find(&list).Error
	return list, err
}
