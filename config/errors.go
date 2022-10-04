package config

import "git.code.oa.com/trpc-go/trpc-go/errs"

// Status code msg 返回状态
type Status struct {
	Code int32
	Msg  string
}

var (
	ResOk                    = Status{0, "Success"}
	ResFail                  = Status{-1, ""}
	InnerProcessingTimeout   = Status{-100, "内部处理超时"}
	InnerServiceOverload     = Status{-200, "内部服务过载"}
	InnerLogicError          = Status{-300, "内部逻辑出错"}
	InnerOtherError          = Status{-400, "内部其他问题"}
	InnerReadDbError         = Status{-401, "读取数据库错误"}
	InnerWriteDbError        = Status{-402, "写入数据库错误"}
	InnerDeleteDbError       = Status{-403, "删除数据库错误"}
	InnerMarshalError        = Status{-427, "数据序列化失败"}
	InnerUnmarshalError      = Status{-428, "数据反序列化失败"}
	ClientUserInfoError      = Status{100, ""}
	ClientParamParsingError  = Status{101, "客户端请求参数解析失败"}
	ClientInvalidParamError  = Status{102, "客户端请求参数校验失败"}
	ClientPermissionError    = Status{200, "客户端没有访问权限"}
	ClientCheckUserNameError = Status{302, "用户名重复"}
	ClientUPInvalid          = Status{303, "用户名或密码不规范"}
	ClientLoginError         = Status{304, "登录失败"}
	ClientNoTokenError       = Status{305, "客户端Token缺失"}
	ClientInvalidTokenError  = Status{306, "客户端Token无效"}
	ClientForgedIdentity     = Status{307, "客户端伪造身份"}
	ClientLoginExpired       = Status{308, "用户登陆过期"}
)

// New 错误构造方法
func New(err Status) error {
	return errs.New(int(err.Code), err.Msg)
}
