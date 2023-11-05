package main

import (
	"fmt"

	"github.com/alireza-frj4/app1/internal/adapters/core/arithmetic"
)

func main() {

	arithAdapter := arithmetic.NewAdapter()
	result, err := arithAdapter.Addition(1, 3)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(result)
}
