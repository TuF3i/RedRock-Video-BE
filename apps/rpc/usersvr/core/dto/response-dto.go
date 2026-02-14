package dto

var (
	OperationSuccess    = Response{Status: 0, Info: "Operation Success"}
	NoUserExist         = Response{Status: 80001, Info: "No User Exist"}
	InvalidRefreshToken = Response{Status: 80002, Info: "Invalid RefreshToken"}
	InvalidUID          = Response{Status: 80003, Info: "Invalid UID"}
	InvalidUserName     = Response{Status: 80004, Info: "Invalid User Name"}
	InvalidAvatarURL    = Response{Status: 80005, Info: "Invalid AvatarURL"}
	InvalidBio          = Response{Status: 80006, Info: "Invalid Bio"}
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
