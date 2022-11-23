package xerr

//全局错误码
const (
	OK               int32 = 0   //成功返回
	DbErr            int32 = 110 //服务器繁忙，稍后再试！
	ValidateParamErr int32 = 111 //校验参数错误
	NoPermission     int32 = 112 //无访问权限
	DataHasExist     int32 = 113 //数据已存在
	DataNoExist      int32 = 114 //数据不存在
	PasswordErr      int32 = 115 //密码错误
	TokenExpired     int32 = 120 //token已过期
	TokenNotValidYet int32 = 121 //token未激活
	TokenMalformed   int32 = 122 //非法的token
	TokenInvalid     int32 = 124 //无效的token
)

/**
    前3位代表业务,后三位代表具体功能:
	101xxx: 用户服务错误码
**/
const (
	UserHasExist int32 = 101001 //用户已存在
	UserNoExist  int32 = 101002 //用户不存在或已经被删除
)
