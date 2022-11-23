package common

type RouteOption func()

var (
	PTypeDirectory = 1 //目录类型

	PTypeMenu = 2 //菜单类型

	PTypeButton = 3 //按钮类型
)

var (
	Admin = "admin"
)

const (
	LoginLog   = "login_log_queue"
	OperateLog = "operate_log_queue"
	ApiCheck   = "api_check_queue"
)
