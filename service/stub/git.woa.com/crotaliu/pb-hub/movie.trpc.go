// Code generated by trpc-go/trpc-go-cmdline. DO NOT EDIT.
// source: movie.proto

package pb_hub

import (
	"context"
	"fmt"

	_ "git.code.oa.com/trpc-go/trpc-go"
	_ "git.code.oa.com/trpc-go/trpc-go/http"

	"git.code.oa.com/trpc-go/trpc-go/client"
	"git.code.oa.com/trpc-go/trpc-go/codec"
	"git.code.oa.com/trpc-go/trpc-go/server"
)

/* ************************************ Service Definition ************************************ */

// UserService defines service
type UserService interface {

	// Register 用户注册
	Register(ctx context.Context, req *RegisterReq, rsp *RegisterRsp) (err error)

	// CheckUserName 检查用户名重复
	CheckUserName(ctx context.Context, req *CheckUserNameReq, rsp *CheckUserNameRsp) (err error)

	// Login 用户登录
	Login(ctx context.Context, req *LoginReq, rsp *LoginRsp) (err error)
}

func UserService_Register_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {

	req := &RegisterReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(UserService).Register(ctx, reqbody.(*RegisterReq), rspbody.(*RegisterRsp))
	}

	rsp := &RegisterRsp{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func UserService_CheckUserName_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {

	req := &CheckUserNameReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(UserService).CheckUserName(ctx, reqbody.(*CheckUserNameReq), rspbody.(*CheckUserNameRsp))
	}

	rsp := &CheckUserNameRsp{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func UserService_Login_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {

	req := &LoginReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(UserService).Login(ctx, reqbody.(*LoginReq), rspbody.(*LoginRsp))
	}

	rsp := &LoginRsp{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

// UserServer_ServiceDesc descriptor for server.RegisterService
var UserServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "trpc.MovieService.operation.User",
	HandlerType: ((*UserService)(nil)),
	Methods: []server.Method{
		{
			Name: "/trpc.MovieService.operation.User/Register",
			Func: UserService_Register_Handler,
		},
		{
			Name: "/trpc.MovieService.operation.User/CheckUserName",
			Func: UserService_CheckUserName_Handler,
		},
		{
			Name: "/trpc.MovieService.operation.User/Login",
			Func: UserService_Login_Handler,
		},
	},
}

// RegisterUserService register service
func RegisterUserService(s server.Service, svr UserService) {
	if err := s.Register(&UserServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("User register error:%v", err))
	}

}

// ListService defines service
type ListService interface {

	// GetList 获取视频列表
	GetList(ctx context.Context, req *GetListReq, rsp *GetListRsp) (err error)

	// GetLeaderboard 获取视频排行榜
	GetLeaderboard(ctx context.Context, req *GetLeaderboardReq, rsp *GetLeaderboardRsp) (err error)
}

func ListService_GetList_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {

	req := &GetListReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(ListService).GetList(ctx, reqbody.(*GetListReq), rspbody.(*GetListRsp))
	}

	rsp := &GetListRsp{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func ListService_GetLeaderboard_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {

	req := &GetLeaderboardReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(ListService).GetLeaderboard(ctx, reqbody.(*GetLeaderboardReq), rspbody.(*GetLeaderboardRsp))
	}

	rsp := &GetLeaderboardRsp{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

// ListServer_ServiceDesc descriptor for server.RegisterService
var ListServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "trpc.MovieService.operation.List",
	HandlerType: ((*ListService)(nil)),
	Methods: []server.Method{
		{
			Name: "/trpc.MovieService.operation.List/GetList",
			Func: ListService_GetList_Handler,
		},
		{
			Name: "/trpc.MovieService.operation.List/GetLeaderboard",
			Func: ListService_GetLeaderboard_Handler,
		},
	},
}

// RegisterListService register service
func RegisterListService(s server.Service, svr ListService) {
	if err := s.Register(&ListServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("List register error:%v", err))
	}

}

// InfoService defines service
type InfoService interface {

	// GetInfo 获取视频详情
	GetInfo(ctx context.Context, req *GetInfoReq, rsp *GetInfoRsp) (err error)

	// GetRecord 获取记录
	GetRecord(ctx context.Context, req *RecordReq, rsp *RecordRsp) (err error)

	// PostRecord 增加记录
	PostRecord(ctx context.Context, req *RecordReq, rsp *RecordRsp) (err error)

	// DelRecord 删除记录
	DelRecord(ctx context.Context, req *RecordReq, rsp *RecordRsp) (err error)
}

func InfoService_GetInfo_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {

	req := &GetInfoReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(InfoService).GetInfo(ctx, reqbody.(*GetInfoReq), rspbody.(*GetInfoRsp))
	}

	rsp := &GetInfoRsp{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func InfoService_GetRecord_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {

	req := &RecordReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(InfoService).GetRecord(ctx, reqbody.(*RecordReq), rspbody.(*RecordRsp))
	}

	rsp := &RecordRsp{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func InfoService_PostRecord_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {

	req := &RecordReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(InfoService).PostRecord(ctx, reqbody.(*RecordReq), rspbody.(*RecordRsp))
	}

	rsp := &RecordRsp{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

func InfoService_DelRecord_Handler(svr interface{}, ctx context.Context, f server.FilterFunc) (interface{}, error) {

	req := &RecordReq{}
	filters, err := f(req)
	if err != nil {
		return nil, err
	}
	handleFunc := func(ctx context.Context, reqbody interface{}, rspbody interface{}) error {
		return svr.(InfoService).DelRecord(ctx, reqbody.(*RecordReq), rspbody.(*RecordRsp))
	}

	rsp := &RecordRsp{}
	err = filters.Handle(ctx, req, rsp, handleFunc)
	if err != nil {
		return nil, err
	}

	return rsp, nil
}

// InfoServer_ServiceDesc descriptor for server.RegisterService
var InfoServer_ServiceDesc = server.ServiceDesc{
	ServiceName: "trpc.MovieService.operation.Info",
	HandlerType: ((*InfoService)(nil)),
	Methods: []server.Method{
		{
			Name: "/trpc.MovieService.operation.Info/GetInfo",
			Func: InfoService_GetInfo_Handler,
		},
		{
			Name: "/trpc.MovieService.operation.Info/GetRecord",
			Func: InfoService_GetRecord_Handler,
		},
		{
			Name: "/trpc.MovieService.operation.Info/PostRecord",
			Func: InfoService_PostRecord_Handler,
		},
		{
			Name: "/trpc.MovieService.operation.Info/DelRecord",
			Func: InfoService_DelRecord_Handler,
		},
	},
}

// RegisterInfoService register service
func RegisterInfoService(s server.Service, svr InfoService) {
	if err := s.Register(&InfoServer_ServiceDesc, svr); err != nil {
		panic(fmt.Sprintf("Info register error:%v", err))
	}

}

/* ************************************ Client Definition ************************************ */

// UserClientProxy defines service client proxy
type UserClientProxy interface {

	// Register 用户注册
	Register(ctx context.Context, req *RegisterReq, opts ...client.Option) (rsp *RegisterRsp, err error)

	// CheckUserName 检查用户名重复
	CheckUserName(ctx context.Context, req *CheckUserNameReq, opts ...client.Option) (rsp *CheckUserNameRsp, err error)

	// Login 用户登录
	Login(ctx context.Context, req *LoginReq, opts ...client.Option) (rsp *LoginRsp, err error)
}

type UserClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewUserClientProxy = func(opts ...client.Option) UserClientProxy {
	return &UserClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *UserClientProxyImpl) Register(ctx context.Context, req *RegisterReq, opts ...client.Option) (*RegisterRsp, error) {

	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)

	msg.WithClientRPCName("/trpc.MovieService.operation.User/Register")
	msg.WithCalleeServiceName(UserServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("MovieService")
	msg.WithCalleeServer("operation")
	msg.WithCalleeService("User")
	msg.WithCalleeMethod("Register")
	msg.WithSerializationType(codec.SerializationTypePB)

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &RegisterRsp{}

	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}

	return rsp, nil
}

func (c *UserClientProxyImpl) CheckUserName(ctx context.Context, req *CheckUserNameReq, opts ...client.Option) (*CheckUserNameRsp, error) {

	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)

	msg.WithClientRPCName("/trpc.MovieService.operation.User/CheckUserName")
	msg.WithCalleeServiceName(UserServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("MovieService")
	msg.WithCalleeServer("operation")
	msg.WithCalleeService("User")
	msg.WithCalleeMethod("CheckUserName")
	msg.WithSerializationType(codec.SerializationTypePB)

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &CheckUserNameRsp{}

	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}

	return rsp, nil
}

func (c *UserClientProxyImpl) Login(ctx context.Context, req *LoginReq, opts ...client.Option) (*LoginRsp, error) {

	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)

	msg.WithClientRPCName("/trpc.MovieService.operation.User/Login")
	msg.WithCalleeServiceName(UserServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("MovieService")
	msg.WithCalleeServer("operation")
	msg.WithCalleeService("User")
	msg.WithCalleeMethod("Login")
	msg.WithSerializationType(codec.SerializationTypePB)

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &LoginRsp{}

	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}

	return rsp, nil
}

// ListClientProxy defines service client proxy
type ListClientProxy interface {

	// GetList 获取视频列表
	GetList(ctx context.Context, req *GetListReq, opts ...client.Option) (rsp *GetListRsp, err error)

	// GetLeaderboard 获取视频排行榜
	GetLeaderboard(ctx context.Context, req *GetLeaderboardReq, opts ...client.Option) (rsp *GetLeaderboardRsp, err error)
}

type ListClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewListClientProxy = func(opts ...client.Option) ListClientProxy {
	return &ListClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *ListClientProxyImpl) GetList(ctx context.Context, req *GetListReq, opts ...client.Option) (*GetListRsp, error) {

	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)

	msg.WithClientRPCName("/trpc.MovieService.operation.List/GetList")
	msg.WithCalleeServiceName(ListServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("MovieService")
	msg.WithCalleeServer("operation")
	msg.WithCalleeService("List")
	msg.WithCalleeMethod("GetList")
	msg.WithSerializationType(codec.SerializationTypePB)

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &GetListRsp{}

	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}

	return rsp, nil
}

func (c *ListClientProxyImpl) GetLeaderboard(ctx context.Context, req *GetLeaderboardReq, opts ...client.Option) (*GetLeaderboardRsp, error) {

	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)

	msg.WithClientRPCName("/trpc.MovieService.operation.List/GetLeaderboard")
	msg.WithCalleeServiceName(ListServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("MovieService")
	msg.WithCalleeServer("operation")
	msg.WithCalleeService("List")
	msg.WithCalleeMethod("GetLeaderboard")
	msg.WithSerializationType(codec.SerializationTypePB)

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &GetLeaderboardRsp{}

	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}

	return rsp, nil
}

// InfoClientProxy defines service client proxy
type InfoClientProxy interface {

	// GetInfo 获取视频详情
	GetInfo(ctx context.Context, req *GetInfoReq, opts ...client.Option) (rsp *GetInfoRsp, err error)

	// GetRecord 获取记录
	GetRecord(ctx context.Context, req *RecordReq, opts ...client.Option) (rsp *RecordRsp, err error)

	// PostRecord 增加记录
	PostRecord(ctx context.Context, req *RecordReq, opts ...client.Option) (rsp *RecordRsp, err error)

	// DelRecord 删除记录
	DelRecord(ctx context.Context, req *RecordReq, opts ...client.Option) (rsp *RecordRsp, err error)
}

type InfoClientProxyImpl struct {
	client client.Client
	opts   []client.Option
}

var NewInfoClientProxy = func(opts ...client.Option) InfoClientProxy {
	return &InfoClientProxyImpl{client: client.DefaultClient, opts: opts}
}

func (c *InfoClientProxyImpl) GetInfo(ctx context.Context, req *GetInfoReq, opts ...client.Option) (*GetInfoRsp, error) {

	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)

	msg.WithClientRPCName("/trpc.MovieService.operation.Info/GetInfo")
	msg.WithCalleeServiceName(InfoServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("MovieService")
	msg.WithCalleeServer("operation")
	msg.WithCalleeService("Info")
	msg.WithCalleeMethod("GetInfo")
	msg.WithSerializationType(codec.SerializationTypePB)

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &GetInfoRsp{}

	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}

	return rsp, nil
}

func (c *InfoClientProxyImpl) GetRecord(ctx context.Context, req *RecordReq, opts ...client.Option) (*RecordRsp, error) {

	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)

	msg.WithClientRPCName("/trpc.MovieService.operation.Info/GetRecord")
	msg.WithCalleeServiceName(InfoServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("MovieService")
	msg.WithCalleeServer("operation")
	msg.WithCalleeService("Info")
	msg.WithCalleeMethod("GetRecord")
	msg.WithSerializationType(codec.SerializationTypePB)

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &RecordRsp{}

	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}

	return rsp, nil
}

func (c *InfoClientProxyImpl) PostRecord(ctx context.Context, req *RecordReq, opts ...client.Option) (*RecordRsp, error) {

	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)

	msg.WithClientRPCName("/trpc.MovieService.operation.Info/PostRecord")
	msg.WithCalleeServiceName(InfoServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("MovieService")
	msg.WithCalleeServer("operation")
	msg.WithCalleeService("Info")
	msg.WithCalleeMethod("PostRecord")
	msg.WithSerializationType(codec.SerializationTypePB)

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &RecordRsp{}

	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}

	return rsp, nil
}

func (c *InfoClientProxyImpl) DelRecord(ctx context.Context, req *RecordReq, opts ...client.Option) (*RecordRsp, error) {

	ctx, msg := codec.WithCloneMessage(ctx)
	defer codec.PutBackMessage(msg)

	msg.WithClientRPCName("/trpc.MovieService.operation.Info/DelRecord")
	msg.WithCalleeServiceName(InfoServer_ServiceDesc.ServiceName)
	msg.WithCalleeApp("MovieService")
	msg.WithCalleeServer("operation")
	msg.WithCalleeService("Info")
	msg.WithCalleeMethod("DelRecord")
	msg.WithSerializationType(codec.SerializationTypePB)

	callopts := make([]client.Option, 0, len(c.opts)+len(opts))
	callopts = append(callopts, c.opts...)
	callopts = append(callopts, opts...)

	rsp := &RecordRsp{}

	if err := c.client.Invoke(ctx, req, rsp, callopts...); err != nil {
		return nil, err
	}

	return rsp, nil
}
