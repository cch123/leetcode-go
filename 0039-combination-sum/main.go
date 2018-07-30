package main

func combinationSum(candidates []int, target int) [][]int {
	fmt.Println(candidates, target)
	if target == 0 {
		return [][]int{[]int{}}
	}

	if len(candidates) == 0 {
		// 无解
		return nil
	}

	if len(candidates) == 1 {
		if target%candidates[0] == 0 {
			result := []int{}
			for i := 0; i < (target / candidates[0]); i++ {
				result = append(result, candidates[0])
			}
			return [][]int{result}
		}

		// 无解
		return nil
	}

	result := [][]int{}
	for i := 0; i <= target/candidates[0]; i++ {
		subResult := combinationSum(candidates[1:], target-i*candidates[0])
		if subResult != nil {
			// [][]int{}
			// 里的每个数组都追加上当前的 candidate 0 即可
			for j := 0; j < len(subResult); j++ {
				for k := 0; k < i; k++ {
					subResult[j] = append(subResult[j], candidates[0])
				}
				result = append(result, subResult[j])
			}
		}
	}
	return result
}

