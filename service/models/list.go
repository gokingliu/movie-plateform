package models

type List struct {
	Mid           uint32  `gorm:"primaryKey;column:mid;type:int(10) unsigned;not null"`                           // 电影ID，自增ID
	MUrl          string  `gorm:"column:mUrl;type:varchar(255);not null"`                                         // 电影URL
	MName         string  `gorm:"column:mName;type:varchar(255);not null"`                                        // 电影名
	MPoster       string  `gorm:"column:mPoster;type:varchar(255);not null"`                                      // 电影海报
	MTypeID       uint32  `gorm:"column:mTypeID;type:tinyint(4) unsigned;not null"`                               // 电影类型ID
	MTypeName     string  `gorm:"column:mTypeName;type:varchar(10);not null"`                                     // 电影类型
	MDouBanScore  float32 `gorm:"column:mDouBanScore;type:float(2,1);not null"`                                   // 豆瓣评分
	MDirector     string  `gorm:"column:mDirector;type:varchar(50);not null"`                                     // 电影导演
	MStarring     string  `gorm:"column:mStarring;type:varchar(255);not null"`                                    // 电影主演
	MCountryID    uint32  `gorm:"column:mCountryID;type:tinyint(4) unsigned;not null"`                            // 电影制片国家/地区ID
	MCountryName  string  `gorm:"column:mCountryName;type:varchar(10);not null"`                                  // 电影制片国家/地区
	MLanguageID   uint8   `gorm:"column:mLanguageID;type:tinyint(4) unsigned;not null"`                           // 电影语言ID
	MLanguageName string  `gorm:"column:mLanguageName;type:varchar(10);not null"`                                 // 电影语言
	MDateYear     uint32  `gorm:"column:mDateYear;type:smallint(6) unsigned;not null"`                            // 电影上映年份
	MDate         string  `gorm:"column:mDate;type:varchar(10);not null"`                                         // 电影上映日期
	MDesc         string  `gorm:"column:mDesc;type:varchar(1024);not null"`                                       // 电影简介
	MStatus       uint32  `gorm:"column:mStatus;type:tinyint(4) unsigned;not null"`                               // 电影状态
	CreateTime    uint64  `gorm:"column:createTime;type:int(10) unsigned;not null;default:unix_timestamp(now())"` // 创建时间
	UpdateTime    uint64  `gorm:"column:updateTime;type:int(10) unsigned;not null;default:unix_timestamp(now())"` // 更新时间
}

type ListResult struct {
	Mid           uint32  `gorm:"primaryKey;column:mid;type:int(10) unsigned;not null"`                           // 电影ID，自增ID
	MName         string  `gorm:"column:mName;type:varchar(255);not null"`                                        // 电影名
	MPoster       string  `gorm:"column:mPoster;type:varchar(255);not null"`                                      // 电影海报
	MTypeName     string  `gorm:"column:mTypeName;type:varchar(10);not null"`                                     // 电影类型
	MDouBanScore  float32 `gorm:"column:mDouBanScore;type:float(2,1);not null"`                                   // 豆瓣评分
	MDirector     string  `gorm:"column:mDirector;type:varchar(50);not null"`                                     // 电影导演
	MStarring     string  `gorm:"column:mStarring;type:varchar(255);not null"`                                    // 电影主演
	MCountryName  string  `gorm:"column:mCountryName;type:varchar(10);not null"`                                  // 电影制片国家/地区
	MLanguageName string  `gorm:"column:mLanguageName;type:varchar(10);not null"`                                 // 电影语言
	MDateYear     uint32  `gorm:"column:mDateYear;type:smallint(6) unsigned;not null"`                            // 电影上映年份
	MDate         string  `gorm:"column:mDate;type:varchar(10);not null"`                                         // 电影上映日期
	CreateTime    uint64  `gorm:"column:createTime;type:int(10) unsigned;not null;default:unix_timestamp(now())"` // 创建时间
	UpdateTime    uint64  `gorm:"column:updateTime;type:int(10) unsigned;not null;default:unix_timestamp(now())"` // 更新时间
	MViews        uint32  `gorm:"column:mViews;type:int(10) unsigned;not null"`                                   // 播放量
	MLikes        uint32  `gorm:"column:mLikes;type:int(10) unsigned;not null"`                                   // 点赞量
	MCollects     uint32  `gorm:"column:mCollects;type:int(10) unsigned;not null"`                                // 收藏量
}
