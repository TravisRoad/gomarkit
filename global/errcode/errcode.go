package errcode

const (
	Success = 0
	Failed  = 1

	// auth
	LoginFailed   = 100
	UsernameOrPwd = 101
	SessionSave   = 102
	NotLogin      = 103
)
