package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(3))
}

var res = []string{}

func goDeeper(curVal int, currentStr string, leftLeft int, rightLeft int) {
	if curVal > 0 {
		// 放 (
		if leftLeft > 0 {
			goDeeper(curVal+1, currentStr+string('('), leftLeft-1, rightLeft)
		}
		// 放 )
		if rightLeft > 0 {
			goDeeper(curVal-1, currentStr+string(')'), leftLeft, rightLeft-1)
		}
	}

	if curVal == 0 {
		if leftLeft == 0 && rightLeft == 0 {
			//res++
			res = append(res, currentStr)
			return
		}

		// 只能放 (
		if leftLeft > 0 {
			goDeeper(curVal+1, currentStr+string('('), leftLeft-1, rightLeft)
		}
	}

	if curVal < 0 {
		return
	}
}

func generateParenthesis(n int) []string {
	res = res[:0]
	goDeeper(0, "", n, n)
	return res
}
