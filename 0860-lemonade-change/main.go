package main

func main() {

}

func lemonadeChange(bills []int) bool {
	var valMap = map[int]int{}
	for i := 0; i < len(bills); i++ {
		if bills[i] == 5 {
			valMap[5]++
		}
		if bills[i] == 10 {
			if valMap[5] == 0 {
				return false
			}
			valMap[5]--
			valMap[10]++
		}
		if bills[i] == 20 {
			if valMap[10] > 0 && valMap[5] > 0 {
				valMap[10]--
				valMap[5]--
				continue
			}
			if valMap[5] > 3 {
				valMap[5] = valMap[5] - 3
				continue
			}
			return false
		}
	}

	return true
}
