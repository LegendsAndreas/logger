package go_logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Info: Yellow
// Warning: Orange.
// Failiure: Red.

// var logName string
// var pathName string

// func SetLogFile(logName string, path string) {
// 	logName = logName
// 	pathName = path
// }

/*
* @brief Takes an error message and the path where that error happend and stores it in "log.txt".
*
* This function opens or creates a log.txt file in append write only mode, and stores the error "errMsg"
* in the appropriate path, provided by "path". It gets the current time and formates it to YYYY-MM-DD HH:MM:SS.
* The time along with the error message then gets written into the log file.
*
* @param errMsg The error message that will be logged.
* @param logName The name of the log file. Default name is "Log". If you want the default name, leave this as empty("").
* @param path The path of where the error occured.
*
 */
func GoLog(errMsg error, errType string, logName string, path string) {
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

// Calls the function "GoLog()" and logs an error with the type "INFO".
func Info(errMsg error, logName string, pathName string) {
	GoLog(errMsg, "INFO", logName, pathName)
}

// Calls the function "GoLog()" and logs an error with the type "WARNING".
func Warning(errMsg error, logName string, pathName string) {
	GoLog(errMsg, "WARNING", logName, pathName)
}

// Calls the function "GoLog()" and logs an error with the type "ERROR".
func Error(errMsg error, logName string, pathName string) {
	GoLog(errMsg, "ERROR", logName, pathName)
}

/*
* @brief GetPath returns the absolute path of the executable.
*
* @note Remember that if you just run your Go program with "go run main.go", it wont make
* 		store the log file in the program directory, since the execution of the program
*		happens in your GOPATH folder.
*
* @returns A string of the directory of the executable.
 */
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
