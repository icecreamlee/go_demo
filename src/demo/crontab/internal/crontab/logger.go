package crontab

import (
	"github.com/IcecreamLee/goutils"
	"log"
	"os"
)

// Program structures.
//  Define Start and Stop methods.
type Logger struct {
	OutFile string
}

var logger *log.Logger

func InitLogger() {
	var err error
	logFile := goutils.GetCurrentPath() + "go_crontab.log"
	loggerFile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	logger = log.New(loggerFile, "", log.LstdFlags)
}
