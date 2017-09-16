package main

import (
	"sync"
	"os"
	"fmt"
	"time"
)

/*
	声明 sync 包下面的 WaitGroup
	用途：它能够一直等到所有的go routine执行完成，并且阻塞主线程的执行，直到所有的go routine执行完成。
	WaitGroup总共有三个方法：Add(delta int),Done(),Wait()。简单的说一下这三个方法的作用。
	Add:添加或者减少等待go routine的数量
    Done:相当于Add(-1)
    Wait:执行阻塞，直到所有的WaitGroup数量变成0
*/
var wg sync.WaitGroup

//声明 文件数组
var fileArr [4] *os.File

/*
	make用于内建类型（map、slice 和channel）的内存分配。new用于各种类型的内存分配。
	make(T, args)与new(T)有着不同的功能，make只能创建slice、map和channel，并且返回一个有初始值(非零)的T类型（引用），而不是*T。
	/创建一个初始元素个数为5的数组切片，元素初始值为0，并预留10个元素的存储空间
	b := make([]int, 5, 10) // len(b)=5, cap(b)=10
	//继续切片，注意len和cap的变化
	b = b[:cap(b)] // len(b)=5, cap(b)=5
	b = b[1:]      // len(b)=4, cap(b)=4
 */

/*
	信道是什么？简单说，是goroutine之间互相通讯的东西。类似我们Unix上的管道（可以在进程间传递消息）， 用来goroutine之间发消息和接收消息。
	其实，就是在做goroutine之间的内存共享。
 */
var chArr [4] chan int

/*
	chan 通道 - 操作不好会造成死锁 | 非缓冲信道上如果发生了流入无流出，或者流出无流入，也就导致了死锁。 |
	比如 1 . ch <- 1 // 1流入信道，堵塞当前线, 没人取走数据信道不会打开  单一goroutine Main 函数中做
	2.   ch1 <- <- ch2 // ch1 等待 ch2流出的数据

	无缓冲信道 make(chan int) (默认为0，即无缓冲)

	缓冲信道 make(chan int,4)  不必阻塞当前线而等待该数据取走。
 */
//缓存信道-缓冲信道是先进先出的，我们可以把缓冲信道看作为一个线程安全的队列
func buffered_chan(){
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	//ch <- 4   如果当流入第四个的时候 会导致死锁

	//range不等到信道关闭是不会结束读取的,如果缓冲通道干涸 range就会阻塞当前goroutine, 所以死锁咯。所以加个判断break | 显式地关闭信道
	//close(ch) 被关闭的信道会禁止数据流入, 是只读的。我们仍然可以从关闭的信道中取出数据，但是不能再写入数据了。

	for v := range ch {
		fmt.Println(v)
		if len(ch) <= 0 { // 如果现有数据量为0，跳出循环
			break
		}
	}
}

//测试信道 等待多goroutine的方案 |开出很多小goroutine,finshed to master。

var quit chan int // 只开一个信道
func foo(id int) {
	fmt.Println(id)
	quit <- 0 // ok, finished
}
func moreGoRun(){
	count := 1000
	quit = make(chan int) // 无缓冲  无缓冲的信道是一批数据一个一个的「流进流出」
//	quit = make(chan int,1000) // 有缓冲  缓冲信道则是一个一个存储，然后一起流出去

	for i := 0; i < count; i++ {
		go foo(i)
	}

	for i := 0; i < count; i++ {
		<- quit
	}
}


func yars_test1() {
	wg.Add(4)
	//定义四个通道


	//chArr := make([]chan int,4)
	for i := 0; i < 4; i++ {
		tmpChan := make(chan int)
		chArr [i] = tmpChan
	}
	fmt.Printf("111")
	//循环创建四个文件
	for i := 0; i < 4; i++ {
		fileName := fmt.Sprintf("/tmp/%d_test.hehe", i)
		fmt.Printf(fileName)
		name, err := os.Create(fileName)
		fileArr[i] = name
		if err != nil {
			fmt.Println(err.Error())
			fmt.Println(err.Error())
		}
		go myfun1(chArr[i], fileArr[i]) // 启动一个goroutine
		chArr[i] <- i
	}
	wg.Wait()
	for i := 0; i < 4; i++ {
		fileArr[i].Close()
	}
}
//递归调度 写入%4的倍数 1 2 3 4 1 2 3 4 | 2 3 4 1  | 3 4 1 2 | 4 1 2 3
func myfun1(ch chan int, f *os.File) {

	num := <-ch
	//从1开始
	t := (num % 4) + 1
	f.Write([]byte(fmt.Sprint(t)))
	num += 1
	if num == 100 {
		wg.Done()
		close(ch)
		return
	}

	go myfun1(ch, f)
	ch <- num
}

/**
程序很简单，只是将每次循环的数量过3秒钟输出。
那么，这个程序如果不用WaitGroup，那么将看不见输出结果。因为go routine还没执行完，主线程已经执行完毕。
注释的defer wg.Done()和defer wg.Add(-1)作用一样。这个很好，原来执行脚本，都是使用time.Sleep，用一个估计的时间等到子线程执行完。WaitGroup很好。虽然chanel也能实现，但是觉得如果涉及不到子线程与主线程数据同步，这个感觉不错。
 */
func yars_test2() {
	for i := 0; i < 5; i ++ {
		wg.Add(1)
		go func(n int) {
			// defer wg.Done()
			defer wg.Add(-1)
			time.Sleep(3e9)
			fmt.Println(i)
		}(i)
	}

	wg.Wait()
}

func main() {
	//yars_test1()
	//buffered_chan()
	moreGoRun()
}
