package main

import (
	"fmt"
	"math"
	"strutil"

	"github.com/raaynaldo/LearnGoLang/src/03_packages/strutil"
)

func main() {
	fmt.Println(math.Floor(2.7))
	fmt.Println(math.Ceil(2.7))
	fmt.Println(math.Sqrt(16))
	fmt.Prinln(strutil.Reverse("Hello"))
}
