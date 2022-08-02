package e

type ErrorReponse struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

var MsgFlags = map[int]string{
	SUCCESS			: "ok",
	ERROR			: "fail",
	InvalidParams   : "请求参数错误",
	ErrorExistUser: "the user has existed",
	ErrorNotExistUser: "the user have not existed",

	ErrorAuthCheckTokenFail:        "Token鉴权失败",
	ErrorAuthCheckTokenTimeout:     "Token已超时",
	ErrorAuthToken:                 "Token生成失败",
	ErrorAuth:                      "Token错误",
	ErrorNotCompare:                "不匹配",
	ErrorDatabase:                  "数据库操作出错,请重试",

}

func GetMsg(code int) string  {
	if msg,ok:=MsgFlags[code];ok{
		return msg
	}
	return MsgFlags[ERROR]
}