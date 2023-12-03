package errcode

const (
	Success = 0
	Failed  = 1

	// common
	ParamParseFailed = 10
	RoleMismatch     = 11

	// auth
	LoginFailed   = 100
	UsernameOrPwd = 101
	SessionSave   = 102
	NotLogin      = 103

	// user
	GetUsersFailed   = 200
	UpdateUserFailed = 201
	DeleteUserFailed = 202
	AddUserFailed    = 203
)
