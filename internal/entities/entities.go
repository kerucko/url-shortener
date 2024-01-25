package entities

type StatusResponse struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

type Request struct {
	Url   string `json:"url" validate:"required,url"`
	Alias string `json:"alias,omitempty"`
}

type Response struct {
	StatusResponse
	Alias string `json:"alias,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func OK() StatusResponse {
	return StatusResponse{
		Status: StatusOK,
	}
}

func Error(msg string) StatusResponse {
	return StatusResponse{
		Status: StatusError,
		Error:  msg,
	}
}
