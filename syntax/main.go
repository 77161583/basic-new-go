package main

import "fmt"

func DeferTwo() (a int) {
	a = 0
	defer func() {
		fmt.Println("Inside defer:", a) // 这将打印 "Inside defer: 1"
		a = 1
	}()
	return a // 这将返回 0
}

func main() {

}
