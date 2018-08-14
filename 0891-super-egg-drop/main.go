/*
 f[step][k]
 step 表示步数
 k 表示鸡蛋数
 这道题的关键点在于，不是用两个输入维度来做动态规划
 而是用一个输入维度和一个输出维度来做，
 其结果 >= 其中的另一个输入维度时，即满足输出条件，
 但因为可能存在多条结果，所以要把所有结果都走完
 只要用来 DP 的两个维度能有确定的上界就可以
*/
package main

import "fmt"

func main() {
	fmt.Println(superEggDrop(1, 2))
	fmt.Println(superEggDrop(2, 6))
	fmt.Println(superEggDrop(2, 4))
}

func superEggDrop(K int, N int) int {
	var f = [][]int{}
	emptyLine := make([]int, K+1)
	f = append(f, emptyLine)
	result := N + 1
	for steps := 1; steps <= N; steps++ {
		line := make([]int, K+1)
		f = append(f, line)
		for k := 1; k <= K; k++ {
			if steps == 1 {
				f[steps][k] = 1
			} else if k == 1 {
				f[steps][k] = steps
			} else {
				// steps > 1
				// k > 1
				f[steps][k] = 1 + f[steps-1][k-1] + f[steps-1][k]
			}
			//fmt.Println(f)

			if f[steps][k] >= N && steps < result {
				//fmt.Println(steps, k, f[steps])
				result = steps
			}
		}
	}
	return result
}
