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

func Test_sliceHelperImpl_RemoveRange(t *testing.T) {
	tests := []struct {
		name    string
		start   int
		end     int
		slice   interface{}
		want    interface{}
		wantErr bool
	}{
		{
			name:    "移除指定范围内的元素",
			start:   2,
			end:     4,
			slice:   []int{1, 2, 3, 4, 5},
			want:    []int{1, 5},
			wantErr: false,
		},
		{
			name:    "移除整个切片",
			start:   1,
			end:     5,
			slice:   []int{1, 2, 3, 4, 5},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "起始索引大于结束索引",
			start:   3,
			end:     2,
			slice:   []int{1, 2, 3, 4, 5},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "起始索引超出范围",
			start:   6,
			end:     7,
			slice:   []int{1, 2, 3, 4, 5},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "结束索引超出范围",
			start:   2,
			end:     6,
			slice:   []int{1, 2, 3, 4, 5},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "输入不是切片",
			start:   1,
			end:     2,
			slice:   "not_a_slice",
			want:    nil,
			wantErr: true,
		},
		{
			name:    "空切片",
			start:   1,
			end:     1,
			slice:   []int{},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sliceHelperImpl{}
			got, err := s.RemoveRange(tt.start, tt.end, tt.slice)
			if (err != nil) != tt.wantErr {
				t.Errorf("RemoveRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveRange() got = %v, want %v", got, tt.want)
			}
		})
	}
}

type args struct {
	slice   interface{}
	element interface{}
}

func Test_sliceHelperImpl_IndexOf(t *testing.T) {
	// 测试用例
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// 测试用例1：元素在切片中，期望返回索引值
		{
			name: "Element exists in slice",
			args: args{
				slice:   []int{1, 2, 3, 4, 5},
				element: 3,
			},
			want:    2,
			wantErr: false,
		},
		// 测试用例2：元素不在切片中，期望返回-1
		{
			name: "Element does not exist in slice",
			args: args{
				slice:   []int{1, 2, 3, 4, 5},
				element: 6,
			},
			want:    -1,
			wantErr: false,
		},
		// TODO: 添加更多测试用例
	}

	// 遍历测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 创建切片辅助对象
			s := &sliceHelperImpl{}

			// 调用 IndexOf 方法获取结果
			got, err := s.IndexOf(tt.args.slice, tt.args.element)

			// 判断是否返回错误，与期望值相符
			if (err != nil) != tt.wantErr {
				t.Errorf("IndexOf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// 判断返回的索引值是否与期望值相符
			if got != tt.want {
				t.Errorf("IndexOf() got = %v, want %v", got, tt.want)
			}
		})
	}
}
