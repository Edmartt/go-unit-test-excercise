package main

import (
	"fmt"
	"testproject/simple/ops"
)

func main(){
	result := ops.Sum(10, 10)

	fmt.Printf("The result is %d \n", result)
}
