package handle

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Data    interface{} `json:"data"`
}

func NewResponse(code int, status, message string, data interface{}) *Response {
	return &Response{
		Code:    code,
		Status:  status,
		Message: message,
		Data:    data,
	}
}
