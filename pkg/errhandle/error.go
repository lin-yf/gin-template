package errhandle

import (
	"errors"

	"gorm.io/gorm"
)

var (
	TokenExpired     error = errors.New("Token is expired")
	TokenNotValidYet error = errors.New("Token not active yet")
	TokenMalformed   error = errors.New("That's not even a token")
	TokenInvalid     error = errors.New("Couldn't handle this token:")
)

// Response json返回格式
type Response struct {
	Code   int         `json:"code"`
	Result interface{} `json:"result,omitempty"`
	Msg    string      `json:"message"`
	Error  string      `json:"error,omitempty"`
}

// CommonError 通用报错格式
type CommonError struct {
	Code     int
	Msg      string
	RawError error
}

// 三位数错误编码为复用http原本含义
// 五位数错误编码为应用自定义错误
const (
	// CodeNotFullySuccess 未完全成功
	AppNotFullySuccess = 203
	// 未登录
	AppCheckLogin = 401
	// 未授权访问
	AppNoPermissionErr = 403
	// 资源未找到
	AppNotFound = 404
	// AppParamErr 参数错误
	AppParamErr = 40001
	// AppCredentialInvalid 凭证无效
	AppCredentialInvalid = 40001
	// 上传出错
	AppUploadFailed = 40002
	// 目录创建失败
	AppCreateFolderFailed = 40003
	// 数据库操作失败
	AppDBError = 40005
	// AppEncryptError
	AppEncryptError = 40006
	// IO操作失败
	AppIOFailed = 40007
	// 回调失败
	AppCallbackError = 40008
	// 未定错误
	AppUnknownErr = -1
)

// NewError 返回新的错误对象
func NewError(code int, msg string, err error) CommonError {
	return CommonError{
		Code:     code,
		Msg:      msg,
		RawError: err,
	}
}

// WithError
func (err *CommonError) WithError(raw error) CommonError {
	err.RawError = raw
	return *err
}

// Error 返回业务代码错误信息
func (err CommonError) Error() string {
	return err.Msg
}

// Err 通用错误处理
func Err(errCode int, msg string, err error) Response {
	if commonErr, ok := err.(CommonError); ok {
		errCode = commonErr.Code
		err = commonErr.RawError
		msg = commonErr.Msg
	}

	res := Response{
		Code: errCode,
		Msg:  msg,
	}
	// if err != nil && gin.Mode() != gin.ReleaseMode {
	if err != nil {
		// res.Error = err.Error()
	}
	return res
}

// Err 通用错误处理
func ErrWithData(errCode int, msg string, err error, data interface{}) Response {
	if commonErr, ok := err.(CommonError); ok {
		errCode = commonErr.Code
		err = commonErr.RawError
		msg = commonErr.Msg
	}

	res := Response{
		Code:   errCode,
		Msg:    msg,
		Result: data,
	}
	// if err != nil && gin.Mode() != gin.ReleaseMode {
	if err != nil {
		res.Error = err.Error()
	}
	return res
}

func IsNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
