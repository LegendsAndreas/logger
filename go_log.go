package go_logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// var logName string
// var pathName string

// func SetLogFile(logName string, path string) {
// 	logName = logName
// 	pathName = path
// }

func goLog(errMsg error, errType string, logName string, path string) {
	// Formates absolute path and log file name.
	logPath := formatLogPath(logName, path)

	logFile, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file.")
		fmt.Println(err)
	}

	defer func(logFile *os.File) {
		err := logFile.Close()
		if err != nil {
			fmt.Println("Error closing file.")
			fmt.Println(err)
		}
	}(logFile)

	// Gets the currentTime during the reported error message.
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")

	// Writes to log file with the format: [DATE]: [ERROR TYPE] [ERROR MESSAGE]
	_, err = logFile.WriteString(fmt.Sprintf("%s: [%s] %s\n", formattedTime, errType, errMsg))
	if err != nil {
		fmt.Println("Error writing to log file.")
		fmt.Println(err)
	}
}

// Formats absolute path for where the log file will be saved.
func formatLogPath(logName string, path string) string {
	// If logName is empty, we set it to "log".
	if logName == "" {
		logName = "Log"
	}

	// Sets the log path.
	formattedLogName := "/" + logName + ".txt"
	logPath := path + formattedLogName

	return logPath
}

// Calls the function "goLog()" and logs an error with the type "INFO".
func Info(errMsg error, logName string, pathName string) {
	goLog(errMsg, "INFO", logName, pathName)
}

// Calls the function "goLog()" and logs an error with the type "WARNING".
func Warning(errMsg error, logName string, pathName string) {
	goLog(errMsg, "WARNING", logName, pathName)
}

// Calls the function "goLog()" and logs an error with the type "ERROR".
func Error(errMsg error, logName string, pathName string) {
	goLog(errMsg, "ERROR", logName, pathName)
}

// Gets the absolute path of the directory, where your program executes.
//
// Remember that if you just use "go run file.go", the code will execute in your GOPATH and not where you have your program is.
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
