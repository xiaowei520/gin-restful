package main

import (
	"fmt"
	"time"
)

//快排
func Quick2Sort(values []int) {
	if len(values) <= 1 {
		return
	}
	mid, i := values[0], 1
	head, tail := 0, len(values)-1
	for head < tail {
		//	fmt.Println(values)
		if values[i] > mid {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	values[head] = mid
	Quick2Sort(values[:head])
	Quick2Sort(values[head+1:])
}

//快速排序的时间性能取决于快速排序递归的深度，可以用递归树来描述递归算法的执行情况
//T（n）≤2T（n/2） +n，T（1）=0
//T（n）≤2（2T（n/4）+n/2） +n=4T（n/4）+2n
//T（n）≤4（2T（n/8）+n/4） +2n=8T（n/8）+3n
//……
//T（n）≤nT（1）+（log2n）×n= O(nlogn)
//在最优的情况下，快速排序算法的时间复杂度为O(nlogn)。
//在最坏的情况下，待排序的序列为正序或者逆序，每次划分只得到一个比上一次划分少一个记录的子序列，注意另一个为空。如果递归树画出来，它就是一棵斜树。
//此时需要执行n‐1次递归调用，且第i次划分需要经过n‐i次关键字的比较才能找到第i个记录，也就是枢轴的位置，
//因此比较次数为 等差数列1...n-1 =>最终其时间复杂度为O(n2)。
//
//就空间复杂度来说，主要是递归造成的栈空间的使用，最好情况，递归树的深度为log2n，其空间复杂度也就为O(logn)，最坏情况，需要进行n‐1递归调用，其空间复杂度为O(n)，平均情况，空间复杂度也为O(logn)。
//由于关键字的比较和交换是跳跃进行的，因此，快速排序是一种不稳定的排序方法。


//优化方案  并行的快速排序  快速排序算法是采用分治技术来进行实现的,所以开启多个线程增加实现速度. 但是要考虑 达到指定大小再去goruntine

func main() {

	t1 := time.Now() // get current time

	//	var number = []int{9,8,7,6,5,4,3,2,1}
	var number = []int{1, 2, 3, 9, 8, 7, 6, 5, 4, 100, 101, 102, 103, 104, 105, 106, 107}
	for k := 0; k < 1000000; k++ {

		Quick2Sort(number) // 168.080817ms
	}

	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)
}
