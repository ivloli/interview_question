package main

import (
	"core/rangeList"
	"fmt"
)

func main() {
	r := rangeList.RangeList{}
	r.AddRange(1, 5)
	r.AddRange(7, 10)
	r.AddRange(15, 20)
	fmt.Println(r.ToString())
}
