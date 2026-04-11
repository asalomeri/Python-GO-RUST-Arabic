// المتغيرات في Go - Variables
package main

import "fmt"

func main() {
	// تعريف المتغيرات بشكل صريح
	var اسم string = "أحمد"
	var عمر int = 25
	var طول float64 = 1.75
	var متزوج bool = false

	fmt.Printf("الاسم: %s\n", اسم)
	fmt.Printf("العمر: %d\n", عمر)
	fmt.Printf("الطول: %.2f متر\n", طول)
	fmt.Printf("متزوج: %v\n", متزوج)

	// تعريف مختصر باستخدام :=
	مدينة := "الرياض"
	تاريخ_الميلاد := 1999
	fmt.Printf("المدينة: %s\n", مدينة)
	fmt.Printf("تاريخ الميلاد: %d\n", تاريخ_الميلاد)

	// تعريف متعدد
	var (
		درجة_الحرارة float64 = 25.5
		حالة_الطقس  string  = "مشمس"
		يوم_العمل   bool    = true
	)
	fmt.Printf("درجة الحرارة: %.1f\n", درجة_الحرارة)
	fmt.Printf("حالة الطقس: %s\n", حالة_الطقس)
	fmt.Printf("يوم العمل: %v\n", يوم_العمل)

	// القيم الافتراضية
	var عدد_صحيح int      // 0
	var عدد_عشري float64 // 0.0
	var نص string         // ""
	var منطقي bool        // false
	fmt.Printf("القيم الافتراضية: %d, %.1f, '%s', %v\n", عدد_صحيح, عدد_عشري, نص, منطقي)

	// ثوابت
	const π = 3.14159
	const اسم_التطبيق = "مشروع Go"
	fmt.Printf("π = %.5f\n", π)
	fmt.Printf("اسم التطبيق: %s\n", اسم_التطبيق)
}
