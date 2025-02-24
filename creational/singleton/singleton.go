package singleton

import (
	"fmt"
	"sync"
)

type Logger struct{}

var instance *Logger
var once sync.Once

func GetLoggerInstance() *Logger {
	once.Do(func() {
		instance = &Logger{}
	})
	return instance
}

func DemoSingleton() {
	logger1 := GetLoggerInstance()
	logger2 := GetLoggerInstance()
	fmt.Println(logger1 == logger2) // true
}
