package main

import "fmt"

type s struct {
	name string
}

func main() {
	name1 := "fadly"
	name2 := "munandar"

	bb := []s{{name1}, {name2}}
	cc := make([]s, 0)

	for _, v := range bb {
		v.name = "x"
		cc = append(cc, v)
	}
	bb = cc
	fmt.Println(bb)
}
