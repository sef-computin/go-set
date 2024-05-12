package main

import (
	"fmt"
	set "github.com/sef-computin/go-set"
)

func main() {
	s := set.NewSet()

	ok := s.Add(23)
	fmt.Println(ok)
	ok = s.Add(-52)
	fmt.Println(ok)
	ok = s.Add(23)
	fmt.Println(ok)
	ok = s.Add("string")
	fmt.Println(ok)

	fmt.Println(s.ToSlice()...)
}
