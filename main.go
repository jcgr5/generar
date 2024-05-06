package main

import (
	"fmt"
	"generar/slices"
)

func main() {
	x := slices.GenSliceStr(3, 5)
	for i := range x {
		fmt.Println(x[i])
	}
}
