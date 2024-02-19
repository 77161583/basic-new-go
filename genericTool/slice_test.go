// slice_test.go
package genericTool

import (
	"fmt"
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

func TestSliceHelperImpl_Remove(t *testing.T) {
	//slice := []int{1, 2, 3, 4, 5}
	slice := []string{"a", "b", "c", "d"}
	index := 0

	helper := NewSliceHelper()
	result, err := helper.Remove(slice, index)
	if err != nil {
		t.Errorf("Remove method returned error: %v", err)
	}

	expected := []string{"b", "c", "d"}
	fmt.Printf("Type of result: %T\n", result)

	resultSlice, ok := result.([]string)
	if !ok {
		t.Errorf("Remove method did not return a slice")
	}

	fmt.Printf("期待结果: %v\n", expected)
	fmt.Printf("输出结果: %v\n", resultSlice)

	if !reflect.DeepEqual(resultSlice, expected) {
		t.Errorf("Remove method returned incorrect result. Expected: %v, got: %v", expected, resultSlice)
	}

}
