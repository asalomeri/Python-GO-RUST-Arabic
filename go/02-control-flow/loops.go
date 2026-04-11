// الحلقات في Go - Loops
package main

import "fmt"

func main() {
	// ===== حلقة for أساسية =====
	fmt.Println("=== حلقة for أساسية ===")
	for i := 0; i < 5; i++ {
		fmt.Printf("العدد: %d\n", i)
	}

	// ===== for كـ while =====
	fmt.Println("\n=== for كـ while ===")
	عداد := 0
	for عداد < 5 {
		fmt.Printf("العداد: %d\n", عداد)
		عداد++
	}

	// ===== التكرار على شريحة =====
	fmt.Println("\n=== التكرار على شريحة ===")
	فواكه := []string{"تفاح", "موز", "برتقال"}
	for فهرس, فاكهة := range فواكه {
		fmt.Printf("%d: %s\n", فهرس, فاكهة)
	}

	// ===== التكرار على خريطة =====
	fmt.Println("\n=== التكرار على خريطة ===")
	درجات := map[string]int{
		"أحمد":  85,
		"محمد":  92,
		"علي":   78,
	}
	for اسم, درجة := range درجات {
		fmt.Printf("%s: %d\n", اسم, درجة)
	}

	// ===== break و continue =====
	fmt.Println("\n=== break و continue ===")
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue // تخطي 3
		}
		if i == 7 {
			break // إيقاف عند 7
		}
		fmt.Println(i)
	}

	// ===== حلقة لا نهائية =====
	fmt.Println("\n=== حلقة مع عداد ===")
	مجموع := 0
	for i := 1; i <= 100; i++ {
		مجموع += i
	}
	fmt.Printf("مجموع 1 إلى 100 = %d\n", مجموع)
}
