package main

import "fmt"

func gcd(A, B int) int {
	if A > B {
		A, B = B, A
	}

	if B%A == 0 {
		return A
	}

	return gcd(B-A, A)
}

func nthMagicalNumber(N int, A int, B int) int {
	// ensure a <= b
	if A > B {
		A, B = B, A
	}

	if A == B || B%A == 0 {
		return int(int64(N*A) % 1000000007)
	}

	// two array
	// first: A, A*2, A*3, A*4
	// secon: B, B*2, ... --B * A/ gcd(B, A)-del-
	zxgbs := int64(A) * int64(B) / int64(gcd(A, B))
	// startPosVal := A * N
	i := int64(N)
	j := int64(A*N/B) - int64(A*N)/(zxgbs)
	high := N
	low := 0
	for low <= high {
		mid := low + (high-low)/2
		i = int64(mid)
		j = int64(A)*int64(i)/int64(B) - int64(A)*int64(i)/(zxgbs)
		if i+j > int64(N) {
			high = mid - 1
			continue
		}
		if i+j < int64(N) {
			low = mid + 1
			continue
		}
		if i+j == int64(N) {
			return int(int64(A) * i % int64(1000000007))
		}
	}

	j = int64(N)
	i = j * int64(B) / int64(A)
	high = int(j)
	low = 0
	for low <= high {
		mid := low + (high-low)/2
		j = int64(mid)
		i = int64(B) * int64(j) / int64(A)
		beforeJ := j * int64(B) / zxgbs
		if i+j-beforeJ > int64(N) {
			high = mid - 1
			continue
		}
		if i+j-beforeJ < int64(N) {
			low = mid + 1
			continue
		}
		if i+j-beforeJ == int64(N) {
			if j*int64(B) == zxgbs {
				return int(j + int64(1)%int64(1000000007))
			}
			return int(j * int64(B) % int64(1000000007))
		}
	}
	return 0
}

func main() {
	fmt.Println(nthMagicalNumber(1, 2, 3))
	fmt.Println(nthMagicalNumber(4, 2, 3))
	fmt.Println(nthMagicalNumber(5, 2, 4))
	fmt.Println(nthMagicalNumber(3, 6, 4))
	fmt.Println(nthMagicalNumber(10, 10, 8))
}
