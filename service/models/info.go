package models

type Record struct {
	// 用户名
	UserName string
	// 电影ID
	Mid uint32
	// 类型 1-点赞 2-收藏 3-播放
	Type uint32
	// 创建时间
	CreateTime uint64
}
