package log

import "time"

func Now() string {
	currentTime := time.Now().Format("2006-01-02 15:04:05")
	return currentTime
}

func HeadAddTime(s string, values ...interface{}) string {
	return Now() + "  " + FormatString(s, values...)
}

func Info(s string, values ...interface{}) {
	ColorPrint(HeadAddTime(s, values...), "")
}

func Warning(s string, values ...interface{}) {
	ColorPrint(HeadAddTime(s, values...), "yellow")

}

func Error(s string, values ...interface{}) {
	ColorPrint(HeadAddTime(s, values...), "red")
}

func Success(s string, values ...interface{}) {
	ColorPrint(HeadAddTime(s, values...), "green")
}
