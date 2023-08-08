package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// 数组定义
	var a [5]int
	var b [3]int

	b[0] = 3
	b[1] = 6
	b[2] = 9
	fmt.Println(a)
	fmt.Println(b)

	var num1 = [3]int{1, 2, 3}
	var num2 = [...]string{"tom", "jack", "mary"}
	num3 := [3]float64{1.23, 3.24, 56.11}
	fmt.Println(num1, num2, num3)

	for i, v := range num3 {
		fmt.Printf("index=%d value=%v\n", i, v)
	}

	for i := 0; i < len(num3); i++ {
		fmt.Println(num3[i])
	}

	// 随机生成5个数，并将其反转
	// 1.随机生成5个数 rand.Intn()
	//2.将随机数放到一个数组
	//3.反转打印
	var intArr [5]int
	rand.NewSource(time.Now().UnixNano()) // 生成随机数种子
	for i := 0; i < len(intArr); i++ {
		intArr[i] = rand.Intn(100)
	}
	fmt.Println("反转前", intArr)
	// 反转打印
	for i := 0; i < len(intArr)/2; i++ {
		temp := intArr[i]
		intArr[i] = intArr[len(intArr)-1-i]
		intArr[len(intArr)-1-i] = temp
	}
	fmt.Println("反转后", intArr)
}
