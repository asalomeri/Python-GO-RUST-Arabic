// HTTP Server بسيط في Go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// هيكل الاستجابة
type استجابة struct {
	رسالة string `json:"message"`
	وقت   string `json:"time"`
}

// معالج الصفحة الرئيسية
func معالج_الرئيسية(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "مرحبا بالعالم! هذا HTTP Server مكتوب بـ Go\n")
}

// معالج صفحة المعلومات
func معالج_المعلومات(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	استجابة_json := map[string]string{
		"message": "مرحبا من Go HTTP Server",
		"time":    time.Now().Format("2006-01-02 15:04:05"),
		"version": "1.0.0",
	}
	json.NewEncoder(w).Encode(استجابة_json)
}

// معالج صفحة الصحة
func معالج_الصحة(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"status": "OK"})
}

func main() {
	// تسجيل المعالجات
	http.HandleFunc("/", معالج_الرئيسية)
	http.HandleFunc("/info", معالج_المعلومات)
	http.HandleFunc("/health", معالج_الصحة)

	المنفذ := ":8080"
	fmt.Printf("الخادم يعمل على http://localhost%s\n", المنفذ)
	fmt.Println("المسارات المتاحة:")
	fmt.Println("  GET /       - الصفحة الرئيسية")
	fmt.Println("  GET /info   - معلومات الخادم")
	fmt.Println("  GET /health - فحص الصحة")

	log.Fatal(http.ListenAndServe(المنفذ, nil))
}
