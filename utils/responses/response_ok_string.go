package responses

type ResponseOkString struct {
	Message   string     `json:"message,omitempty"`
}


func NewResponseOkString(message string) ResponseOkString{
	return ResponseOkString{
		Message: message,
	}
}

