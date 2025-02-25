package lib

import (
	"errors"
	"testing"
)

type testCalc struct {
	expression string
	answer     float64
	err        error
}

func TestCalculateExpression(t *testing.T) {
	testData := []testCalc{
		{
			expression: "2+3",
			answer:     5,
		},
		{
			expression: "1/2",
			answer:     0.5,
		},
		{
			expression: "10/2",
			answer:     5,
		},
		{
			expression: "10*5",
			answer:     50,
		},
		{
			expression: "10-20",
			answer:     -10,
		},
		{
			expression: "27+(17*3)/25-13*(25+(15*3+4))",
			answer:     -932.96,
		},
		{
			expression: "19/0",
			err:        ErrorDivisionByZero,
		},
		{
			expression: "-(-11-(1*20/2)-11/2*3)",
			answer:     37.5,
		},
		{
			expression: "(-1)(-1)",
			answer:     1,
		},
		{
			expression: "-1(2*2-1)",
			answer:     -3,
		},
		{
			expression: "*32",
			answer:     32,
		},
	}
	for _, test := range testData {
		answer, err := CalculateExpression(test.expression)
		if !errors.Is(test.err, err) {
			t.Errorf("error mismatch, expected %v, got %v", test.err, err)
		}
		if answer != test.answer {
			t.Errorf("calculateExpression(%s) => %f, want %f", test.expression, answer, test.answer)
		}
	}
}
