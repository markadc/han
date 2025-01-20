package loger

import "testing"

func TestLoger(t *testing.T) {
	Info("hello world")
	Warning("hello world")
	Error("hello world")
	Success("hello world")
}
