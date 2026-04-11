// مدير الملفات - File Manager في Go
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// قراءة محتوى ملف
func قراءة_ملف(مسار string) (string, error) {
	محتوى, خطأ := os.ReadFile(مسار)
	if خطأ != nil {
		return "", fmt.Errorf("خطأ في قراءة الملف: %w", خطأ)
	}
	return string(محتوى), nil
}

// كتابة محتوى لملف
func كتابة_ملف(مسار, محتوى string) error {
	خطأ := os.WriteFile(مسار, []byte(محتوى), 0644)
	if خطأ != nil {
		return fmt.Errorf("خطأ في كتابة الملف: %w", خطأ)
	}
	return nil
}

// عرض محتوى مجلد
func عرض_المجلد(مسار string) ([]string, error) {
	مدخلات, خطأ := os.ReadDir(مسار)
	if خطأ != nil {
		return nil, fmt.Errorf("خطأ في قراءة المجلد: %w", خطأ)
	}

	var ملفات []string
	for _, مدخل := range مدخلات {
		نوع := "📄"
		if مدخل.IsDir() {
			نوع = "📁"
		}
		ملفات = append(ملفات, fmt.Sprintf("%s %s", نوع, مدخل.Name()))
	}
	return ملفات, nil
}

// البحث عن ملفات
func بحث_عن_ملفات(مجلد, امتداد string) ([]string, error) {
	var نتائج []string
	خطأ := filepath.Walk(مجلد, func(مسار string, info os.FileInfo, خطأ error) error {
		if خطأ != nil {
			return خطأ
		}
		if !info.IsDir() && strings.HasSuffix(مسار, امتداد) {
			نتائج = append(نتائج, مسار)
		}
		return nil
	})
	return نتائج, خطأ
}

func main() {
	fmt.Println("=== مدير الملفات ===\n")

	// إنشاء ملف تجريبي
	محتوى_تجريبي := "مرحبا! هذا ملف تجريبي من Go\nالسطر الثاني\nالسطر الثالث"
	خطأ := كتابة_ملف("/tmp/test_file.txt", محتوى_تجريبي)
	if خطأ != nil {
		fmt.Println("خطأ:", خطأ)
		return
	}
	fmt.Println("✅ تم إنشاء الملف بنجاح")

	// قراءة الملف
	محتوى, خطأ := قراءة_ملف("/tmp/test_file.txt")
	if خطأ != nil {
		fmt.Println("خطأ:", خطأ)
		return
	}
	fmt.Printf("\nمحتوى الملف:\n%s\n", محتوى)

	// عرض محتوى المجلد
	fmt.Println("\nمحتوى /tmp:")
	ملفات, خطأ := عرض_المجلد("/tmp")
	if خطأ != nil {
		fmt.Println("خطأ:", خطأ)
		return
	}
	for _, ملف := range ملفات[:min(5, len(ملفات))] {
		fmt.Println(" ", ملف)
	}
}

func min(أ, ب int) int {
	if أ < ب {
		return أ
	}
	return ب
}
