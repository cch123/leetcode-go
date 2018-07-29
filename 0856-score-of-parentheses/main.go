package main

type stack struct {
	arr []int
}

func (s *stack) push(elem int) {
	s.arr = append(s.arr, elem)
}

func (s *stack) pop() {
	s.arr = s.arr[:len(s.arr)-1]
}

func (s *stack) top() int {
	return s.arr[len(s.arr)-1]
}

func (s *stack) findTopMostLeftParen() int {
	for i := len(s.arr) - 1; i >= 0; i-- {
		if s.arr[i] == int('(') {
			return i
		}
	}
	return -1
}

func scoreOfParentheses(S string) int {
	st := stack{arr: make([]int, 0)}
	for _, b := range S {
		if b == '(' {
			st.push(int(b))
		}

		if b == ')' {
			if st.top() == int('(') {
				st.pop()
				st.push(1)
			} else {
				// 一直找到匹配的括号
				idx := st.findTopMostLeftParen()
				sum := 0
				for i := idx + 1; i < len(st.arr); i++ {
					sum += st.arr[i]
				}
				sum *= 2

				for st.top() != int('(') {
					st.pop()
				}
				// pop the '('
				st.pop()
				st.push(sum)
			}
		}
	}

	res := 0
	for _, n := range st.arr {
		res += n
	}

	return res
}
