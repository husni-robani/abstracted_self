package logger

import (
	"log"
	"os"
)

var Error = log.New(os.Stderr, "[ERROR] ", (log.Ldate | log.Lshortfile))
var Info = log.New(os.Stdout, "[INFO] ", (log.Lmsgprefix))
var Debug = log.New(os.Stdout, "[DEBUG] ", (log.Lmsgprefix))