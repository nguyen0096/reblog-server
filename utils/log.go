package utils

import "log"

func Info(message string, params ...interface{}) {
	log.Printf(message, params...)
}

func Error(message string, params ...interface{}) {
	log.Printf(message, params...)
}
