package errno

var (
	OK = NewError(0, "OK")

	ErrServer             = NewError(10001, "服务异常，请联系管理员")
	ErrParam              = NewError(10002, "参数有误")
	ErrSignParam          = NewError(10003, "签名参数有误")
	ErrForbiddenOperation = NewError(10004, "非法操作")

	ErrUserPhone                        = NewError(20101, "用户手机号不合法")
	ErrUserCaptcha                      = NewError(20102, "用户验证码有误")
	ErrUserPassword                     = NewError(20103, "用户密码不正确")
	ErrUserNotExist                     = NewError(20104, "用户不存在")
	ErrUserAuthorizationMissing         = NewError(20105, "未认证的用户")
	ErrUserAuthorizationFormat          = NewError(20106, "用户的认证格式不对")
	ErrUserUnsupportedAuthorizationType = NewError(20107, "不支持的认证类型")
	ErrUserInvalidToken                 = NewError(20108, "认证令牌有误")

	ErrAccountNotBelongToUser = NewError(20201, "账户不属于用户")
)
