package utils

import "fmt"

func DebugStruct(s interface{}) {
	fmt.Printf("%+v\n", s)
}
