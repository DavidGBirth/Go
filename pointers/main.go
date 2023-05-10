package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var newVar int32 = 2
	fmt.Println("Print Main:", newVar, &newVar)

	printReference(&newVar)

	printAsCopy(newVar)

	fmt.Println("Print Main:", newVar, &newVar)
}

func printReference(value *int32) {
	if value != nil {
		*value = 8
		fmt.Println("Print reference:", *value, value)
		return
	}
	fmt.Println("No pointer")
}

func printAsCopy(value int32) {
	value = 4
	fmt.Println("Print as copy:", value, &value, "bytes:", unsafe.Sizeof(value))
}