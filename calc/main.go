package main

import (
	"fmt"
	"github.com/SlashLight/tp_golang_2sem/lib"
)

func main() {
	var expression string
	fmt.Scan(&expression)

	ans, err := lib.CalculateExpression(expression)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.5f\n", ans)
	return
}
