package logic

import (
	"MovieService/models"
	"git.code.oa.com/trpc-go/trpc-go/log"
	pb "git.woa.com/crotaliu/pb-hub"
	"gorm.io/gorm"
)

func AddInfoLogic(db *gorm.DB, req *pb.ManageInfoReq) error {
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
	log.Error(addInfoSplice)
	dbResult := db.Debug().Model(&models.List{}).Create(&addInfoSplice)

	return dbResult.Error
}
