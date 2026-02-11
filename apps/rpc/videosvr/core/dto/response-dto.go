package dto

var (
	OperationSuccess   = Response{Status: 50200, Info: "Operation Success"}
	InvalidRVID        = Response{Status: 50001, Info: "Invalid RVID"}
	InvalidFaceUrl     = Response{Status: 50002, Info: "Invalid FaceUrl"}
	InvalidMinioKey    = Response{Status: 50003, Info: "Invalid MinioKey"}
	InvalidDescription = Response{Status: 50004, Info: "Invalid Description"}
	InvalidAuthorID    = Response{Status: 50005, Info: "Invalid AuthorID"}
	InvalidAuthorName  = Response{Status: 50006, Info: "Invalid AuthorName"}
	NoPermission       = Response{Status: 50007, Info: "No Permission"}
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
