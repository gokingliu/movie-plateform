package logic

import (
	"MovieService/config"
	"MovieService/models"
	pb "git.woa.com/crotaliu/pb-hub"
	"gorm.io/gorm"
)

// AddInfoLogic 增加视频逻辑
func AddInfoLogic(db *gorm.DB, req *pb.AddInfoReq) error {
	var addInfoSplice = make([]map[string]interface{}, 0)
	for _, value := range req.List {
		addInfoSplice = append(addInfoSplice, map[string]interface{}{
			"MUrl":          value.MUrl,
			"MName":         value.MName,
			"MPoster":       value.MPoster,
			"MTypeID":       value.MTypeID,
			"MTypeName":     value.MTypeName,
			"MDouBanScore":  value.MDouBanScore,
			"MDirector":     value.MDirector,
			"MStarring":     value.MStarring,
			"MCountryID":    value.MCountryID,
			"MCountryName":  value.MCountryName,
			"MLanguageID":   value.MLanguageID,
			"MLanguageName": value.MLanguageName,
			"MDateYear":     value.MDateYear,
			"MDate":         value.MDate,
			"MDesc":         value.MDesc,
		})
	}
	dbResult := db.Debug().Model(&models.List{}).Create(&addInfoSplice)

	return dbResult.Error
}

// UpdateInfoLogic 修改视频逻辑
func UpdateInfoLogic(db *gorm.DB, req *pb.UpdateInfoReq) (error, int64) {
	dbResult := db.Debug().Model(&models.List{}).Select(config.UpdateInfoFields).
		Where("mid IN ?", req.Mids).Updates(map[string]interface{}{
		"mTypeID":       req.Info.MTypeID,
		"mTypeName":     req.Info.MTypeName,
		"mDouBanScore":  req.Info.MDouBanScore,
		"mCountryID":    req.Info.MCountryID,
		"mCountryName":  req.Info.MCountryName,
		"mLanguageID":   req.Info.MLanguageID,
		"mLanguageName": req.Info.MLanguageName,
		"mDateYear":     req.Info.MDateYear,
	})

	return dbResult.Error, dbResult.RowsAffected
}

// UpdateInfoStatusLogic 修改视频状态逻辑
func UpdateInfoStatusLogic(db *gorm.DB, req *pb.UpdateInfoStatusReq) (error, int64) {
	dbResult := db.Debug().Model(&models.List{}).Select("mStatus").
		Where("mid IN ?", req.Mids).Updates(map[string]interface{}{"mStatus": req.MStatus})

	return dbResult.Error, dbResult.RowsAffected
}

// DelInfoLogic 删除视频逻辑
func DelInfoLogic(db *gorm.DB, req *pb.DelInfoReq) (error, int64) {
	dbResult := db.Debug().Delete(&models.List{}, req.Mids)

	return dbResult.Error, dbResult.RowsAffected
}
