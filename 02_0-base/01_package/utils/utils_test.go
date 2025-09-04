package utils

import "testing"

// 测试用例
func TestAdd(t *testing.T) {
	res := Add(3, 7)
	if res != 10 {
		t.Error("Coding函数测试失败")
	}
	t.Log("Coding函数测试通过")
}
