package config

// ListFields 列表展示字段
var ListFields = []string{"baseTable.mid", "baseTable.mUrl", "baseTable.mName", "baseTable.mPoster",
	"baseTable.mTypeName", "baseTable.mDoubanScore", "baseTable.mDirector", "baseTable.mStarring",
	"baseTable.mCountryName", "baseTable.mLanguageName", "baseTable.mDate",
	"viewTable.mViews", "likeTable.mLikes", "collectTable.mCollects",
}

// LeaderboardFields 排行榜列表展示字段
var LeaderboardFields = []string{"baseTable.mid", "baseTable.mName", "totalTable.mTotal"}
