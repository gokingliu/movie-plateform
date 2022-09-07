package services

import (
	"MovieService/config"
	"MovieService/logic"
	"MovieService/utils"
	"context"
	pb "git.woa.com/crotaliu/pb-hub"
)

type ManageImpl struct{}

// AddInfo 增加视频
func (s *ManageImpl) AddInfo(ctx context.Context, req *pb.AddInfoReq, rsp *pb.ManageInfoRsp) error {
	// ConnDB 实例
	db := utils.ConnDB()
	// 判断 token，并获取用户名、用户角色
	tokenBol, _, role, tokenErr := logic.PreHandleTokenLogic(db, ctx)
	// 不正常时不允许请求 DB，并返回相应的错误码，清空用户信息
	if !tokenBol || tokenErr != nil {
		rsp.Code, rsp.Msg = config.ClientInvalidTokenError.Code, utils.GetErrorMap(tokenErr.Error()).Msg
		return nil
	}
	if role != 3 {
		rsp.Code, rsp.Msg = config.ClientPermissionError.Code, config.ClientPermissionError.Msg
		return nil
	}
	// 增加视频逻辑
	dbErr := logic.AddInfoLogic(db, req)
	// 添加 DB 错误
	if dbErr != nil {
		rsp.Code, rsp.Msg = config.InnerWriteDbError.Code, config.InnerWriteDbError.Msg
		return nil
	}
	// 正常返回
	rsp.Code, rsp.Msg = config.ResOk.Code, config.ResOk.Msg

	return nil
}

// UpdateInfo 修改视频
func (s *ManageImpl) UpdateInfo(ctx context.Context, req *pb.UpdateInfoReq, rsp *pb.ManageInfoRsp) error {
	// ConnDB 实例
	db := utils.ConnDB()
	// 判断 token，并获取用户名、用户角色
	tokenBol, _, role, tokenErr := logic.PreHandleTokenLogic(db, ctx)
	// 不正常时不允许请求 DB，并返回相应的错误码，清空用户信息
	if !tokenBol || tokenErr != nil {
		rsp.Code, rsp.Msg = config.ClientInvalidTokenError.Code, utils.GetErrorMap(tokenErr.Error()).Msg
		return nil
	}
	if role != 3 {
		rsp.Code, rsp.Msg = config.ClientPermissionError.Code, config.ClientPermissionError.Msg
		return nil
	}
	// 修改视频逻辑
	dbErr, _ := logic.UpdateInfoLogic(db, req)
	// 修改 DB 错误
	if dbErr != nil {
		rsp.Code, rsp.Msg = config.InnerWriteDbError.Code, config.InnerWriteDbError.Msg
		return nil
	}
	// 正常返回
	rsp.Code, rsp.Msg = config.ResOk.Code, config.ResOk.Msg

	return nil
}

// UpdateInfoStatus 修改视频状态
func (s *ManageImpl) UpdateInfoStatus(ctx context.Context, req *pb.UpdateInfoStatusReq, rsp *pb.ManageInfoRsp) error {
	// ConnDB 实例
	db := utils.ConnDB()
	// 判断 token，并获取用户名、用户角色
	tokenBol, _, role, tokenErr := logic.PreHandleTokenLogic(db, ctx)
	// 不正常时不允许请求 DB，并返回相应的错误码，清空用户信息
	if !tokenBol || tokenErr != nil {
		rsp.Code, rsp.Msg = config.ClientInvalidTokenError.Code, utils.GetErrorMap(tokenErr.Error()).Msg
		return nil
	}
	if role != 3 {
		rsp.Code, rsp.Msg = config.ClientPermissionError.Code, config.ClientPermissionError.Msg
		return nil
	}
	// 修改视频状态逻辑
	dbErr, _ := logic.UpdateInfoStatusLogic(db, req)
	// 修改 DB 错误
	if dbErr != nil {
		rsp.Code, rsp.Msg = config.InnerWriteDbError.Code, config.InnerWriteDbError.Msg
		return nil
	}
	// 正常返回
	rsp.Code, rsp.Msg = config.ResOk.Code, config.ResOk.Msg

	return nil
}

// DelInfo 删除视频
func (s *ManageImpl) DelInfo(ctx context.Context, req *pb.DelInfoReq, rsp *pb.ManageInfoRsp) error {
	// ConnDB 实例
	db := utils.ConnDB()
	// 判断 token，并获取用户名、用户角色
	tokenBol, _, role, tokenErr := logic.PreHandleTokenLogic(db, ctx)
	// 不正常时不允许请求 DB，并返回相应的错误码，清空用户信息
	if !tokenBol || tokenErr != nil {
		rsp.Code, rsp.Msg = config.ClientInvalidTokenError.Code, utils.GetErrorMap(tokenErr.Error()).Msg
		return nil
	}
	if role != 3 {
		rsp.Code, rsp.Msg = config.ClientPermissionError.Code, config.ClientPermissionError.Msg
		return nil
	}
	// 修改视频状态逻辑
	dbErr, _ := logic.DelInfoLogic(db, req)
	// 修改 DB 错误
	if dbErr != nil {
		rsp.Code, rsp.Msg = config.InnerWriteDbError.Code, config.InnerWriteDbError.Msg
		return nil
	}
	// 正常返回
	rsp.Code, rsp.Msg = config.ResOk.Code, config.ResOk.Msg

	return nil
}
