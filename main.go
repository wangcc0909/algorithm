package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.NumGoroutine()//获取goroutine的数量
	runtime.GOMAXPROCS(runtime.NumCPU())
	array := make([]int, 100000)
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(a []int) {
			now := time.Now()
			print(a)
			fmt.Println("耗时:", time.Since(now))
			wg.Done()
		}(array)
	}
	wg.Wait()
}
