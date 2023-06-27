// Package sporttery
// @Description: 竞彩网开奖结果抓取
// @Auth shigx 2023-04-24 16:58:25
package sporttery

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
	"gspider/internal/model"
	"gspider/internal/model/entity"
	"gspider/internal/pkg/spider"
	"strconv"
	"time"
)

var ruleScore = &spider.TaskRule{
	Name:        "竞彩网足球赛果",
	Description: "竞彩足球赛果信息抓取",
	Rule: &spider.Rule{
		Url: getJczqResultUrl(),
		Nodes: map[int]*spider.Node{
			0: stepScore1,
		},
	},
}

func getJczqResultUrl() string {
	searchDate := time.Now().Format("2006-01-02")
	return "https://webapi.sporttery.cn/gateway/jc/football/getMatchResultV1.qry?matchPage=1&matchBeginDate=" + searchDate + "&matchEndDate=" + searchDate + "&leagueId=&pageSize=200&pageNo=1&isFix=0&pcOrWap=1"
}

// 页面结果
type scoreResult struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	Value        struct {
		Total       int `json:"total"`
		MatchResult []struct {
			A                 string `json:"a"`
			AllAwayTeam       string `json:"allAwayTeam"`
			AllHomeTeam       string `json:"allHomeTeam"`
			AwayTeam          string `json:"awayTeam"`
			AwayTeamId        int    `json:"awayTeamId"`
			BettingSingle     int    `json:"bettingSingle"`
			D                 string `json:"d"`
			GoalLine          string `json:"goalLine"`
			H                 string `json:"h"`
			HomeTeam          string `json:"homeTeam"`
			HomeTeamId        int    `json:"homeTeamId"`
			LeagueBackColor   string `json:"leagueBackColor"`
			LeagueId          int    `json:"leagueId"`
			LeagueName        string `json:"leagueName"`
			LeagueNameAbbr    string `json:"leagueNameAbbr"`
			MatchDate         string `json:"matchDate"`
			MatchId           int    `json:"matchId"`
			MatchNum          string `json:"matchNum"`
			MatchNumStr       string `json:"matchNumStr"`
			MatchResultStatus string `json:"matchResultStatus"`
			PoolStatus        string `json:"poolStatus"`
			SectionsNo1       string `json:"sectionsNo1"`
			SectionsNo999     string `json:"sectionsNo999"`
			WinFlag           string `json:"winFlag"`
		} `json:"matchResult"`
	} `json:"value"`
}

var stepScore1 = &spider.Node{
	OnRequest: func(req *colly.Request) {
		log.Infof("Visiting %s", req.URL.String())
	},
	OnError: func(res *colly.Response, err error) error {
		log.Errorf("Visiting failed! url:%s, err:%s", res.Request.URL.String(), err.Error())
		// 出错时重试三次
		return Retry(res, 3)
	},
	OnResponse: func(res *colly.Response, nextC *colly.Collector) error {
		if res.StatusCode != 200 {
			return fmt.Errorf("Response status:%d", res.StatusCode)
		}

		var response scoreResult
		err := json.Unmarshal(res.Body, &response)
		if err != nil {
			return err
		}
		if response.ErrorCode == "" {
			return fmt.Errorf("json.Unmarshal error, body: %v", string(res.Body))
		}
		if response.ErrorCode != "0" {
			return fmt.Errorf("errorCode:%v, errorMessage:%v", response.ErrorCode, response.ErrorMessage)
		}
		if response.Value.Total == 0 || len(response.Value.MatchResult) == 0 {
			return nil
		}
		scoreModel := model.NewSportteryJczqScoreModel("default")
		insertData := make([]entity.SportteryJczqScore, 0)
		for _, v := range response.Value.MatchResult {
			matchNum, _ := strconv.Atoi(v.MatchNum)
			goalLine, _ := strconv.Atoi(v.GoalLine)
			matchStatus, _ := strconv.Atoi(v.MatchResultStatus)
			insertData = append(insertData, entity.SportteryJczqScore{
				MatchId:     v.MatchId,
				MatchNum:    matchNum,
				MatchName:   v.MatchNumStr,
				MatchDate:   v.MatchDate,
				LeagueId:    v.LeagueId,
				League:      v.LeagueNameAbbr,
				LeagueColor: v.LeagueBackColor,
				HomeId:      v.HomeTeamId,
				Home:        v.HomeTeam,
				AwayId:      v.AwayTeamId,
				Away:        v.AwayTeam,
				GoalLine:    goalLine,
				MatchStatus: matchStatus,
				HalfScore:   v.SectionsNo1,
				FullScore:   v.SectionsNo999,
			})
		}

		if len(insertData) > 0 {
			err = scoreModel.BatchInsertOnUpdate(insertData, []string{"match_num", "match_name", "match_date", "league_id", "league", "league_color", "home_id", "home", "away_id", "away", "goal_line", "match_status", "half_score", "full_score"})
			if err != nil {
				return err
			}
		}

		return nil
	},
}
