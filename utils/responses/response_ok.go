package responses

type ResponseOk[T any] struct {
	Message   string     `json:"message,omitempty"`
	Data      *T         `json:"data,omitempty"`
}

func NewResponseOk[T any](data *T, message string) ResponseOk[T]{
	return ResponseOk[T]{
		Message: message,
		Data: data,
	}
}
