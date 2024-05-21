package main

import "fmt"

func NewsUser() {
	//初始化结构体
	u := User{}
	fmt.Printf("%+v \n", u)

	//up
	up := &User{}
	fmt.Printf("up %+v \n", up)
	up2 := new(User)
	println(up2.Name)
	fmt.Printf("up2 %+v ", up2)

}

type User struct {
	Name string
	Age  int
}
