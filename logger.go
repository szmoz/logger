package logger

import "fmt"

func Info(message string, params interface{}) {
	fmt.Println("Log info message: ", message, "params: ", params)
}

func Error(message string, params interface{}) {
	fmt.Println("Log error message: ", message, "params: ", params)
}

func Warning(message string, params interface{}) {
	fmt.Println("Log warning message: ", message, "params: ", params)
}

func Debug(message string, params interface{}) {
	fmt.Println("Log debug message: ", message, "params: ", params)
}

func Fatal(message string, params interface{}) {
	fmt.Println("Lof fatal message: ", message, "params: ", params)
}
