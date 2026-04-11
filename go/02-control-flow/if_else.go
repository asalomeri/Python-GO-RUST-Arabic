// الشروط في Go - If/Else
package main

import "fmt"

func main() {
	// ===== الشرط البسيط =====
	درجة := 75

	if درجة >= 90 {
		fmt.Println("ممتاز!")
	} else if درجة >= 80 {
		fmt.Println("جيد جداً")
	} else if درجة >= 70 {
		fmt.Println("جيد")
	} else if درجة >= 60 {
		fmt.Println("مقبول")
	} else {
		fmt.Println("راسب")
	}

	// ===== if مع تهيئة مسبقة =====
	if x := 10; x > 5 {
		fmt.Printf("%d أكبر من 5\n", x)
	}

	// ===== switch =====
	fmt.Println("\n=== switch ===")
	يوم := "الإثنين"
	switch يوم {
	case "السبت", "الأحد":
		fmt.Println("عطلة نهاية الأسبوع!")
	case "الإثنين":
		fmt.Println("بداية أسبوع عمل جديد")
	case "الجمعة":
		fmt.Println("نهاية الأسبوع!")
	default:
		fmt.Println("يوم عمل عادي")
	}

	// ===== switch بدون شرط (مثل if-else) =====
	درجة_الحرارة := 25
	switch {
	case درجة_الحرارة > 35:
		fmt.Println("حار جداً!")
	case درجة_الحرارة > 25:
		fmt.Println("دافئ")
	case درجة_الحرارة > 15:
		fmt.Println("معتدل")
	default:
		fmt.Println("بارد")
	}
}
