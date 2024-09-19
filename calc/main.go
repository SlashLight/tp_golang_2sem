package main

import (
	"fmt"
	"github.com/SlashLight/tp_golang_2sem/lib"
	"strings"
)

func main() {
	var expression string
	fmt.Scan(&expression)
	cleanExpression := strings.Replace(expression, " ", "", -1)
	ans, err := lib.CalculateExpression(cleanExpression)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%.5f\n", ans)
	return
}
