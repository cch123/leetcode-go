package main

func findComplement(num int) int {
	var mask = ^0
	for num&mask > 0 {
		mask <<= 1
	}
	mask = ^mask
	return mask ^ num
}
