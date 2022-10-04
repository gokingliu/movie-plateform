package logic

import (
	"MovieService/config"
	"MovieService/models"
	"gorm.io/gorm"
)

// GetInfoLogic 获取视频详情
func GetInfoLogic(db *gorm.DB, mid uint32) (models.Info, error) {
	var info models.Info
	// SELECT 字段
	selectSQL := config.InfoFields
	// 查询列表
	dbResult := db.Debug().Model(&models.List{}).Select(selectSQL).Where("mid = ?", mid).Take(&info)

	return info, dbResult.Error
}

// GetRecordLogic 获取记录
func GetRecordLogic(db *gorm.DB, userName string, mid uint32, mType uint32) (bool, error) {
	var count int64
	// 查询列表
	midResult := db.Debug().Model(&models.Record{}).Select("mid").
		Where("userName = ? AND mid = ? AND mType = ? ", userName, mid, mType).Count(&count)

	return count > 0, midResult.Error
}

// AddRecordLogic 添加记录
func AddRecordLogic(db *gorm.DB, userName string, mid uint32, mType uint32) error {
	dbResult := db.Debug().Model(&models.Record{}).Create(map[string]interface{}{
		"userName": userName,
		"mid":      mid,
		"mType":    mType,
	})

	return dbResult.Error
}

// DelRecordLogic 删除记录
func DelRecordLogic(db *gorm.DB, userName string, mid uint32, mType uint32) error {
	dbResult := db.Debug().
		Where("userName = ? AND mid = ? AND mType = ?", userName, mid, mType).Delete(&models.Record{})

	return dbResult.Error
}
