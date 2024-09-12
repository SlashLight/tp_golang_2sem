package main

import (
	"fmt"
	"github.com/SlashLight/tp_golang_2sem/lib"
)

func main() {
	var expression string
	fmt.Scan(&expression)
	ans := lib.CalculateExpression(expression)
	fmt.Println(ans)
	return
}
