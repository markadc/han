package loger

import "testing"

func TestPrint(t *testing.T) {
	s1 := PyFormat("当前时间是：{}", Now())
	Print(s1, "blue")
	name, code := "王天风", "毒蜂"
	s2 := PyFormat("姓名：{}，代号：{}", name, code)
	Print(s2, "red")
}
