package main

func Sum[T Number](vals ...T) T {
	var res T
	for _, val := range vals {
		res = res + val
	}
	return res

}

type Number interface {
	//加上 ~ 表示 int 以及int的衍生类型
	~int | float64
}

type Integer int
