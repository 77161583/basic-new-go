// slice_test.go
package genericTool

import (
	"reflect"
	"testing"
)

func TestSliceHelper_Add(t *testing.T) {
	// 准备测试数据
	slice := []int{1, 2, 3}
	element := 4
	expected := []int{1, 2, 3, 4}

	// 调用 Add 方法
	helper := NewSliceHelper()
	result, err := helper.Add(slice, element)
	if err != nil {
		t.Errorf("Add method returned error: %v", err)
	}

	// 检查结果类型
	resultSlice, ok := result.([]int)
	if !ok {
		t.Errorf("Add method did not return a slice")
	}

	// 检查结果是否正确
	if !reflect.DeepEqual(resultSlice, expected) {
		t.Errorf("Add method returned incorrect result. Expected: %v, got: %v", expected, resultSlice)
	}

	// 输出结果
	t.Logf("原始切片: %v", slice)
	t.Logf("要添加的元素: %v", element)
	t.Logf("预期结果: %v", expected)
	t.Logf("实际结果: %v", resultSlice)

	// 检查结果是否正确
	if !reflect.DeepEqual(resultSlice, expected) {
		t.Errorf("Add方法返回了不正确的结果: %v, got: %v", expected, resultSlice)
	}
}
