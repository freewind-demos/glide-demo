package main

import (
	"github.com/golang-demos/go-concat-strings-library/strlib"
	"fmt"
)

func main() {
	hello := strlib.ConcatStrs("Hello, Glide")
	fmt.Println(hello)
}
