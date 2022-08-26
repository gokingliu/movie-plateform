package logic

import (
	"MovieService/config"
	"MovieService/models"
	"MovieService/utils"
	"fmt"
	pb "git.woa.com/crotaliu/pb-hub"
	"gorm.io/gorm"
	"strings"
)

// GetListLogic 获取电影列表
func GetListLogic(db *gorm.DB, data map[string]interface{}, role uint32, pageNo, pageSize uint32) ([]*pb.GetListRsp_List, int64, error) {
	var list []*pb.GetListRsp_List
	var count int64
	// 获取 WHERE 条件
	whereSQL := utils.SpliceWhereSql(data, role)
	// 不同角色 WHERE 条件不同
	if role != uint32(3) {
		whereSQL += fmt.Sprintf(" AND mStatus = %d", 1)
	}
	// SELECT 字段
	selectSQL := config.ListFields
	// 分页列表
	listResult := db.Debug().Raw(
		"SELECT "+strings.Join(selectSQL, ",")+" FROM list AS baseTable "+
			"LEFT JOIN (SELECT mid, COUNT(*) AS mViews FROM record WHERE type = 3 GROUP BY mid) AS viewTable "+
			"ON baseTable.mid = viewTable.mid "+
			"LEFT JOIN (SELECT mid, COUNT(*) AS mLikes FROM record WHERE type = 1 GROUP BY mid) AS likeTable "+
			"ON baseTable.mid = likeTable.mid "+
			"LEFT JOIN (SELECT mid, COUNT(*) AS mCollects FROM record WHERE type = 2 GROUP BY mid) AS collectTable "+
			"ON baseTable.mid = collectTable.mid "+
			"WHERE baseTable.mid >= (SELECT mid FROM list WHERE "+whereSQL+" LIMIT ?,1) LIMIT ?",
		(pageNo-1)*pageSize, pageSize,
	).Scan(&list)
	// 总量
	countResult := db.Debug().Model(&models.List{}).Where(whereSQL).Count(&count)
	// 任意 db 读取错误，则返回错误
	var dbError error
	if listResult.Error != nil || countResult.Error != nil {
		dbError = config.New(config.InnerReadDbError)
	}

	return list, count, dbError
}

// GetLeaderboardLogic 获取视频排行榜
func GetLeaderboardLogic(db *gorm.DB, mType uint32) ([]*pb.GetLeaderboardRsp_List, error) {
	var list []*pb.GetLeaderboardRsp_List
	// SELECT 字段
	selectSQL := config.LeaderboardFields
	// 查询列表
	listResult := db.Debug().Raw(
		"SELECT "+strings.Join(selectSQL, ",")+" FROM list AS baseTable "+
			"RIGHT JOIN "+
			"(SELECT mid, COUNT(*) AS mTotal FROM record WHERE type = ? GROUP BY mid ORDER BY mTotal DESC LIMIT 5) "+
			"AS totalTable ON baseTable.mid = totalTable.mid WHERE mStatus = 1", mType,
	).Scan(&list)

	return list, listResult.Error
}
