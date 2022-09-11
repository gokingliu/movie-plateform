package logic

import (
	"MovieService/config"
	"MovieService/models"
	"fmt"
	pb "git.woa.com/crotaliu/pb-hub"
	"gorm.io/gorm"
	"strings"
)

// AddInfoLogic 增加视频逻辑
func AddInfoLogic(db *gorm.DB, req *pb.AddInfoReq) error {
	var addInfoSplice = make([]map[string]interface{}, 0)
	// TODO 不优雅，后续要优化
	for _, value := range req.List {
		addInfoSplice = append(addInfoSplice, map[string]interface{}{
			"mUrl":          value.MUrl,
			"mName":         value.MName,
			"mPoster":       value.MPoster,
			"mTypeID":       value.MTypeID,
			"mTypeName":     value.MTypeName,
			"mDouBanScore":  value.MDouBanScore,
			"mDirector":     value.MDirector,
			"mStarring":     value.MStarring,
			"mCountryID":    value.MCountryID,
			"mCountryName":  value.MCountryName,
			"mLanguageID":   value.MLanguageID,
			"mLanguageName": value.MLanguageName,
			"mDateYear":     value.MDateYear,
			"mDate":         value.MDate,
			"mDesc":         value.MDesc,
		})
	}
	dbResult := db.Debug().Model(&models.List{}).Create(&addInfoSplice)

	return dbResult.Error
}

// UpdateInfoLogic 修改视频逻辑
func UpdateInfoLogic(db *gorm.DB, req *pb.UpdateInfoReq) (error, int64) {
	// TODO 不优雅，后续要优化
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

// AddPropLogic 添加属性逻辑
func AddPropLogic(db *gorm.DB, req *pb.AddPropReq) error {
	var addPropSplice = make([]map[string]interface{}, 0)
	// TODO 不优雅，后续要优化
	for _, value := range req.Props {
		addPropSplice = append(addPropSplice, map[string]interface{}{
			"label": value.Label,
			"value": value.Value,
			"mType": value.MType,
		})
	}
	dbResult := db.Debug().Model(&models.Prop{}).Create(&addPropSplice)

	return dbResult.Error
}

// GetPropLogic 获取属性逻辑
func GetPropLogic(db *gorm.DB, req *pb.GetPropReq) (error, []*pb.GetPropRsp_Result) {
	var props []*pb.GetPropRsp_Result
	dbResult := db.Debug().Model(&models.Prop{}).Select(config.PropFields).
		Where("mType = ?", req.MType).Scan(&props)

	return dbResult.Error, props
}

// UpdatePropLogic 更新属性逻辑
func UpdatePropLogic(db *gorm.DB, req *pb.UpdatePropReq) (error, int64) {
	var updateValueSplice []string
	for _, value := range req.Props {
		updateValueSplice = append(
			updateValueSplice,
			"("+fmt.Sprintf("%d, '%s', %d, 0", value.Id, value.Label, value.Value)+")",
		)
	}
	dbResult := db.Debug().Exec(
		"INSERT INTO prop " +
			"(id, label, value, mType) " +
			"VALUES " + strings.Join(updateValueSplice, ", ") + " " +
			"ON DUPLICATE KEY UPDATE " +
			"label = IF (VALUES(label) > '', VALUES(label), label), " +
			"value = IF (VALUES(value) > 0, VALUES(value), value), " +
			"mType = IF (VALUES(mType) > 0, VALUES(mType), mType)",
	)

	return dbResult.Error, dbResult.RowsAffected
}

// DelPropLogic 删除属性逻辑
func DelPropLogic(db *gorm.DB, req *pb.DelPropReq) (error, int64) {
	dbResult := db.Debug().Delete(&models.Prop{}, req.Ids)

	return dbResult.Error, dbResult.RowsAffected
}
