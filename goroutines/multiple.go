package goroutines

import (
	"fmt"
	"time"
)

// func PrintNumbers(num int) {
// 	fmt.Println(num)
// }

func Multiple() {
	for i := 1; i <= 5; i++ {
		go func(n int) {
			fmt.Println(n)
		}(i)
	}
	time.Sleep(1 * time.Second)
}
