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
func GetListLogic(db *gorm.DB, data map[string]interface{}, role int64, pageNo, pageSize uint32) ([]models.ListInfo, uint32, error) {
	var ListInfo []models.ListInfo
	var count int64
	// 获取 WHERE 条件
	whereSQL := utils.SpliceWhereSql(data)
	// 不同角色 WHERE 条件不同
	if role != 3 {
		if whereSQL == "" {
			whereSQL += fmt.Sprintf(" mStatus = %d", 1)
		} else {
			whereSQL += fmt.Sprintf(" AND mStatus = %d", 1)
		}
	}
	// SELECT 字段
	selectSQL := config.ListFields
	// 分页列表
	dbResult := db.Debug().Raw(
		"SELECT "+strings.Join(selectSQL, ",")+" FROM list AS baseTable "+
			"LEFT JOIN (SELECT mid, COUNT(*) AS mViews FROM record WHERE type = 3 GROUP BY mid) AS viewTable "+
			"ON baseTable.mid = viewTable.mid "+
			"LEFT JOIN (SELECT mid, COUNT(*) AS mLikes FROM record WHERE type = 1 GROUP BY mid) AS likeTable "+
			"ON baseTable.mid = likeTable.mid "+
			"LEFT JOIN (SELECT mid, COUNT(*) AS mCollects FROM record WHERE type = 2 GROUP BY mid) AS collectTable "+
			"ON baseTable.mid = collectTable.mid "+
			"WHERE baseTable.mid >= (SELECT mid FROM list WHERE "+whereSQL+" LIMIT ?,1) LIMIT ?",
		(pageNo-1)*pageSize, pageSize,
	).Scan(&ListInfo)
	// 总量
	countResult := db.Debug().Model(&models.List{}).Where(whereSQL).Count(&count)
	// 任意 db 读取错误，则返回错误
	var dbError error
	if dbResult.Error != nil || countResult.Error != nil {
		dbError = config.New(config.InnerReadDbError)
	}

	return ListInfo, uint32(count), dbError
}

// GetLeaderboardLogic 获取视频排行榜
func GetLeaderboardLogic(db *gorm.DB, mType uint32) ([]*pb.GetLeaderboardRsp_List, error) {
	var list []*pb.GetLeaderboardRsp_List
	// SELECT 字段
	selectSQL := config.LeaderboardFields
	// 查询列表
	dbResult := db.Debug().Raw(
		"SELECT "+strings.Join(selectSQL, ",")+" FROM list AS baseTable "+
			"RIGHT JOIN "+
			"(SELECT mid, COUNT(*) AS mTotal FROM record WHERE type = ? GROUP BY mid ORDER BY mTotal DESC LIMIT 5) "+
			"AS totalTable ON baseTable.mid = totalTable.mid WHERE mStatus = 1", mType,
	).Scan(&list)

	return list, dbResult.Error
}
