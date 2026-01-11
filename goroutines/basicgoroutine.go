package goroutines

import (
	"fmt"
	"time"
)

func SayHello() {
	fmt.Println("Hello go!")
}
func BasicGoroutine() {
	go SayHello()

	time.Sleep(1 * time.Second)

	fmt.Println("main func ends")
}
