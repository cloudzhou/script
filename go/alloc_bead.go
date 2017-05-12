package main

import "fmt"

func main() {
	fmt.Println(alloc(7, 3, 7))
}

// m 是珠子个数， n 是盘子个数，max 是第一个盘子最大允许分配数
func alloc(m int, n int, max int) int {
	// 没有珠子了，一种方法
	if m == 0 {
		return 1
	}
	// 只有最后一个盘子
	if n == 1 {
		return 1
	}
	// 珠子少于盘子，多余的没有必要
	if m < n {
		n = m
	}
	// max 不可能超过珠子总数
	if max > m {
		max = m
	}
	// 算出最小值
	min := m / n
	if m%n != 0 {
		min = min + 1
	}
	// 递归累加
	count := 0
	for i := max; i >= min; i-- {
		count += alloc(m-i, n-1, i)
	}
	return count
}
