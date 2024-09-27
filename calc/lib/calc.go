package lib

import (
	"fmt"
	"strconv"
)

type Stack []byte

var ErrorDivisionByZero = fmt.Errorf("Division by zero")

func (s *Stack) Push(v byte) {
	*s = append(*s, v)
}

func (s *Stack) Top() byte {
	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() byte {
	v := s.Top()
	*s = (*s)[:len(*s)-1]
	return v
}

func (s *Stack) isEmpty() bool {
	return len(*s) == 0
}

func CalculateExpression(s string) (float64, error) {
	polishNotation := getReversePolishNotation(s)
	nums := make([]float64, 0, 10)
	var (
		num1, num2 float64
	)
	for _, val := range polishNotation {
		if val == "+" || val == "-" || val == "*" || val == "/" {
			num2 = nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			if len(nums) < 1 {
				num1 = 0
			} else {
				num1 = nums[len(nums)-1]
				nums = nums[:len(nums)-1]
			}

			switch val {
			case "+":
				num1 += num2
			case "-":
				num1 -= num2
			case "*":
				num1 *= num2
			case "/":
				if num2 == 0 {
					return 0, ErrorDivisionByZero
				}
				num1 /= num2
			}

			nums = append(nums, num1)
		} else {
			num, err := strconv.ParseFloat(val, 64)
			if err != nil {
				return 0, err
			}
			nums = append(nums, num)
		}
	}

	return nums[0], nil
}

func getNumberFromString(s string, pos *int) string {
	var number string
	for ; *pos < len(s); *pos++ {
		_, err := strconv.Atoi(string(s[*pos]))
		if err == nil {
			number += string(s[*pos])
		} else {
			*pos--
			break
		}
	}

	return number
}

func getReversePolishNotation(s string) []string {
	priorityMap := map[byte]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}
	var polishNotation []string
	st := new(Stack)

	for i := 0; i < len(s); i++ {
		ch := s[i]
		_, err := strconv.Atoi(string(ch))
		switch {
		case err == nil:
			polishNotation = append(polishNotation, getNumberFromString(s, &i))
		case ch == '(':
			st.Push(ch)
		case ch == ')':
			for !st.isEmpty() && st.Top() != '(' {
				polishNotation = append(polishNotation, string(st.Pop()))
			}
			st.Pop()
		default:
			for !st.isEmpty() && priorityMap[st.Top()] >= priorityMap[ch] {
				polishNotation = append(polishNotation, string(st.Pop()))
			}
			st.Push(ch)
		}
	}

	for !st.isEmpty() {
		polishNotation = append(polishNotation, string(st.Pop()))
	}
	return polishNotation
}
