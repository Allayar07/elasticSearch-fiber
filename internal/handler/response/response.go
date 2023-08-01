package response

type RespMessage struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

func NewResponseMessage(code int, message interface{}) RespMessage {
	return RespMessage{
		Code:    code,
		Message: message,
	}
}
