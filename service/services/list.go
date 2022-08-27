package services

import (
	"MovieService/config"
	"MovieService/logic"
	"MovieService/utils"
	"context"
	"encoding/json"
	pb "git.woa.com/crotaliu/pb-hub"
)

type ListImpl struct{}

// GetList 获取视频列表
func (s *ListImpl) GetList(ctx context.Context, req *pb.GetListReq, rsp *pb.GetListRsp) error {
	// ConnDB 实例
	db := utils.ConnDB()
	// 判断 token，并获取用户名、用户角色
	tokenBol, _, role, tokenErr := logic.PreHandleTokenLogic(db, ctx)
	// struct 转 map
	mapData := make(map[string]interface{})
	reqData, marshalErr := json.Marshal(&req)
	if marshalErr != nil {
		rsp.Code, rsp.Msg = config.InnerMarshalError.Code, config.InnerMarshalError.Msg
		return nil
	}
	unMarshalErr := json.Unmarshal(reqData, &mapData)
	if unMarshalErr != nil {
		rsp.Code, rsp.Msg = config.InnerUnmarshalError.Code, config.InnerUnmarshalError.Msg
		return nil
	}
	// 查询电影列表逻辑
	result, count, dbErr := logic.GetListLogic(db, mapData, role, req.PageNo, req.PageSize)
	if dbErr != nil {
		rsp.Code, rsp.Msg = config.InnerReadDbError.Code, config.InnerReadDbError.Msg
		return nil
	}
	// 客户端接口，判断 token 解析是否正常
	// 不正常时返回相应的错误码，同时返回列表信息，前端正常展示列表，但清空用户信息
	if !tokenBol {
		rsp.Code, rsp.Msg = config.ClientUserInfoError.Code, utils.GetErrorMap(tokenErr.Error()).Msg
	} else {
		rsp.Code, rsp.Msg = config.ResOk.Code, config.ResOk.Msg
	}
	rsp.Result = &pb.GetListRsp_Result{
		List:  result,
		Count: count,
	}

	return nil
}

// GetLeaderboard 获取视频排行榜
func (s *ListImpl) GetLeaderboard(ctx context.Context, req *pb.GetLeaderboardReq, rsp *pb.GetLeaderboardRsp) error {
	// ConnDB 实例
	db := utils.ConnDB()
	// 判断 token，并获取用户名、用户角色
	tokenBol, _, _, tokenErr := logic.PreHandleTokenLogic(db, ctx)
	// 查询视频排行榜逻辑
	result, dbErr := logic.GetLeaderboardLogic(db, req.MType)
	if dbErr != nil {
		rsp.Code, rsp.Msg = config.InnerReadDbError.Code, config.InnerReadDbError.Msg
		return nil
	}
	// 客户端接口，判断 token 解析是否正常
	// 不正常时返回相应的错误码，同时返回列表信息，前端正常展示列表，但清空用户信息
	if !tokenBol {
		rsp.Code, rsp.Msg = config.ClientUserInfoError.Code, utils.GetErrorMap(tokenErr.Error()).Msg
	} else {
		rsp.Code, rsp.Msg = config.ResOk.Code, config.ResOk.Msg
	}
	rsp.Result = result

	return nil
}
