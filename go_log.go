package go_logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func GoLog(errMsg error, path string) {
	logPath := path + "/log.txt"

	// Opens or creates file.
	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}

	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {

		}
	}(logFile)

	// Gets the currentTime during the reported error message.
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	_, err = logFile.WriteString(fmt.Sprintf("%s: %s\n", formattedTime, errMsg))
	if err != nil {
		panic(err)
	}
}

func GetPath() string {
	// Get the absolute path of the executable
	execPath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
	}

	// Get the directory of the executable
	execDir := filepath.Dir(execPath)
	fmt.Println("Executable Directory:", execDir)

	return execDir
}
