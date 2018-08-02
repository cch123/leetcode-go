package main

func distributeCandies(candies []int) int {
    recMap := map[int]bool{}
    typeCnt := 0
    for i:=0;i<len(candies);i++{
        if recMap[candies[i]] == false {
            recMap[candies[i]] = true
            typeCnt++
        }
    }
    
    if typeCnt > len(candies)/2 {
        return len(candies)/2
    }
    
    return typeCnt
}
