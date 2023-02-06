package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	ID            int64  `gorm:"column:id; not null; type:int primary key auto_increment; comment:'ID';" json:"id"`
	AuthorID      int64  `gorm:"column:author_id; not null; type:int;index:idx_userid comment:'用户ID';" json:"author_id"`
	PlayURL       string `gorm:"column:play_url; not null; type:varchar(1024); comment:'视频播放地址'" json:"play_url"`
	CoverURL      string `gorm:"column:cover_url; not null; type:varchar(1024); comment:'视频封面地址'" json:"cover_url"`
	PublishTime   string `gorm:"column:publish_time; not null; type:datetime; comment:'发布时间'" json:"publish_time"`
	Title         string `gorm:"column:title; type:varchar(255); comment:'标题'" json:"title"`
	FavoriteCount int64  `gorm:"column:favorite_count; not null; type:int default 0; comment:'点赞数';" json:"favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count; not null; type:int default 0; comment:'评论数';" json:"comment_count"`
	ExtraA        string `gorm:"column:extraA; type:varchar(255); comment:'额外字段A'" json:"extraA"`
	ExtraB        string `gorm:"column:extraB; type:varchar(255); comment:'额外字段B'" json:"extraB"`
}

func (table Video) TableName() string {
	return "video"
}
