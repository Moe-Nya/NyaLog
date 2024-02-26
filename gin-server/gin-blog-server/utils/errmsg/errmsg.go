package errmsg

const (
	SUCCESS = 200
	ERROR   = 500
)

var codeMsg = map[int]string{
	SUCCESS: "ok",
	ERROR:   "error",
}

func GetErrorMsg(code int) string {
	return codeMsg[code]
}
