package e

const (
	CodeSuccess       = 200
	CodeInvalidParams = 400
	CodeError         = 500
)

var Code2Msg = map[int]string{
	CodeError:         "fail",
	CodeSuccess:       "ok",
	CodeInvalidParams: "请求参数有误",
}

func GetMsg(code int) string {
	if msg, ok := Code2Msg[code]; ok {
		return msg
	}
	return Code2Msg[CodeError]
}
