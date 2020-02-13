package log_utils

import (
	"log"
	"os"
)

var (
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	Debug = log.New(os.Stdout, "[DEBUG] ", log.LstdFlags|log.Lshortfile)
	Info = log.New(os.Stdout, "[INFO] ", log.LstdFlags|log.Lshortfile)
	Warning = log.New(os.Stdout, "[WARNING] ", log.LstdFlags|log.Lshortfile)
	Error = log.New(os.Stdout, "[ERROR] ", log.LstdFlags|log.Lshortfile)
}
