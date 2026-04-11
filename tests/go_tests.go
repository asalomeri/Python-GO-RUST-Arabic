// اختبارات Go
package main

import (
	"fmt"
	"math"
	"testing"
)

// ===== دوال مساعدة للاختبار =====
func جمع(أ, ب int) int {
	return أ + ب
}

func ضرب(أ, ب int) int {
	return أ * ب
}

func قسمة(أ, ب float64) (float64, error) {
	if ب == 0 {
		return 0, fmt.Errorf("لا يمكن القسمة على صفر")
	}
	return أ / ب, nil
}

func فيبوناتشي(n int) int {
	if n <= 1 {
		return n
	}
	return فيبوناتشي(n-1) + فيبوناتشي(n-2)
}

func هل_عدد_أولي(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// ===== الاختبارات =====
func TestJamع(t *testing.T) {
	cases := []struct {
		أ, ب, متوقع int
	}{
		{2, 3, 5},
		{-1, 1, 0},
		{0, 0, 0},
		{-5, -3, -8},
	}

	for _, c := range cases {
		نتيجة := جمع(c.أ, c.ب)
		if نتيجة != c.متوقع {
			t.Errorf("جمع(%d, %d) = %d; المتوقع %d", c.أ, c.ب, نتيجة, c.متوقع)
		}
	}
}

func TestDarab(t *testing.T) {
	cases := []struct {
		أ, ب, متوقع int
	}{
		{3, 4, 12},
		{-2, 3, -6},
		{0, 100, 0},
	}

	for _, c := range cases {
		نتيجة := ضرب(c.أ, c.ب)
		if نتيجة != c.متوقع {
			t.Errorf("ضرب(%d, %d) = %d; المتوقع %d", c.أ, c.ب, نتيجة, c.متوقع)
		}
	}
}

func TestFibonacci(t *testing.T) {
	متوقعة := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
	for i, متوقع := range متوقعة {
		نتيجة := فيبوناتشي(i)
		if نتيجة != متوقع {
			t.Errorf("فيبوناتشي(%d) = %d; المتوقع %d", i, نتيجة, متوقع)
		}
	}
}

func TestPrime(t *testing.T) {
	أولية := []int{2, 3, 5, 7, 11, 13}
	غير_أولية := []int{1, 4, 6, 8, 9, 10}

	for _, n := range أولية {
		if !هل_عدد_أولي(n) {
			t.Errorf("%d يجب أن يكون عدداً أولياً", n)
		}
	}

	for _, n := range غير_أولية {
		if هل_عدد_أولي(n) {
			t.Errorf("%d لا يجب أن يكون عدداً أولياً", n)
		}
	}
}

func main() {}

