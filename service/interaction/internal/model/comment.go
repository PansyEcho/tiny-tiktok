package model

import (
	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	ID         int64  `gorm:"column:id; not null; type:int primary key auto_increment; comment:'ID'" json:"id"`
	UserID     int64  `gorm:"column:user_id; not null; type:int; index:idx_userid; comment:'评论人ID'" json:"user_id"`
	VideoID    int64  `gorm:"column:video_id; not null; type:int; index:idx_videoid; comment:'被评论视频ID'" json:"video_id"`
	Content    string `gorm:"column:content; type:varchar(255); comment:'评论内容'" json:"content"`
	Cancel     int    `gorm:"column:cancel; not null; type:tinyint default 0; comment:'是否删除评论'" json:"cancel"`
	CreateDate string `gorm:"column:create_date; not null; type:string; comment:'评论发布日期'" json:"create_date"`
	ExtraA     string `gorm:"column:extraA; type:varchar(255); comment:'额外字段A'" json:"extraA"`
	ExtraB     string `gorm:"column:extraB; type:varchar(255); comment:'额外字段B'" json:"extraB"`
}

func (table Comment) TableName() string {
	return "comment"
}
