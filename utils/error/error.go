package error

type Error struct {
	message string `json:"message"`
	err     error  `json:"error"`
}
