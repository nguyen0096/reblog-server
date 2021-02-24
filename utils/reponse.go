package utils

type AppError struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Detail   string `json:"detail"`
	Where    string `json:"where"`
	RawError error  `json:""`
}

func (c *AppError) Error() string {
	return c.Message
}

type Response struct {
	Message      string `json:"message"`
	RowsAffected int    `json:"rows_affected"`
}

// Result new service response
type Result struct {
	Message string `json:"message"`
	Error   error  `json:"error"`
}
