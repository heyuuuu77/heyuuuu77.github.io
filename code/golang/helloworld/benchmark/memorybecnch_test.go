package main

import "testing"

// 原始结构体：字段顺序可能导致内存浪费和缓存效率低
type User struct {
	Age     int32 // 4字节
	IsAdmin bool  // 1字节（会因对齐产生3字节填充）
	Score   int64 // 8字节（需从8字节边界开始，前面会产生4字节填充）
}

// 优化结构体：按字段大小排序（从大到小），减少内存填充
type UserOptimized struct {
	Score   int64 // 8字节（放在最前，无填充）
	Age     int32 // 4字节（紧随8字节后，无填充）
	IsAdmin bool  // 1字节（最后，仅可能有3字节填充，总填充减少）
}

func BenchmarkUser(b *testing.B) {
	users := make([]User, 1000000)
	for i := 0; i < b.N; i++ {
		for _, u := range users {
			_ = u.Age + int32(u.Score)
		}
	}
}

// BenchmarkUser-8             4597            249822 ns/op            3482 B/op          0 allocs/op

func BenchmarkUserOptimized(b *testing.B) {
	users := make([]UserOptimized, 1000000)
	for i := 0; i < b.N; i++ {
		for _, u := range users {
			_ = u.Age + int32(u.Score)
		}
	}
}

// BenchmarkUserOptimized-8            4448            264298 ns/op            3598 B/op          0 allocs/op
