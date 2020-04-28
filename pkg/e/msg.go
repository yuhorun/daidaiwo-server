package e

var MsgFlags = map[int]string{
	SUCCESS:               "ok",
	LOGIN_SUCCESS:         "登陆成功",
	LOGOUT_SUCCESS:        "退出登录成功",
	SIGNUP_SUCCESS:        "注册成功",
	GETTASKLIST_SUCCESS:   "获取任务列表成功",
	GETVERIFYCODE_SUCCESS: "获取验证码成功",

	FAIL:               "fail",
	LOGIN_FAIL:         "登陆失败",
	LOGOUT_FAIL:        "退出登陆失败",
	SIGNUP_FAIL:        "注册失败",
	GETTASKLIST_FAIL:   "读取任务列表失败",
	GETVERIFYCODE_FAIL: "获取验证码失败",

	INVALID_PARAMS: "请求参数错误",

	ERROR_COOKIE_NOT_SET:         "Cookie未设置",
	ERROR_AUTH_CHECK_JWT_FAIL:    "JWT鉴权失败",
	ERROR_AUTH_CHECK_JWT_TIMEOUT: "JWT已超时",
	ERROR_AUTH_JWT:               "JWT生成失败",
	ERROR_AUTH:                   "JWT错误",

	ERROR_SIGNUP:              "C错误，注册失败",
	ERROR_USERNAME_OR_PAWWORD: "用户手机号或者密码错误",

	ERROR_POST_C: "C错误，发布任务失败",
	ERROR_POST_M: "M错误，发布任务失败",

	ER_DUP_ENTRY:       "键值重复",
	ER_DUP_ENTRY_NAME:  "用户名称已存在",
	ER_DUP_ENTRY_PHONE: "手机号码已被注册过",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[FAIL]
}
