// القنوات في Go - Channels
package main

import (
	"fmt"
	"time"
)

// ===== قناة بسيطة =====
func مثال_قناة_بسيطة() {
	fmt.Println("=== قناة بسيطة ===")
	قناة := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		قناة <- "مرحبا من goroutine!"
	}()

	رسالة := <-قناة
	fmt.Println(رسالة)
}

// ===== قناة ذات سعة (Buffered) =====
func مثال_قناة_ذات_سعة() {
	fmt.Println("\n=== قناة ذات سعة ===")
	قناة := make(chan int, 3)

	قناة <- 1
	قناة <- 2
	قناة <- 3

	fmt.Println(<-قناة)
	fmt.Println(<-قناة)
	fmt.Println(<-قناة)
}

// ===== إرسال أعداد عبر قناة =====
func توليد_أعداد(قناة chan<- int, حتى int) {
	for i := 1; i <= حتى; i++ {
		قناة <- i
	}
	close(قناة)
}

func مثال_توليد() {
	fmt.Println("\n=== توليد الأعداد ===")
	قناة := make(chan int)
	go توليد_أعداد(قناة, 5)

	for عدد := range قناة {
		fmt.Printf("استقبلت: %d\n", عدد)
	}
}

// ===== select =====
func مثال_select() {
	fmt.Println("\n=== select ===")
	قناة1 := make(chan string)
	قناة2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		قناة1 <- "من القناة 1"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		قناة2 <- "من القناة 2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case رسالة := <-قناة1:
			fmt.Println(رسالة)
		case رسالة := <-قناة2:
			fmt.Println(رسالة)
		}
	}
}

func main() {
	مثال_قناة_بسيطة()
	مثال_قناة_ذات_سعة()
	مثال_توليد()
	مثال_select()
}
