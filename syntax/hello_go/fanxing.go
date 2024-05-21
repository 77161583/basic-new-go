package main

type fanxing[T any] interface {
	Add(inx int, t T)
	Append(t T)
}

func UserList() {
	var i fanxing[int]
	i.Add(1, 33)
	i.Append(12)
}

type LinkList[T any] struct {
	head *node[T]
}

type node[T any] struct {
	val T
}
