package dto

type ErrorResponse struct {
	ErrMsg string `json:"error"`
}

func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		ErrMsg: message,
	}
}
