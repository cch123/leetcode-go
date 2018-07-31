package main

func judgeCircle(moves string) bool {
    var stepMap = map[byte]int{}
    for i:=0;i<len(moves);i++{
        stepMap[moves[i]]++
    }

    if stepMap['U'] == stepMap['D'] && stepMap['L'] == stepMap['R'] {
        return true
    }
    return false
}
