package logger

import "fmt"

func Log(v ...interface{}) {
	fmt.Println("log: ", v)
}

func Debug(v ...interface{}) {
	fmt.Println("debug: ", v)
}

func Warn(v ...interface{}) {
	fmt.Println("warn: ", v)
}

func Error(v ...interface{}) {
	fmt.Println("error: ",v)
}