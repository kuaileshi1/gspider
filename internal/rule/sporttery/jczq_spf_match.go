// Package sporttery
// @Description: 竞彩网开奖结果抓取
// @Auth shigx 2023-04-24 16:58:25
package sporttery

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
	"gspider/internal/pkg/spider"
)

var ruleSpfMatch = &spider.TaskRule{
	Name:        "竞彩网足球胜平负比赛",
	Description: "竞彩足球胜平负比赛信息抓取",
	Rule: &spider.Rule{
		EnterFun: func(c *colly.Collector) error {
			return c.Visit("https://webapi.sporttery.cn/gateway/jc/football/getMatchCalculatorV1.qry?poolCode=hhad,had&channel=c")
		},
		Nodes: map[int]*spider.Node{
			0: stepSpf1,
			//1: stepSpf2,
		},
	},
}

// 比赛信息
type matchResult struct {
	ErrorCode    string `json:"errorCode"`
	ErrorMessage string `json:"errorMessage"`
	Success      bool   `json:"success"`
	Value        struct {
		TotalCount    int `json:"totalCount"`
		MatchInfoList []struct {
			BusinessDate string `json:"businessDate"`
			SubMatchList []struct {
				AwayRank        string `json:"awayRank"`
				AwayTeamAbbName string `json:"awayTeamAbbName"`
				AwayTeamId      int    `json:"awayTeamId"`
				BackColor       string `json:"backColor"`
				BusinessDate    string `json:"businessDate"`
				Had             struct {
					A        string `json:"a"`
					Af       string `json:"af"`
					D        string `json:"d"`
					Df       string `json:"df"`
					GoalLine string `json:"goalLine"`
					H        string `json:"h"`
					Hf       string `json:"hf"`
				} `json:"had"`
				Hhad struct {
					A        string `json:"a"`
					Af       string `json:"af"`
					D        string `json:"d"`
					Df       string `json:"df"`
					GoalLine string `json:"goalLine"`
					H        string `json:"h"`
					Hf       string `json:"hf"`
				} `json:"hhad"`
				HomeRank        string `json:"homeRank"`
				HomeTeamAbbName string `json:"homeTeamAbbName"`
				HomeTeamId      int    `json:"homeTeamId"`
				LeagueAbbName   string `json:"leagueAbbName"`
				LeagueId        int    `json:"leagueId"`
				MatchDate       string `json:"matchDate"`
				MatchId         int    `json:"matchId"`
				MatchNum        int    `json:"matchNum"`
				MatchNumStr     string `json:"matchNumStr"`
				MatchStatus     string `json:"matchStatus"`
				MatchTime       string `json:"matchTime"`
				SellStatus      int    `json:"sellStatus"`
			} `json:"subMatchList"`
		} `json:"matchInfoList"`
	} `json:"value"`
}

var stepSpf1 = &spider.Node{
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

		var response matchResult
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

		return nil
	},
}

var stepSpf2 = &spider.Node{
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

		return nil
	},
}
