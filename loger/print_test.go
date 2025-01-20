package loger

import "testing"

func TestPrint(t *testing.T) {
	name := "wauo"
	flag := "626"
	Print("{} is {}", name, flag)
}
