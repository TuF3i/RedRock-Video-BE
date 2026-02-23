package response

type FinalResponse struct {
	Status int64       `json:"status"`
	Info   string      `json:"info"`
	Data   interface{} `json:"data"`
}

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

func InternalError(err error) Response {
	return Response{
		Status: 500,
		Info:   err.Error(),
	}
}
