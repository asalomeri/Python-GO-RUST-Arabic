// أنواع البيانات في Go - Data Types
package main

import (
	"fmt"
	"strings"
)

func main() {
	// ===== الأعداد الصحيحة =====
	fmt.Println("=== الأعداد الصحيحة ===")
	var عدد8 int8 = 127
	var عدد16 int16 = 32767
	var عدد32 int32 = 2147483647
	var عدد64 int64 = 9223372036854775807
	fmt.Printf("int8: %d, int16: %d, int32: %d, int64: %d\n", عدد8, عدد16, عدد32, عدد64)

	// ===== الأعداد العشرية =====
	fmt.Println("\n=== الأعداد العشرية ===")
	var عشري32 float32 = 3.14
	var عشري64 float64 = 3.141592653589793
	fmt.Printf("float32: %.2f, float64: %.15f\n", عشري32, عشري64)

	// ===== النصوص =====
	fmt.Println("\n=== النصوص ===")
	نص := "مرحبا بالعالم"
	fmt.Println(نص)
	fmt.Printf("الطول: %d\n", len([]rune(نص)))
	fmt.Println(strings.ToUpper("hello"))
	fmt.Println(strings.Contains(نص, "مرحبا"))

	// ===== المصفوفات =====
	fmt.Println("\n=== المصفوفات ===")
	var أعداد [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println(أعداد)
	fmt.Printf("الطول: %d\n", len(أعداد))

	// ===== الشرائح (Slices) =====
	fmt.Println("\n=== الشرائح ===")
	فواكه := []string{"تفاح", "موز", "برتقال"}
	fmt.Println(فواكه)
	فواكه = append(فواكه, "مانجو")
	fmt.Println(فواكه)
	fmt.Println(فواكه[1:3])

	// ===== الخرائط (Maps) =====
	fmt.Println("\n=== الخرائط ===")
	شخص := map[string]interface{}{
		"اسم": "محمد",
		"عمر": 30,
	}
	fmt.Println(شخص["اسم"])

	// ===== القيم المنطقية =====
	fmt.Println("\n=== القيم المنطقية ===")
	صحيح := true
	خاطئ := false
	fmt.Printf("صحيح: %v, خاطئ: %v\n", صحيح, خاطئ)
	fmt.Printf("AND: %v, OR: %v, NOT: %v\n", صحيح && خاطئ, صحيح || خاطئ, !صحيح)
}
