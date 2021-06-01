package main

import (
	"flag"
)

type Flags struct {
	Port  int64
	Debug bool
}

func (f *Flags) ParseAgrs() {
	flag.Int64Var(&f.Port, "port", 14070, "server port")
	flag.BoolVar(&f.Debug, "debug", false, "debug")
}
