package _case

import (
	"fmt"
	"time"
)

func Communication() {
	//可读可写的通道
	ch := make(chan int)
	go communicationF1(ch)
	go communicationF2(ch)
}

// F1接收一个只写通道
func communicationF1(ch chan<- int) {
	//通过循环向通道写入0-99
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

// F2接收一个只读通道
func communicationF2(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

// 并发场景下的同步机制
func ConcurrentSync() {
	ch := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()

}

// 通知协程退出与多路复用
func NotifyAndMultiplex() {
	ch := make(chan int, 10)
	strCh := make(chan string, 10)
	done := make(chan struct{})

	go noticeAndMultiplexingF1(ch)
	go noticeAndMultiplexingF2(strCh)
	go noticeAndMultiplexingF3(ch, strCh, done)

	time.Sleep(1 * time.Second) // 等待任务完成（生产环境应用 sync.WaitGroup）
	close(done)
}

func noticeAndMultiplexingF1(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}

}

func noticeAndMultiplexingF2(ch chan<- string) {
	for i := 0; i < 10; i++ {
		ch <- fmt.Sprintf("hello %d", i)
	}

}

func noticeAndMultiplexingF3(ch <-chan int, strCh <-chan string, done <-chan struct{}) {
	i := 0
	for {
		select {
		case num, ok := <-ch:
			if !ok { // ch 已关闭
				ch = nil // 将 ch 设为 nil，该 case 将永远阻塞（select 会忽略 nil 通道）
				continue
			}
			fmt.Println(num)
		case str, ok := <-strCh:
			if !ok { // strCh 已关闭
				strCh = nil
				continue
			}
			fmt.Println(str)
		case <-done:
			fmt.Println("收到退出通知")
			return
		}
		i++
		fmt.Println("累计执行次数：", i)
	}
	fmt.Println("结束")
}
