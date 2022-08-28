package services

import (
	"MovieService/config"
	"MovieService/logic"
	"MovieService/utils"
	"context"
	pb "git.woa.com/crotaliu/pb-hub"
)

type InfoImpl struct{}

// GetInfo 获取视频详情
func (s *InfoImpl) GetInfo(ctx context.Context, req *pb.GetInfoReq, rsp *pb.GetInfoRsp) error {
	// ConnDB 实例
	db := utils.ConnDB()
	// 判断 token，并获取用户名、用户角色
	tokenBol, _, _, tokenErr := logic.PreHandleTokenLogic(db, ctx)
	// 查询视频详情逻辑
	info, dbErr := logic.GetInfoLogic(db, req.Mid)
	if dbErr != nil {
		rsp.Code, rsp.Msg = config.InnerReadDbError.Code, config.InnerReadDbError.Msg
		return nil
	}
	// 客户端接口，判断 token 解析是否正常
	// 不正常时返回相应的错误码，同时返回视频详情，前端正常展示，但清空用户信息
	if !tokenBol {
		rsp.Code, rsp.Msg = config.ClientUserInfoError.Code, utils.GetErrorMap(tokenErr.Error()).Msg
	} else {
		rsp.Code, rsp.Msg = config.ResOk.Code, config.ResOk.Msg
	}
	// 解构类型重新赋值
	// TODO 不优雅，后续要优化
	rsp.Result = &pb.GetInfoRsp_Result{
		Mid:           info.Mid,
		MUrl:          info.MUrl,
		MName:         info.MName,
		MPoster:       info.MPoster,
		MTypeName:     info.MTypeName,
		MDouBanScore:  info.MDouBanScore,
		MDirector:     info.MDirector,
		MStarring:     info.MStarring,
		MCountryName:  info.MCountryName,
		MLanguageName: info.MLanguageName,
		MDate:         info.MDate,
		MDesc:         info.MDesc,
	}

	return nil
}

// GetRecord 获取记录
func (s *InfoImpl) GetRecord(ctx context.Context, req *pb.RecordReq, rsp *pb.RecordRsp) error {
	// ConnDB 实例
	db := utils.ConnDB()
	// 判断 token，并获取用户名、用户角色
	tokenBol, userName, _, tokenErr := logic.PreHandleTokenLogic(db, ctx)
	// 客户端接口，判断 token 解析是否正常
	// 不正常时不允许请求 DB，并返回相应的错误码，清空用户信息
	if !tokenBol || tokenErr != nil {
		rsp.Code, rsp.Msg = config.ClientUserInfoError.Code, utils.GetErrorMap(tokenErr.Error()).Msg
		return nil
	}
	// 查询记录逻辑
	result, dbErr := logic.GetRecordLogic(db, userName, req.Mid, req.MType)
	// 查询 DB 错误
	if dbErr != nil {
		rsp.Code, rsp.Msg = config.InnerReadDbError.Code, config.InnerReadDbError.Msg
		rsp.Result = false
		return nil
	}
	// 正常返回
	rsp.Code, rsp.Msg = config.ResOk.Code, config.ResOk.Msg
	rsp.Result = result

	return nil
}

// PostRecord 添加记录
func (s *InfoImpl) PostRecord(ctx context.Context, req *pb.RecordReq, rsp *pb.RecordRsp) error {
	// ConnDB 实例
	db := utils.ConnDB()
	// 判断 token，并获取用户名、用户角色
	tokenBol, userName, _, tokenErr := logic.PreHandleTokenLogic(db, ctx)
	// 客户端接口，判断 token 解析是否正常
	// 不正常时不允许请求 DB，并返回相应的错误码，清空用户信息
	if !tokenBol || tokenErr != nil {
		rsp.Code, rsp.Msg = config.ClientUserInfoError.Code, utils.GetErrorMap(tokenErr.Error()).Msg
		return nil
	}
	// 添加记录逻辑
	dbErr := logic.PostRecordLogic(db, userName, req.Mid, req.MType)
	// 添加 DB 错误
	if dbErr != nil {
		rsp.Code, rsp.Msg = config.InnerReadDbError.Code, config.InnerReadDbError.Msg
		return nil
	}
	// 正常返回
	rsp.Code, rsp.Msg = config.ResOk.Code, config.ResOk.Msg

	return nil
}

func (s *InfoImpl) DelRecord(ctx context.Context, req *pb.RecordReq, rsp *pb.RecordRsp) error {
	// ConnDB 实例
	db := utils.ConnDB()
	// 判断 token，并获取用户名、用户角色
	tokenBol, userName, _, tokenErr := logic.PreHandleTokenLogic(db, ctx)
	// 客户端接口，判断 token 解析是否正常
	// 不正常时不允许请求 DB，并返回相应的错误码，清空用户信息
	if !tokenBol || tokenErr != nil {
		rsp.Code, rsp.Msg = config.ClientUserInfoError.Code, utils.GetErrorMap(tokenErr.Error()).Msg
		return nil
	}
	// 删除记录逻辑
	dbErr := logic.DelRecordLogic(db, userName, req.Mid, req.MType)
	// 添加 DB 错误
	if dbErr != nil {
		rsp.Code, rsp.Msg = config.InnerReadDbError.Code, config.InnerReadDbError.Msg
		return nil
	}
	// 正常返回
	rsp.Code, rsp.Msg = config.ResOk.Code, config.ResOk.Msg

	return nil
}
