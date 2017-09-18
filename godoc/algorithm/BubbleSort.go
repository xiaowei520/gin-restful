package main

import (
	"fmt"
	"time"
)

//算法原理:  正序排序
//比较相邻的元素,如果第一个比第二个大,就交换他们两个.
//算法分析:
//	时间复杂度:
//		最好时间复杂度O(n)  本身是正序 ,比较次数 n-1 移动次数 0
//
//		最坏时间复杂度O(n*n) 本身是逆序  比较次数 n-1,n-2,.... 1 等差数列  Sn=[n*(a1+an)]/2。 =>  (n-1+1)*n /2
//		移动次数:--3倍 =>   ((n-1+1)*n /2) * 3   tmp=a;a=b;b=tmp
//算法稳定性排序

//冒泡排序 -基本实现
func BubbleSort(values [] int) ([] int) {
	lens := len(values)
	var tmp int
	flag := true
	for i := 0; i < lens; i++ {
		for j := lens - 1; j > i; j-- {
			if values[j-1] > values[j] {
				tmp = values[j-1]
				values[j-1] = values[j]
				values[j] = tmp

				flag = false
				continue
			}
		}
		//一遍排序下来没有变为false 证明正序排序 退出
		if flag {
			break
		}
	}
	return values
}

//冒泡排序-优化
func BubbleSortOpt(values [] int) ([] int) {
	lens := len(values)
	var tmp int

	var lastSwapPos int = 0
	var lastSwapPosTemp int = 0

	for i := 0; i < lens; i++ {

		lastSwapPos = lastSwapPosTemp;
		for j := lens - 1; j > lastSwapPos; j-- {
			if values[j-1] > values[j] {
				tmp = values[j-1]
				values[j-1] = values[j]
				values[j] = tmp

				lastSwapPosTemp = j;
			}
		}
		if lastSwapPos == lastSwapPosTemp {
			break;
		}

	}
	return values
}
//写了好多冒泡-本质上效率并没有太大差距-
func main() {
	t1 := time.Now() // get current time

	//	var number = []int{9,8,7,6,5,4,3,2,1}
	var number = []int{1, 2, 3, 9, 8, 7, 6, 5, 4,100,101,102,103,104,105,106,107}
	var result = []int{}
	for k := 0; k < 1000000; k++ {

			result = BubbleSort(number)// 13.500457ms
	//	result = BubbleSortOpt(number) //  17.319948ms
	}

	for i, value := range result {
		fmt.Printf("a[%d]=%d\n", i, value)
	}

	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)
}
