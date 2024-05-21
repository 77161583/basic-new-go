package main

type Inner struct {
}

type Outer struct {
	Inner
}

func (o Outer) Name() string {
	return "Outer"
}

func (i Inner) SayHello() {
	println("hello", i.Name())
}

func (i Inner) Name() string {
	return "Inner"
}

func UserOuter() {
	var o Outer
	o.SayHello()
}

func main() {
	//UserOuter()
	println(Sum[int](2, 2, 2))
	println(Sum[Integer](2, 3, 2))

}
