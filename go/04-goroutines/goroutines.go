// Goroutines في Go - التزامن
package main

import (
	"fmt"
	"sync"
	"time"
)

// ===== goroutine بسيطة =====
func طباعة_أرقام(الاسم string, عدد int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= عدد; i++ {
		fmt.Printf("%s: %d\n", الاسم, i)
		time.Sleep(100 * time.Millisecond)
	}
}

// ===== مثال على WaitGroup =====
func مثال_wait_group() {
	fmt.Println("=== WaitGroup ===")
	var wg sync.WaitGroup

	wg.Add(2)
	go طباعة_أرقام("المهمة_أ", 3, &wg)
	go طباعة_أرقام("المهمة_ب", 3, &wg)

	wg.Wait()
	fmt.Println("انتهت جميع المهام!")
}

// ===== Mutex للحماية من race conditions =====
type عداد_آمن struct {
	mu    sync.Mutex
	قيمة int
}

func (ع *عداد_آمن) زيادة() {
	ع.mu.Lock()
	ع.قيمة++
	ع.mu.Unlock()
}

func مثال_mutex() {
	fmt.Println("\n=== Mutex ===")
	عداد := &عداد_آمن{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			عداد.زيادة()
		}()
	}

	wg.Wait()
	fmt.Printf("القيمة النهائية: %d (يجب أن تكون 1000)\n", عداد.قيمة)
}

func main() {
	مثال_wait_group()
	مثال_mutex()
}
