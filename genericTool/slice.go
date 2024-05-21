package genericTool

import (
	"errors"
	"reflect"
)

// SliceHelper 声明一个通用的操作接口
type SliceHelper interface {
	Add(slice interface{}, element interface{}) (interface{}, error)
	Remove(slice interface{}, index int) (interface{}, error)
	RemoveRange(start int, end int, slice interface{}) (interface{}, error)
	IndexOf(slice interface{}, element interface{}) (int, error)
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

// Remove 单个下标删除
func (s *sliceHelperImpl) Remove(slice interface{}, index int) (interface{}, error) {
	// 使用类型断言将 slice 转换为具体的切片类型
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return nil, errors.New("输入的不是一个切片")
	}

	// 检查索引是否越界
	if index < 0 || index >= sliceValue.Len() {
		return nil, errors.New("越界了")
	}

	// 创建一个新的切片，长度为原始切片长度减1
	newSlice := reflect.MakeSlice(sliceValue.Type(), sliceValue.Len()-1, sliceValue.Len()-1)

	// 将原始切片中索引之前的元素复制到新切片中
	reflect.Copy(newSlice, sliceValue.Slice(0, index))

	// 将原始切片中索引之后的元素复制到新切片中
	reflect.Copy(newSlice.Slice(index, newSlice.Len()), sliceValue.Slice(index+1, sliceValue.Len()))

	// 返回新的切片
	return newSlice.Interface(), nil
}

// RemoveRange 删除指定范围的元素
func (s *sliceHelperImpl) RemoveRange(start int, end int, slice interface{}) (interface{}, error) {
	// 使用类型断言将 slice 转换为具体的切片类型
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Kind() != reflect.Slice {
		return nil, errors.New("input is not a slice")
	}

	sliceLen := sliceValue.Len()
	if sliceLen == 0 {
		return nil, errors.New("长度不能为空")
	}

	// 如果超出范围就返回错误
	if start < 1 || start > sliceLen || end < start || end > sliceLen {
		return nil, errors.New("超出范围")
	}

	// 重置初始化下标
	newStart := start - 1

	// 计算要删除的下标
	switch {
	// 整个切片就直接返回空
	case newStart == 0 && end == sliceLen:
		return nil, nil
	case newStart == 0:
		return sliceValue.Slice(end, sliceLen).Interface(), nil
	case end == sliceLen:
		return sliceValue.Slice(newStart, sliceLen).Interface(), nil
	default:
		// 直接在原始切片上进行操作
		newSlice := reflect.MakeSlice(sliceValue.Type(), 0, 0)
		for i := 0; i < newStart; i++ {
			newSlice = reflect.Append(newSlice, sliceValue.Index(i))
		}
		for i := end; i < sliceLen; i++ {
			newSlice = reflect.Append(newSlice, sliceValue.Index(i))
		}
		return newSlice.Interface(), nil
	}
}

// IndexOf 方法用于查找元素在切片中的索引值
func (s *sliceHelperImpl) IndexOf(slice interface{}, element interface{}) (int, error) {
	// 使用反射获取传入切片的值
	sliceValue := reflect.ValueOf(slice)

	// 检查传入的值是否是切片类型
	if sliceValue.Kind() != reflect.Slice {
		return -1, errors.New("输入的不是一个切片")
	}

	// 将要查找的元素转换为 reflect.Value 类型的值
	elementValue := reflect.ValueOf(element)

	// 遍历切片中的每个元素
	for i := 0; i < sliceValue.Len(); i++ {
		// 使用反射比较切片中的当前元素与要查找的元素是否相等
		if reflect.DeepEqual(sliceValue.Index(i).Interface(), elementValue.Interface()) {
			// 如果相等，则返回当前元素的索引值
			return i, nil
		}
	}

	// 如果未找到相等的元素，则返回 -1 表示未找到
	return -1, nil
}
