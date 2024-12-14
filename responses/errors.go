package responses

type ResponseError struct {
	ErrorMessage string `json:"error_message,omitempty"`
}

func (r *ResponseError) SetErrorMessage(message string) {
	r.ErrorMessage = message
}

func (r *ResponseError) GetErrorMessage() string {
	return r.ErrorMessage
}
