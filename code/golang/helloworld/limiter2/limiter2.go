// Package limiter2 show limiter of way 2
package limiter2

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	chLimit := make(chan struct{}, 3)

	for i := 0; i < 20; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			chLimit <- struct{}{}

			time.Sleep(time.Second * 3)

			defer func() { <-chLimit }() // 释放许可

			fmt.Println("下游服务处理逻辑....")

		}(i)
	}

	wg.Wait()
}

// Limiter2 show how to use this func
func Limiter2() {

}
