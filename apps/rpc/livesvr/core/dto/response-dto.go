package dto

var (
	OperationSuccess = Response{Status: 0, Info: "Operation Success"}
	InvalidRVID      = Response{Status: 90001, Info: "Invalid RVID"}
	InvalidUID       = Response{Status: 90002, Info: "Invalid UID"}
	NoPermission     = Response{Status: 90003, Info: "No Permission"}
	LiveNotExist     = Response{Status: 90004, Info: "Live Not Exist"}
)

// Response 业务层错误封装
type Response struct {
	Status int64  `json:"status"`
	Info   string `json:"info"`
}

func (r Response) Error() string {
	return r.Info
}

func (r Response) GetStatus() int64 {
	return r.Status
}

func (r Response) GetInfo() string {
	return r.Info
}

// ServerInternalError 服务器内部错误封装
func ServerInternalError(err error) Response {
	return Response{
		Status: 500,
		Info:   err.Error(),
	}
}
