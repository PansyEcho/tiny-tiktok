package model

import "gorm.io/gorm"

type Favorite struct {
	gorm.Model
	ID      int64  `gorm:"column:id; not null; type:int primary key auto_increment; comment:'ID'" json:"id"`
	UserID  int64  `gorm:"column:user_id; not null; type:int; index:idx_userid; comment:'点赞人ID'" json:"user_id"`
	VideoID int64  `gorm:"column:video_id; not null; type:int; index:idx_videoid; comment:'被点赞视频ID'" json:"video_id"`
	Cancel  int    `gorm:"column:cancel; not null; type:tinyint default 0; comment:'是否取消点赞'" json:"cancel"`
	ExtraA  string `gorm:"column:extraA; type:varchar(255); comment:'额外字段A'" json:"extraA"`
	ExtraB  string `gorm:"column:extraB; type:varchar(255); comment:'额外字段B'" json:"extraB"`
}

func (table Favorite) TableName() string {
	return "favorite"
}
