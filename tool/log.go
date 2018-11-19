package tool

import (
	"fmt"
	"os"
	"time"
)

func Writelog(txt string) {
	filename := "log/" + time.Now().Format("20060102") + ".log"
	logFile, logErr := os.OpenFile(filename, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if logErr != nil {
		fmt.Println("Fail to find", *logFile, "cServer start Failed")
		os.Exit(1)
	}
	logFile.WriteString(time.Now().Format("[15:04:05] ") + txt + "\r\n")
}
