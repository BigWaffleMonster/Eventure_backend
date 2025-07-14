package responses

type ResponseFailed struct {
	Message   string     `json:"message,omitempty"`
	Details   []string   `json:"details,omitempty"`
}

func NewResponseFailed(message string, details []string) ResponseFailed{
	return ResponseFailed{
		Message: message,
		Details: details,
	}
}