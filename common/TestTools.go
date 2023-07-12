package common

import(
	"fmt"
)

/****
Print slice support any data type
***/
func PrintSlice[T any](s []T){
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}