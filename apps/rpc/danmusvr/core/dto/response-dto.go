package dto

var (
	OperationSuccess = Response{Status: 0, Info: "Operation Success"}
	InvalidRoomID    = Response{Status: 40001, Info: "Invalid RoomID"}
	InvalidUserID    = Response{Status: 40002, Info: "Invalid UserID"}
	InvalidColor     = Response{Status: 40003, Info: "Invalid Color"}
	InvalidContent   = Response{Status: 40004, Info: "Invalid Content"}
	InvalidDanID     = Response{Status: 40005, Info: "Invalid DanID"}
	DanmuNotExist    = Response{Status: 40006, Info: "Danmu Not Exist"}
	NoPermission     = Response{Status: 40007, Info: "No Permission"}
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
