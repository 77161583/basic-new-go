package main

import (
	"basic-new-go/genericTool"
	"fmt"
)

func main() {
	// Your main function code here
	slice := []string{"3", "asd", "qq"}
	element := "11"

	helper := genericTool.NewSliceHelper()
	res, err := helper.Add(slice, element)
	if err != nil {
		fmt.Println("报错了")
		return
	}

	fmt.Println("Result:", res)
}
