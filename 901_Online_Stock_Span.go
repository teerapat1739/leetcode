package leetcode

/*
Runtime: 227 ms, faster than 50.00% of Go online submissions for Online Stock Span.
Memory Usage: 9.6 MB, less than 9.09% of Go online submissions for Online Stock Span.
*/

type StockSpanner struct {
	prices  []int
	results []int
	stack   []int
}

func Constructor901() StockSpanner {
	return StockSpanner{
		prices:  []int{},
		results: []int{},
		stack:   []int{},
	}
}

func (this *StockSpanner) Next(price int) int {
	p := len(this.prices)
	result := 1
	this.prices = append(this.prices, price)
	// stack saves the price index in price decreasing order
	for len(this.stack) > 0 && this.prices[this.stack[len(this.stack)-1]] <= price {
		result += this.results[this.stack[len(this.stack)-1]]
		this.stack = this.stack[:len(this.stack)-1]
	}

	this.stack = append(this.stack, p)
	this.results = append(this.results, result)
	return result
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
