package limiter

import (
	"context"
	"sync"
	"time"
)

type TickerLimiter struct {
	ticker *time.Ticker
	mu     sync.Mutex
	closed bool
}

// 创建新的限流器，rps 表示每秒请求数
func NewTickerLimiter(rps int) *TickerLimiter {
	if rps <= 0 {
		rps = 1 // 默认每秒1个请求
	}

	interval := time.Second / time.Duration(rps)
	return &TickerLimiter{
		ticker: time.NewTicker(interval),
	}
}

func (l *TickerLimiter) Allow() bool {
	select {
	case <-l.ticker.C:
		return true
	default:
		return false
	}
}

// Wait 等待直到可以处理请求
func (l *TickerLimiter) Wait(ctx context.Context) error {
	select {
	case <-l.ticker.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Close 关闭限流器
func (l *TickerLimiter) Close() {
	l.mu.Lock()
	defer l.mu.Unlock()
	if !l.closed {
		l.ticker.Stop()
		l.closed = true
	}
}
