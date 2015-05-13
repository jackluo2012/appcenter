package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%qn", strings.Split("a,b,c,", ","))
}
