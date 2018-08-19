package main

func main() {}

func spiralMatrixIII(R int, C int, r0 int, c0 int) [][]int {
	traverseCnt := 1
	res := [][]int{}
	res = append(res, []int{r0, c0})
	stepLen := 0
	matrixSize := R * C
	for {
		if traverseCnt == matrixSize {
			break
		}
		stepLen++
		// 右, c0++, r0 不变
		for i := 0; i < stepLen; i++ {
			c0++
			if r0 >= 0 && r0 < R && c0 >= 0 && c0 < C {
				res = append(res, []int{r0, c0})
				traverseCnt++
			}
		}
		// 下，r0++, c0 不变
		for i := 0; i < stepLen; i++ {
			r0++
			if r0 >= 0 && r0 < R && c0 >= 0 && c0 < C {
				res = append(res, []int{r0, c0})
				traverseCnt++
			}
		}
		stepLen++
		// 左，c0--, r0 不变
		for i := 0; i < stepLen; i++ {
			c0--
			if r0 >= 0 && r0 < R && c0 >= 0 && c0 < C {
				res = append(res, []int{r0, c0})
				traverseCnt++
			}
		}
		// 上，r0--, c0 不变
		for i := 0; i < stepLen; i++ {
			r0--
			if r0 >= 0 && r0 < R && c0 >= 0 && c0 < C {
				res = append(res, []int{r0, c0})
				traverseCnt++
			}
		}
	}
	return res
}
