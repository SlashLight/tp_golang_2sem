package lib

import "strconv"

type Stack []byte

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

func CalculateExpression(s string) int {
	polishNotation := getReversePolishNotation(s)
	nums := make([]int, 0, 10)

	for _, val := range polishNotation {
		if val == "+" || val == "-" || val == "*" || val == "/" {
			num1 := nums[len(nums)-2]
			num2 := nums[len(nums)-1]
			nums = nums[:len(nums)-2]

			switch val {
			case "+":
				num1 += num2
			case "-":
				num1 -= num2
			case "*":
				num1 *= num2
			case "/":
				num1 /= num2
			}

			nums = append(nums, num1)
		} else {
			num, _ := strconv.Atoi(val)
			nums = append(nums, num)
		}
	}

	return nums[0]
}

func isDigit(r byte) bool {
	if '0' <= r && r <= '9' {
		return true
	}
	return false
}

func getNumberFromString(s string, pos *int) string {
	var number string
	for ; *pos < len(s); *pos++ {
		if isDigit(s[*pos]) {
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

		switch {
		case isDigit(ch):
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
