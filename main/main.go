package main

import (
	"fmt"
	"sort"
)

func main() {
	//bs := []byte("hello world")
	bs := "string"
	sort.Slice([]byte(bs), func(i, j int) bool {
		return bs[i] < bs[j]
	})
	fmt.Printf("%#v\n", bs)
	println(string(bs))
}
