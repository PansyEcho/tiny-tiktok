package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID             int64  `gorm:"column:id; not null; type:int primary key auto_increment; comment:'ID';" json:"id"`
	Username       string `gorm:"column:username; not null; type:varchar(255);index:idx_username; comment:'用户名'" json:"username"`
	Password       string `gorm:"column:password; not null; type:varchar(255); comment:'密码'" json:"password"`
	Follow_count   int64  `gorm:"column:follow_count; type:int default 0; comment:'关注数'" json:"follow_count"`
	Follower_count int64  `gorm:"column:follower_count; type:int default 0; comment:'粉丝数'" json:"follower_count"`
	Avatar         string `gorm:"column:avatar; type:varchar(255); comment:'头像地址'" json:"avatar"`
	ExtraA         string `gorm:"column:extraA; type:varchar(255); comment:'额外字段A'" json:"extraA"`
	ExtraB         string `gorm:"column:extraB; type:varchar(255); comment:'额外字段B'" json:"extraB"`
}

func (table User) TableName() string {
	return "user"
}
