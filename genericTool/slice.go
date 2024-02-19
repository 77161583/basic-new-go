package genericTool

import (
	"errors"
	"reflect"
)

// SliceHelper 声明一个通用的操作接口
type SliceHelper interface {
	Add(slice interface{}, element interface{}) (interface{}, error)
}

// 切片辅助方法的具体实现
type sliceHelperImpl struct{}

// NewSliceHelper 创建并返回 SliceHelper 接口的实例
func NewSliceHelper() SliceHelper {
	return &sliceHelperImpl{}
}

func (s *sliceHelperImpl) Add(slice interface{}, element interface{}) (interface{}, error) {
	// 使用类型断言将 slice 转换为具体的切片类型
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return nil, errors.New("Input is not a slice")
	}

	// 创建一个新的切片，长度为原始切片长度加1
	newSlice := reflect.MakeSlice(sliceValue.Type(), sliceValue.Len()+1, sliceValue.Len()+1)

	// 将原始切片的元素复制到新切片中
	reflect.Copy(newSlice, sliceValue)

	// 将新元素添加到新切片的末尾
	newElement := reflect.ValueOf(element)
	newSlice.Index(sliceValue.Len()).Set(newElement)

	// 返回新的切片
	return newSlice.Interface(), nil
}

func (s *sliceHelperImpl) Remove(slice interface{}, index int) (interface{}, error) {
	// 使用类型断言将 slice 转换为具体的切片类型
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return nil, errors.New("输入的不是一个切片")
	}

	// 检查索引是否越界
	if index < 0 || index > sliceValue.Len() {
		return nil, errors.New("越界了")
	}

	// 创建一个新的切片，长度为原始切片长度减1
	newSlice := reflect.MakeSlice(sliceValue.Type(), sliceValue.Len()-1, sliceValue.Len()-1)

	// 将原始切片中索引之前的元素复制到新切片中
	reflect.Copy(newSlice, sliceValue.Slice(0, index))

	// 将原始切片中索引之后的元素复制到新切片中
	reflect.Copy(newSlice.Slice(index, newSlice.Len()), sliceValue.Slice(index+1, sliceValue.Len()-1))
}
