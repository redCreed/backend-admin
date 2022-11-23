package xerr

var errMap map[int32]string

func init() {
	errMap = make(map[int32]string)
	errMap[OK] = "SUCCESS"
	errMap[DbErr] = "服务器繁忙，稍后再试！"
	errMap[ValidateParamErr] = "参数非法"
	errMap[NoPermission] = "无访问权限"
	errMap[DataHasExist] = "数据已存在"
	errMap[DataNoExist] = "数据不存在"
	errMap[PasswordErr] = "密码错误"
	errMap[TokenExpired] = "token已过期"
	errMap[TokenNotValidYet] = "token未激活"
	errMap[TokenMalformed] = "非法的token"
	errMap[TokenInvalid] = "无效的token"

	errMap[UserHasExist] = "用户已存在"
	errMap[UserNoExist] = "该用户不存在或已被删除"
}

func ErrMapMsg(errCode int32) string {
	if msg, ok := errMap[errCode]; ok {
		return msg
	} else {
		return "服务器繁忙，稍后再试！"
	}
}
