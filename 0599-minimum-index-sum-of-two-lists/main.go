package main

import "fmt"

func findRestaurant(list1 []string, list2 []string) []string {
	l1Map := map[string]int{}
	l2Map := map[string]int{}
	var res []string
	for i := 0; i < len(list1); i++ {
		l1Map[list1[i]] = i
	}
	for i := 0; i < len(list2); i++ {
		l2Map[list2[i]] = i
	}

	var max = len(list1) + len(list2)
	for i := 0; i < len(list1); i++ {
		l1Str := list1[i]
		if _, ok := l2Map[l1Str]; ok {
			sum := l2Map[l1Str] + i
			if sum < max {
				max = sum
				res = res[:0]
				res = append(res, list1[i])
				continue
			}

			if sum == max {
				res = append(res, list1[i])
				continue
			}
		}
	}

	return res

}

func main() {
	fmt.Println(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}))
}
