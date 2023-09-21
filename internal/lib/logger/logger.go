package logger

import (
	"log"
	"os"
)

var (
	INFO = log.New(os.Stdout, "[INFO] ", log.Ltime)
	WARN = log.New(os.Stdout, "[WARNING] ", log.Ltime|log.Llongfile)
	ERR  = log.New(os.Stderr, "[ERROR] ", log.Ltime|log.Llongfile)
)
