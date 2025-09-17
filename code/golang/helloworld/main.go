package main

import (
	"fmt"

	"github.com/gogf/gf/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

func main() {

	fmt.Println("Hello goFrame!", gf.VERSION)

	s := g.Server()
	s.BindHandler("/", func(r *ghttp.Request) {
		r.Response.Write("Hello goFrame!")
	})
	s.SetPort(8199)
	s.Run()
}
