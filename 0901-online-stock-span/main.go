package main

type StockSpanner struct {
	elemList []int
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	if len(this.elemList) == 0 {
		this.elemList = append(this.elemList, price)
		return 1
	}
	var res = 1
	for i := len(this.elemList) - 1; i >= 0; i-- {
		if this.elemList[i] <= price {
			res++
		} else {
			break
		}
	}
	this.elemList = append(this.elemList, price)
	return res
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */

func main() {
	var s = Constructor()
	println(s.Next(100))
	println(s.Next(80))
	println(s.Next(60))
	println(s.Next(70))
	println(s.Next(60))
	println(s.Next(75))
	println(s.Next(85))
}
