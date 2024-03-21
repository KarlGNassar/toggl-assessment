package api

type ErrorResponse struct {
	ErrorCode    string `json:"error_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
	Details      string `json:"details,omitempty"`
}
