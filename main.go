package main

import (
	rb_app "reblog-server/app"
)

func main() {
	s := rb_app.NewServer()

	s.Start()
}
