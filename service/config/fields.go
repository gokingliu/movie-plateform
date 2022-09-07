package config

// ListFields 列表展示字段
var ListFields = []string{
	"baseTable.mid", "baseTable.mName", "baseTable.mPoster", "baseTable.mTypeName", "baseTable.mDouBanScore",
	"baseTable.mDirector", "baseTable.mStarring", "baseTable.mCountryName", "baseTable.mLanguageName",
	"baseTable.mDateYear", "baseTable.mDate", "baseTable.createTime", "baseTable.updateTime",
	"viewTable.mViews", "likeTable.mLikes", "collectTable.mCollects",
}

// LeaderboardFields 排行榜列表展示字段
var LeaderboardFields = []string{"baseTable.mid", "baseTable.mName", "totalTable.mTotal"}

// InfoFields 视频详情展示字段
var InfoFields = []string{
	"mid", "mUrl", "mName", "mPoster", "mTypeName", "mDouBanScore", "mDirector", "mStarring", "mCountryName",
	"mLanguageName", "mDate", "mDesc",
}

// UpdateInfoFields 修改视频详情字段
var UpdateInfoFields = []string{
	"mTypeID", "mTypeName", "mDouBanScore", "mCountryID", "mCountryName", "mLanguageID", "mLanguageName", "mDateYear",
}
