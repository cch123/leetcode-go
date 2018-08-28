package main

func main() {
}
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	// 在 x 轴上的投影没有重合，或者在 y 轴上的投影没有重合，即认为两个矩形不相交
	if rec1[0] >= rec2[2] || rec2[0] >= rec1[2] {
		return false
	}
	if rec1[1] >= rec2[3] || rec2[1] >= rec1[3] {
		return false
	}
	return true
}
