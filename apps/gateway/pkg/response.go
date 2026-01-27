package pkg

var (
	OperationSuccess = Response{Status: 20000, Info: "Operation Success"}         // 操作成功
	RevDataError     = Response{Status: 40000, Info: "Hertz Validate Data Error"} // api数据错误
)

// 业务层错误封装
type Response struct {
	Status uint   `json:"status"`
	Info   string `json:"info"`
}

func (r Response) Error() string {
	return r.Info
}

// 请求层错误封装
type FinalResponse struct {
	Status uint        `json:"status"`
	Info   string      `json:"info"`
	Data   interface{} `json:"data"`
}

// 服务器内部错误封装
func ServerInternalError(err error) Response {
	return Response{
		Status: 500,
		Info:   err.Error(),
	}
}
