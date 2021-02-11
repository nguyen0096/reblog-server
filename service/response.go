package service

type Error struct {
	Code     int    `json:"code"`
	Message  string `json:"message"`
	Detail   string `json:"detail"`
	Where    string `json:"where"`
	RawError error  `json:""`
}

func (c *Error) Error() string {
	return c.Message
}

type Response struct {
	Message      string `json:"message"`
	RowsAffected int    `json:"rows_affected"`
}
