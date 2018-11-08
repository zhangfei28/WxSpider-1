package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",
	INVALID_MP:     "公众号不存在",
}

//获取错误消息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]

	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
