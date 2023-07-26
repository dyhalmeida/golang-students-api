package controller

type ResponseMessage struct {
	Message string `json:"message"`
}

func NewResponseMessage(msg string) ResponseMessage {
	return ResponseMessage{
		Message: msg,
	}
}
