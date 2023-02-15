package model

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	ID       int64  `gorm:"column:id; not null; type:int primary key auto_increment; comment:'ID'" json:"id"`
	UserID   int64  `gorm:"column:user_id; not null; type:int; index:idx_userid; comment:'关注人ID'" json:"user_id"`
	FollowID int64  `gorm:"column:follow_id; not null; type:int; index:idx_followid; comment:'被关注人ID'" json:"follow_id"`
	Cancel   int    `gorm:"column:cancel; not null; type:tinyint default 0; comment:'是否取消关注'" json:"cancel"`
	ExtraA   string `gorm:"column:extraA; type:varchar(255); comment:'额外字段A'" json:"extraA"`
	ExtraB   string `gorm:"column:extraB; type:varchar(255); comment:'额外字段B'" json:"extraB"`
}

func (table Follow) TableName() string {
	return "follow"
}
