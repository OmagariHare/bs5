package core

import (
	"math/rand"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var rander = rand.New(rand.NewSource(time.Now().UnixNano()))

// 生成随机字符串
func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rander.Intn(len(letterBytes))]
	}
	return string(b)
}
