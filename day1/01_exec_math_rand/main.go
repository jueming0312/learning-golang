package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 练习随机数生成
	// 1.生成随机数种子
	rand.NewSource(time.Now().UnixNano())
	// 2.生成 0-100之间的随机整数
	fmt.Println(rand.Intn(100))
}
