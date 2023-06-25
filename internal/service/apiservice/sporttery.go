package apiservice

import (
	"gspider/internal/model"
	"gspider/internal/model/entity"
)

type SportteryService struct {
}

func NewSportteryService() *SportteryService {
	return &SportteryService{}
}

func (a SportteryService) JczqScoreList() []entity.SportteryJczqScore {
	jczqModel := model.NewSportteryJczqScoreModel("default")
	list, _ := jczqModel.GetJczqScoreList()
	return list
}
