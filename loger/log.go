package loger

import "time"

// Now 返回当前时间的
func Now() string {
	currTime := time.Now().Format("2006-01-02 15:04:05")
	return currTime
}

// Default 开头加个时间
func Default(s string, values ...interface{}) string {
	return Now() + "  " + PyFormat(s, values...)
}

func Debug(s string, values ...interface{}) {
	Print(Default(s, values...), "blue")
}

func Info(s string, values ...interface{}) {
	Print(Default(s, values...), "")
}

func Warning(s string, values ...interface{}) {
	Print(Default(s, values...), "yellow")
}

func Error(s string, values ...interface{}) {
	Print(Default(s, values...), "red")
}

func Success(s string, values ...interface{}) {
	Print(Default(s, values...), "green")
}
