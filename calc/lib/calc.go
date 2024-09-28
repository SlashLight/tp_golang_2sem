package lib

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
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
	s = clean(s)
	polishNotation := getReversePolishNotation(s)
	nums := make([]float64, 0, 10)
	var (
		num1, num2 float64
	)
	if len(polishNotation) <= 2 {
		return strconv.ParseFloat(polishNotation[0], 64)
	}
	for _, val := range polishNotation {
		if val == "+" || val == "-" || val == "*" || val == "/" || val == "~" {
			if val == "~" {
				nums[len(nums)-1] = 0 - nums[len(nums)-1]
				continue
			}
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
	/*if s[*pos] == '-' {
		number += string(s[*pos])
		*pos++
	}*/
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

func clean(expression string) string {
	addMult := regexp.MustCompile("(\\d+)(\\()")
	cleanExpression := strings.Replace(expression, " ", "", -1)
	cleanExpression = strings.Replace(cleanExpression, ")(", ")*(", -1)
	cleanExpression = addMult.ReplaceAllString(cleanExpression, "${1}*$2")
	return cleanExpression
}

func getReversePolishNotation(s string) []string {
	priorityMap := map[byte]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
		'~': 3,
	}
	var polishNotation []string
	st := new(Stack)

	for i := 0; i < len(s); i++ {
		ch := s[i]
		_, err := strconv.Atoi(string(ch))
		switch {
		case err == nil:
			polishNotation = append(polishNotation, getNumberFromString(s, &i))
			/*if ch == '-' {
				for !st.isEmpty() && priorityMap[st.Top()] >= priorityMap[ch] {
					polishNotation = append(polishNotation, string(st.Pop()))
				}
				st.Push('+')
			}*/
		case ch == '(':
			st.Push(ch)
		case ch == ')':
			for !st.isEmpty() && st.Top() != '(' {
				polishNotation = append(polishNotation, string(st.Pop()))
			}
			st.Pop()
		default:
			if ch == '-' && (i == 0 || s[i-1] == '(') {
				ch = '~'
			}
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
