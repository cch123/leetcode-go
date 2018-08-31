package main

import "fmt"

func main() {
	//fmt.Println(canVisitAllRooms([][]int{[]int{1}, []int{2}, []int{3}, []int{}}))
	//fmt.Println(canVisitAllRooms([][]int{[]int{1, 3, 2}, []int{2, 3}, []int{2, 1, 3, 1}, []int{}}))
	fmt.Println(canVisitAllRooms([][]int{[]int{1}, []int{}, []int{0, 3}, []int{1}}))
}

func canVisitAllRooms(rooms [][]int) bool {
	//for i := 0; i < len(rooms); i++ {
	// start from rooms[i]
	visited := map[int]bool{0: true}
	nextVisit := rooms[0]
	for {
		newToVisit := []int{}
		for j := 0; j < len(nextVisit); j++ {
			if visited[nextVisit[j]] == true {
				continue
			}
			visited[nextVisit[j]] = true
			newToVisit = append(newToVisit, rooms[nextVisit[j]]...)
		}
		if len(newToVisit) == 0 {
			break
		}
		nextVisit = newToVisit
		//fmt.Println(nextVisit)
	}

	// check all
	if len(visited) == len(rooms) {
		//fmt.Println(visited, i)
		return true
	}
	//}
	return false
}
