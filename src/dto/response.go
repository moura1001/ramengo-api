package dto

type ErrorResponse struct {
	ErrMsg string `json:"error"`
}

func NewErrorResponse(message string) ErrorResponse {
	return ErrorResponse{
		ErrMsg: message,
	}
}

type OrderResponse struct {
	Id          string `json:"id"`
	Description string `json:"description"`
	Image       string `json:"image"`
}
