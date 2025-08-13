package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/v2"
	"github.com/gogf/gf/v2/frame/g"
)

func main() {

	// orderCoroutinue()

	// rateLimit()
	fmt.Println("Hello goFrame!", gf.VERSION)

	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("Hello goFrame!")
	})
	s.SetPort(8199)
	s.Run()
}

func orderCoroutinue() {
	chs := make([]chan int, 4)

	for i := range chs {
		fmt.Println(i)

		chs[i] = make(chan int)

		go func(i int) {
			for {
				v := <-chs[i]
				fmt.Println(v + 1)
				time.Sleep(time.Second)

				chs[(i+1)%4] <- (v + 1) % 4
			}
		}(i)
	}
	chs[0] <- 0
	select {}
}

// 限流器
func rateLimit() {
	var wg sync.WaitGroup

	chLimit := make(chan struct{}, 3)
	for i := 0; i < 20; i++ {
		chLimit <- struct{}{}
		wg.Add(1)

		go func(i int) {
			fmt.Println("下游服务处理逻辑...", i)
			time.Sleep(time.Second * 3)
			<-chLimit
			wg.Done()
		}(i)
	}
	wg.Wait()
}
