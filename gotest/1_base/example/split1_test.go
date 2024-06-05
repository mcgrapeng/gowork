package example

import (
	"reflect"
	"testing"
)

//func TestSplit1(t *testing.T) { // 测试函数名必须以Test开头，必须接收一个*testing.T类型参数
//	got := Split1("a:b:c", ":")        // 程序输出的结果
//	want := []string{"a", "b", "c"}    // 期望的结果
//	if !reflect.DeepEqual(want, got) { // 因为slice不能比较直接，借助反射包中的方法比较
//		t.Errorf("expected:%v, got:%v", want, got) // 测试失败输出错误提示
//	}
//	t.Log(got, want)
//}

func TestSplitWithComplexSep1(t *testing.T) {
	got := Split1("abcd", "bc")
	want := []string{"a", "d"}
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}
