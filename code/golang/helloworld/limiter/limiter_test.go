package limiter_test

import (
	"testing"
	"time"

	"example.com/m/limiter"
)

func TestTickerLimiter_Allow(t *testing.T) {
	rps := 2
	l := limiter.NewTickerLimiter(rps)
	defer l.Close()

	// 测试初始请求
	if l.Allow() {
		t.Error("第一个请求应该被允许")
	}

	// 测试第二个请求(同一时间窗口)
	if l.Allow() {
		t.Error("第二个请求应该被拒绝")
	}

	time.Sleep(time.Second/2 + 10*time.Millisecond)

	if !l.Allow() {
		t.Error("第三个请求应该被允许")
	}
}
