package entity

import "time"

type SportteryJczqScore struct {
	Id          int64     `gorm:"column:id;primary_key;AUTO_INCREMENT;NOT NULL"`
	MatchId     int       `gorm:"column:match_id;default:0;NOT NULL;comment:'比赛编号'"`
	MatchNum    int       `gorm:"column:match_num;default:0;NOT NULL;comment:'场次编号'"`
	MatchName   string    `gorm:"column:match_name;default:;NOT NULL;comment:'场次名称'"`
	MatchDate   string    `gorm:"column:match_date;default:1000-01-01;NOT NULL;comment:'比赛日期'"`
	LeagueId    int       `gorm:"column:league_id;default:0;NOT NULL;comment:'联赛编号'"`
	League      string    `gorm:"column:league;default:;NOT NULL;comment:'联赛名称'"`
	LeagueColor string    `gorm:"column:league_color;default:;NOT NULL;comment:'联赛背景颜色'"`
	HomeId      int       `gorm:"column:home_id;default:0;NOT NULL;comment:'主队编号'"`
	Home        string    `gorm:"column:home;default:;NOT NULL;comment:'主队'"`
	AwayId      int       `gorm:"column:away_id;default:0;NOT NULL;comment:'客队编号'"`
	Away        string    `gorm:"column:away;default:;NOT NULL;comment:'客队'"`
	GoalLine    int       `gorm:"column:goal_line;default:0;NOT NULL;comment:'让球数'"`
	HalfScore   string    `gorm:"column:half_score;default:;NOT NULL;comment:'半场比分'"`
	FullScore   string    `gorm:"column:full_score;default:;NOT NULL;comment:'全场比分'"`
	MatchStatus int       `gorm:"column:match_status;default:0;NOT NULL;comment:'比赛结果'"`
	CreatedAt   time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP;NOT NULL"`
	UpdatedAt   time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP;NOT NULL"`
}

func (s *SportteryJczqScore) TableName() string {
	return "sporttery_jczq_score"
}
