package utils

import (
	"log"
)

// Info ...
func Info(message string, params ...interface{}) {
	// pc := make([]uintptr, 10) // at least 1 entry needed
	// runtime.Callers(5, pc)
	// f := runtime.FuncForPC(pc[0])
	// file, line := f.FileLine(pc[0])
	// fmt.Printf("%s:%d %s\n", file, line, f.Name())
	log.Printf("[info] "+message, params...)
}

// Error ...
func Error(message string, params ...interface{}) {
	log.Printf("[error] "+message, params...)
}

// Panic ...
func Panic(message string, params ...interface{}) {
	log.Panicf("[panic] "+message, params)
}
