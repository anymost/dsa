package main

import "fmt"

/**
计算的几个特点：输入，输出，正确性，确定性，可行性，有穷性
*/

/**
对于算法分析，复杂度计算常用的方式是：
迭代：级数求和

递归：递归跟踪，递推方程
递推方程需要递归基
*/


/**
复杂度:
    常数复杂度: O(1)
    对数复杂度: O(lgn)
    多项式复杂度: O(n^c)
    指数复杂度: O(2^n)
*/

/*
复杂级数：
    算术级数: 1 + 2 + 3 + 4 + ... + n = O(n^2)
    幂方级数: 1^2 + 2^2 + 3^2 + 4^2 + ... + n^2 = O(n^3)
    几何级数: 2^1 + 2^2 + 2^3 + 2^4 + ... + 2^n = O(2^n)
	收敛级数: 1/1/2 + 1/2/3 + 1/3/4 + 1/4/5 + 1/(n-1)/n = O(1)
	调和级数: 1   + 1/2 + 1/3 + 1/4 + ... + 1/n = O(logn)
    对数级数: log1+ log2 + log3 + log4 + ... + logn = O(nlogn)
*/

/**
推导算法的可行性
由 不变性+单调性 -> 正确性
*/

/**
封底估算:
 1 day = 10^5s
 300 year = 10^10s
 100 year = 3*10^9s
*/

/**
减而治之
分而治之
 */

func hailstone(n int) []int {
	array := make([]int, 0)
	for n > 1 {
		if n % 2 == 0 {
			n = n / 2
		} else {
			n = 3 * n + 1
		}
		array = append(array, n)
	}
	return array
}

func sum(array []int, n int) int {
	if n < 1 {
		return 0
	} else {
		return sum(array, n-1) + array[n-1]
	}
}

func main() {
	val := sum([]int{1, 2, 3, 4, 5}, 5)
	fmt.Println(val)
}


