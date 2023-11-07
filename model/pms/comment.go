package pms

import (
	"github.com/AuroraOps04/mall-go/model/base"
)

type Comment struct {
	base.BaseModel
	ProductID        uint            `json:"product_id"`
	MemberNickName   string          `json:"member_nick_name" gorm:"size:255"`
	ProductName      string          `json:"product_name" gorm:"size:255"`
	Star             uint8           `json:"star" gorm:"size:3"`
	MemberIp         string          `json:"member_ip" gorm:"size:64"`
	ShowStatus       int8            `json:"show_status" gorm:"size:1"`
	ProductAttribute string          `json:"product_attribute" gorm:"size:255"`
	CollectCount     uint            `json:"collect_count"`
	ReadCount        uint            `json:"read_count"`
	Content          string          `json:"content" gorm:"type:text"`
	Pics             string          `json:"pics" gorm:"size:1000"`
	MemberIcon       string          `json:"member_icon" gorm:"size:255"`
	ReplayCount      uint            `json:"replay_count"`
	Replies          []*CommentReply `json:"replies" gorm:"foreignKey:CommentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (p Comment) TableName() string {
	return "pms_comment"
}

type CommentReply struct {
	base.BaseModel
	CommentID      uint     `json:"comment_id"`
	Comment        *Comment `json:"comment" gorm:"foreignKey:CommentID"`
	MemberNickName string   `json:"member_nick_name" gorm:"size:255"`
	MemberIcon     string   `json:"member_icon" gorm:"size:255"`
	Content        string   `json:"content" gorm:"size:1000"`
	Type           int8     `json:"type" gorm:"size:1"`
}

func (p CommentReply) TableName() string {
	return "pms_comment_reply"
}
