// الدوال في Go - Functions
package main

import (
	"errors"
	"fmt"
	"math"
)

// ===== دالة بسيطة =====
func مرحبا() {
	fmt.Println("مرحبا بالعالم!")
}

// ===== دالة بمعاملات =====
func تحية(الاسم string) {
	fmt.Printf("مرحبا يا %s!\n", الاسم)
}

// ===== دالة بعودة =====
func جمع(أ, ب int) int {
	return أ + ب
}

// ===== دالة تعيد عدة قيم =====
func قسمة(أ, ب float64) (float64, error) {
	if ب == 0 {
		return 0, errors.New("لا يمكن القسمة على صفر")
	}
	return أ / ب, nil
}

// ===== دالة بقيم عودة مسماة =====
func مستطيل(طول, عرض float64) (مساحة, محيط float64) {
	مساحة = طول * عرض
	محيط = 2 * (طول + عرض)
	return
}

// ===== دالة variadic =====
func مجموع(أعداد ...int) int {
	مجموع := 0
	for _, عدد := range أعداد {
		مجموع += عدد
	}
	return مجموع
}

// ===== دالة من نوع أول (First-class function) =====
func تطبيق(أ, ب float64, عملية func(float64, float64) float64) float64 {
	return عملية(أ, ب)
}

// ===== دالة recursive =====
func فيبوناتشي(n int) int {
	if n <= 1 {
		return n
	}
	return فيبوناتشي(n-1) + فيبوناتشي(n-2)
}

func main() {
	مرحبا()
	تحية("أحمد")

	fmt.Printf("3 + 4 = %d\n", جمع(3, 4))

	نتيجة, خطأ := قسمة(10, 2)
	if خطأ != nil {
		fmt.Println("خطأ:", خطأ)
	} else {
		fmt.Printf("10 ÷ 2 = %.2f\n", نتيجة)
	}

	_, خطأ = قسمة(5, 0)
	if خطأ != nil {
		fmt.Println("خطأ:", خطأ)
	}

	مساحة, محيط := مستطيل(5, 3)
	fmt.Printf("المساحة: %.1f, المحيط: %.1f\n", مساحة, محيط)

	fmt.Printf("المجموع: %d\n", مجموع(1, 2, 3, 4, 5))

	// استخدام دالة كمعامل
	حاصل := تطبيق(4, 2, func(أ, ب float64) float64 {
		return math.Pow(أ, ب)
	})
	fmt.Printf("4^2 = %.0f\n", حاصل)

	// فيبوناتشي
	for i := 0; i <= 10; i++ {
		fmt.Printf("فيبوناتشي(%d) = %d\n", i, فيبوناتشي(i))
	}
}
