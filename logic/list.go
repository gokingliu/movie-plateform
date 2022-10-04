package logic

import (
	"MovieService/config"
	"MovieService/models"
	"MovieService/utils"
	"fmt"
	"gorm.io/gorm"
	"strings"
)

// GetListLogic 获取电影列表
func GetListLogic(db *gorm.DB, data map[string]interface{}, role int64, pageNo, pageSize uint32) ([]models.List, uint32, error) {
	var ListInfo []models.List
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
	// 分页列表
	dbResult := db.Debug().Raw(
		"SELECT "+strings.Join(config.ListFields, ",")+" FROM list AS baseTable "+
			"LEFT JOIN (SELECT mid, COUNT(*) AS mViews FROM record WHERE mType = 3 GROUP BY mid) AS viewTable "+
			"ON baseTable.mid = viewTable.mid "+
			"LEFT JOIN (SELECT mid, COUNT(*) AS mLikes FROM record WHERE mType = 1 GROUP BY mid) AS likeTable "+
			"ON baseTable.mid = likeTable.mid "+
			"LEFT JOIN (SELECT mid, COUNT(*) AS mCollects FROM record WHERE mType = 2 GROUP BY mid) AS collectTable "+
			"ON baseTable.mid = collectTable.mid "+
			"WHERE "+whereSQL+" LIMIT ?,?",
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
func GetLeaderboardLogic(db *gorm.DB, mType uint32) ([]models.Record, error) {
	var list []models.Record
	// 查询列表
	dbResult := db.Debug().Raw(
		"SELECT "+strings.Join(config.LeaderboardFields, ",")+" FROM list AS baseTable "+
			"RIGHT JOIN "+
			"(SELECT mid, COUNT(*) AS mTotal FROM record WHERE mType = ? GROUP BY mid ORDER BY mTotal DESC LIMIT 5) "+
			"AS totalTable ON baseTable.mid = totalTable.mid WHERE mStatus = 1", mType,
	).Scan(&list)

	return list, dbResult.Error
}
