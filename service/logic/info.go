package logic

import (
	"MovieService/config"
	"MovieService/models"
	pb "git.woa.com/crotaliu/pb-hub"
	"gorm.io/gorm"
)

// GetInfoLogic 获取视频详情
func GetInfoLogic(db *gorm.DB, mid uint32) (*pb.GetInfoRsp_Result, error) {
	var info *pb.GetInfoRsp_Result
	// SELECT 字段
	selectSQL := config.InfoFields
	// 查询列表
	infoResult := db.Debug().Model(&models.List{}).Select(selectSQL).Where("mid = ?", mid).First(&info)

	return info, infoResult.Error
}

// GetRecordLogic 获取记录
func GetRecordLogic(db *gorm.DB, userName string, mid uint32, mType uint32) (bool, error) {
	var selectMid int32
	// 查询列表
	midResult := db.Debug().Model(&models.Record{}).Select("mid").
		Where("userName = ? AND mid = ? AND type = ? ", userName, mid, mType).First(&selectMid)

	return selectMid >= 0, midResult.Error
}
