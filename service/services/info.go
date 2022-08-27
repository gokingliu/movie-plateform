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
	tokenBol, _, _, _ := logic.PreHandleTokenLogic(db, ctx)
	// 查询视频详情逻辑
	result, err := logic.GetInfoLogic(db, req.Mid)
	// 客户端接口，判断 token 解析是否正常
	// 不正常时返回相应的错误码，同时返回视频详情，前端正常展示，但清空用户信息
	if !tokenBol || err != nil {
		rsp.Code, rsp.Msg = config.ClientUserInfoError.Code, config.ClientUserInfoError.Msg
	} else {
		rsp.Code, rsp.Msg = config.ResOk.Code, config.ResOk.Msg
	}
	rsp.Result = result

	return nil
}

// GetRecord 获取记录
func (s *InfoImpl) GetRecord(ctx context.Context, req *pb.RecordReq, rsp *pb.RecordRsp) error {
	// ConnDB 实例
	db := utils.ConnDB()
	// 判断 token，并获取用户名、用户角色
	tokenBol, userName, _, tokenErr := logic.PreHandleTokenLogic(db, ctx)
	// 查询记录逻辑
	result, err := logic.GetRecordLogic(db, userName, req.Mid, req.MType)
	// 客户端接口，判断 token 解析是否正常
	// 不正常时返回相应的错误码，清空用户信息，并返回 false
	if !tokenBol || tokenErr != nil || err != nil {
		rsp.Code, rsp.Msg = config.ClientUserInfoError.Code, config.ClientUserInfoError.Msg
		rsp.Result = false
	} else {
		rsp.Code, rsp.Msg = config.ResOk.Code, config.ResOk.Msg
		rsp.Result = result
	}

	return nil
}
