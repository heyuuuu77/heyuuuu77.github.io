package limiter2_test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestLimit2(t *testing.T) {
	tests := []struct {
		name           string
		concurrency    int
		totalRequest   int
		expectedMaxCon int
		expectedTime   time.Duration
	}{
		{
			name:           "3并发限制",
			concurrency:    3,
			totalRequest:   20,
			expectedMaxCon: 3,
			expectedTime:   21 * time.Second,
		},
		{
			name:           "5并发限制",
			concurrency:    5,
			totalRequest:   20,
			expectedMaxCon: 5,
			expectedTime:   12 * time.Second,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var (
				currentConcurrency int
				maxConcurrency     int
				mu                 sync.Mutex
				wg                 sync.WaitGroup
			)

			chLimit := make(chan struct{}, tt.concurrency)

			start := time.Now()

			for i := 0; i < tt.totalRequest; i++ {
				chLimit <- struct{}{}

				mu.Lock()
				currentConcurrency++
				if currentConcurrency > maxConcurrency {
					maxConcurrency = currentConcurrency
				}

				mu.Unlock()

				wg.Add(1)
				go func() {
					defer wg.Done()
					defer func() {
						<-chLimit
						mu.Lock()
						currentConcurrency--
						mu.Unlock()
					}()

					time.Sleep(3 * time.Second)
				}()
			}
			wg.Wait()
			elapsed := time.Since(start)

			if maxConcurrency != tt.expectedMaxCon {
				t.Errorf("最大并发数不匹配； 实际 %d, 期望 %d", maxConcurrency, tt.expectedMaxCon)
			}

			if elapsed < tt.expectedTime-1*time.Second || elapsed > tt.expectedTime+1*time.Second {
				t.Errorf("最大并发数不匹配； 实际 %d, 期望 %d", maxConcurrency, tt.expectedMaxCon)
			}

			fmt.Printf("%s: 最大并发数=%d, 总耗时=%v\n", tt.name, maxConcurrency, elapsed)
		})
	}
}
