// REST API بسيط في Go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"
)

// هيكل المهمة
type مهمة struct {
	ID        int    `json:"id"`
	العنوان   string `json:"title"`
	المكتملة  bool   `json:"completed"`
	التاريخ   string `json:"created_at"`
}

// قاعدة البيانات في الذاكرة
var (
	mu     sync.Mutex
	مهام   []مهمة
	عداد_ID = 1
)

func init() {
	مهام = []مهمة{
		{ID: 1, العنوان: "تعلم Go", المكتملة: false, التاريخ: time.Now().Format("2006-01-02")},
		{ID: 2, العنوان: "بناء REST API", المكتملة: true, التاريخ: time.Now().Format("2006-01-02")},
	}
	عداد_ID = 3
}

// جلب جميع المهام
func جلب_المهام(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "الطريقة غير مسموحة", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	mu.Lock()
	defer mu.Unlock()
	json.NewEncoder(w).Encode(مهام)
}

// إنشاء مهمة جديدة
func إنشاء_مهمة(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "الطريقة غير مسموحة", http.StatusMethodNotAllowed)
		return
	}

	var مهمة_جديدة مهمة
	if خطأ := json.NewDecoder(r.Body).Decode(&مهمة_جديدة); خطأ != nil {
		http.Error(w, "بيانات غير صحيحة", http.StatusBadRequest)
		return
	}

	mu.Lock()
	مهمة_جديدة.ID = عداد_ID
	عداد_ID++
	مهمة_جديدة.التاريخ = time.Now().Format("2006-01-02")
	مهام = append(مهام, مهمة_جديدة)
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(مهمة_جديدة)
}

// معالج موحد للمهام
func معالج_المهام(w http.ResponseWriter, r *http.Request) {
	// استخراج المعرف من المسار
	أجزاء := strings.Split(r.URL.Path, "/")
	if len(أجزاء) < 3 || أجزاء[2] == "" {
		switch r.Method {
		case http.MethodGet:
			جلب_المهام(w, r)
		case http.MethodPost:
			إنشاء_مهمة(w, r)
		default:
			http.Error(w, "الطريقة غير مسموحة", http.StatusMethodNotAllowed)
		}
		return
	}

	معرف, خطأ := strconv.Atoi(أجزاء[2])
	if خطأ != nil {
		http.Error(w, "معرف غير صحيح", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()
	for i, م := range مهام {
		if م.ID == معرف && r.Method == http.MethodGet {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(مهام[i])
			return
		}
	}
	http.Error(w, "المهمة غير موجودة", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/tasks", معالج_المهام)
	http.HandleFunc("/tasks/", معالج_المهام)

	المنفذ := ":8081"
	fmt.Printf("REST API يعمل على http://localhost%s\n", المنفذ)
	fmt.Println("\nنقاط النهاية:")
	fmt.Println("  GET  /tasks      - جلب جميع المهام")
	fmt.Println("  POST /tasks      - إنشاء مهمة جديدة")
	fmt.Println("  GET  /tasks/{id} - جلب مهمة محددة")

	log.Fatal(http.ListenAndServe(المنفذ, nil))
}
