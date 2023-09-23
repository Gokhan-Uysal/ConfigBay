package logger

import (
	"log"
	"os"
)

var (
	INFO  = log.New(os.Stdout, "[INFO] ", log.Ltime)
	DEBUG = log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime|log.Llongfile)
	WARN  = log.New(os.Stdout, "[WARNING] ", log.Ltime|log.Llongfile)
	ERR   = log.New(os.Stderr, "[ERROR] ", log.Ltime|log.Llongfile)
)
