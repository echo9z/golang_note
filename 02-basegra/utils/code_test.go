package utils

import "testing"

func TestCoding(t *testing.T) {
	res := Coding("hello")
	if res != "hellotest" {
		t.Error("Coding函数测试失败")
	}
	t.Log("Coding函数测试通过")
}
