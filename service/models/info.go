package models

type Info struct {
	Mid           uint32  `gorm:"primaryKey;column:mid;type:int(10) unsigned;not null"` // 电影ID，自增ID
	MUrl          string  `gorm:"column:mUrl;type:varchar(255);not null"`               // 电影URL
	MName         string  `gorm:"column:mName;type:varchar(255);not null"`              // 电影名
	MPoster       string  `gorm:"column:mPoster;type:varchar(255);not null"`            // 电影海报
	MTypeName     string  `gorm:"column:mTypeName;type:varchar(10);not null"`           // 电影类型
	MDouBanScore  float32 `gorm:"column:mDouBanScore;type:float(2,1);not null"`         // 豆瓣评分
	MDirector     string  `gorm:"column:mDirector;type:varchar(50);not null"`           // 电影导演
	MStarring     string  `gorm:"column:mStarring;type:varchar(255);not null"`          // 电影主演
	MCountryName  string  `gorm:"column:mCountryName;type:varchar(10);not null"`        // 电影制片国家/地区
	MLanguageName string  `gorm:"column:mLanguageName;type:varchar(10);not null"`       // 电影语言
	MDate         string  `gorm:"column:mDate;type:varchar(10);not null"`               // 电影上映日期
	MDesc         string  `gorm:"column:mDesc;type:varchar(1024);not null"`             // 电影简介
}

type Record struct {
	UserName   string `gorm:"column:userName;type:varchar(32);not null"`                           // 用户名
	Mid        uint32 `gorm:"column:mid;type:tinyint(4) unsigned;not null"`                        // 电影ID
	MType      uint32 `gorm:"column:mType;type:tinyint(4) unsigned;not null"`                      // 类型 1-点赞 2-收藏 3-播放
	CreateTime string `gorm:"column:createTime;type:timestamp;not null;default:CURRENT_TIMESTAMP"` // 创建时间
}
