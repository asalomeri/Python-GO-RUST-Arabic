// الحزم في Go - Packages
package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func main() {
	// ===== حزمة fmt =====
	fmt.Println("=== حزمة fmt ===")
	fmt.Printf("النص: %s\n", "مرحبا")
	fmt.Printf("العدد: %d\n", 42)
	fmt.Printf("العشري: %.2f\n", 3.14)

	رسالة := fmt.Sprintf("تمت المعالجة في %s", "2024")
	fmt.Println(رسالة)

	// ===== حزمة math =====
	fmt.Println("\n=== حزمة math ===")
	fmt.Printf("π = %.5f\n", math.Pi)
	fmt.Printf("الجذر التربيعي لـ 16 = %.0f\n", math.Sqrt(16))
	fmt.Printf("2^10 = %.0f\n", math.Pow(2, 10))
	fmt.Printf("القيمة المطلقة لـ -5 = %.0f\n", math.Abs(-5))

	// ===== حزمة strings =====
	fmt.Println("\n=== حزمة strings ===")
	نص := "مرحبا بالعالم"
	fmt.Println(strings.ToUpper("hello world"))
	fmt.Println(strings.Contains(نص, "مرحبا"))
	fmt.Println(strings.Replace("a-b-c", "-", " ", -1))
	أجزاء := strings.Split("a,b,c", ",")
	fmt.Println(أجزاء)
	fmt.Println(strings.Join(أجزاء, " | "))

	// ===== حزمة sort =====
	fmt.Println("\n=== حزمة sort ===")
	أعداد := []int{5, 2, 8, 1, 9, 3}
	sort.Ints(أعداد)
	fmt.Println(أعداد)

	كلمات := []string{"موز", "تفاح", "برتقال"}
	sort.Strings(كلمات)
	fmt.Println(كلمات)

	// ===== حزمة time =====
	fmt.Println("\n=== حزمة time ===")
	الآن := time.Now()
	fmt.Println(الآن.Format("2006-01-02 15:04:05"))

	// ===== حزمة rand =====
	fmt.Println("\n=== حزمة rand ===")
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 5; i++ {
		fmt.Printf("رقم عشوائي: %d\n", rng.Intn(100))
	}
}
