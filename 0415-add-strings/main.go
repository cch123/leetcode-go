package main

func reverse(str string) string {
	strBytes := []byte(str)
	for i := 0; i < len(strBytes)/2; i++ {
		strBytes[i], strBytes[len(strBytes)-1-i] = strBytes[len(strBytes)-1-i], strBytes[i]
	}
	return string(strBytes)
}

var byteArr = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}
var byteMap = map[byte]int{'0': 0, '1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}

func addStrings(num1 string, num2 string) string {
	num1 = reverse(num1)
	num2 = reverse(num2)
	if len(num2) < len(num1) {
		num1, num2 = num2, num1
	}
	resInts := []int{}
	// num1 比 num2 短
	for i := 0; i < len(num1); i++ {
		plusRes := byteMap[num1[i]] + byteMap[num2[i]]
		resInts = append(resInts, plusRes)
	}
	for i := len(num1); i < len(num2); i++ {
		resInts = append(resInts, byteMap[num2[i]])
	}

	// 进位
	for i := 0; i < len(resInts)-1; i++ {
		if resInts[i] >= 10 && i+1 < len(resInts) {
			resInts[i+1]++
			resInts[i] -= 10
			continue
		}
	}

	// 最后一位特殊处理
	if resInts[len(resInts)-1] >= 10 {
		resInts[len(resInts)-1] -= 10
		resInts = append(resInts, 1)
	}

	resBytes := []byte{}
	// gen bytes
	for i := 0; i < len(resInts); i++ {
		resBytes = append(resBytes, byteArr[resInts[i]])
	}
	return reverse(string(resBytes))
}

func main() {
	println(addStrings("1", "9"))
}
