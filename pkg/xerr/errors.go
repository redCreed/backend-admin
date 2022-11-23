package xerr

import "fmt"

/**
自定义错误信息
*/

type CodeError struct {
	errCode int32
	errMsg  string
}

// GetCode  返回给前端的错误码
func (e *CodeError) GetCode() int32 {
	return e.errCode
}

// GetMsg 返回给前端显示端错误信息
func (e *CodeError) GetMsg() string {
	if msg, ok := errMap[e.errCode]; ok {
		return msg
	}
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

func NewErrCodeMsg(errCode int32, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errMsg}
}
func NewErrCode(errCode int32) *CodeError {
	return &CodeError{errCode: errCode, errMsg: ErrMapMsg(errCode)}
}

// NewErrMsg 通用数据库操作错误返回
func NewErrMsg(errMsg string) *CodeError {
	return &CodeError{errCode: DbErr, errMsg: errMsg}
}
