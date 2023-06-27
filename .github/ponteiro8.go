package main

import (
	"fmt"
)

func modifyValue(ptr *int) {
	*ptr = 42
}
func main() {
	ptr := new(int)
	*ptr = 10
	fmt.Println("Antes da modificação:", *ptr)
	modifyValue(ptr)
	fmt.Println("Após a modificação:", *ptr)
}
