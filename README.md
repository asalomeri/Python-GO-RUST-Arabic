# 🚀 Python-GO-RUST-Arabic

> شرح شامل للغات Python و Go و Rust من الصفر مع مشاريع عملية باللغة العربية

[![Python Tests](https://github.com/asalomeri/Python-GO-RUST-Arabic/actions/workflows/python-tests.yml/badge.svg)](https://github.com/asalomeri/Python-GO-RUST-Arabic/actions/workflows/python-tests.yml)
[![Go Tests](https://github.com/asalomeri/Python-GO-RUST-Arabic/actions/workflows/go-tests.yml/badge.svg)](https://github.com/asalomeri/Python-GO-RUST-Arabic/actions/workflows/go-tests.yml)
[![Rust Tests](https://github.com/asalomeri/Python-GO-RUST-Arabic/actions/workflows/rust-tests.yml/badge.svg)](https://github.com/asalomeri/Python-GO-RUST-Arabic/actions/workflows/rust-tests.yml)

---

## 📋 جدول المحتويات

- [مقدمة](#مقدمة)
- [هيكل المستودع](#هيكل-المستودع)
- [أمثلة سريعة](#أمثلة-سريعة)
- [المقارنة بين اللغات](#المقارنة-بين-اللغات)
- [كيفية البدء](#كيفية-البدء)
- [المشاريع العملية](#المشاريع-العملية)
- [المساهمة](#المساهمة)
- [الموارد والمراجع](#الموارد-والمراجع)

---

## مقدمة

هذا المستودع يحتوي على دروس تعليمية شاملة لثلاث لغات برمجة رائدة:

| اللغة | الاستخدام الرئيسي | مستوى الصعوبة |
|-------|------------------|---------------|
| 🐍 Python | الذكاء الاصطناعي، ويب، أتمتة | سهل |
| 🐹 Go | الخوادم، الأنظمة الموزعة | متوسط |
| 🦀 Rust | الأنظمة، التطبيقات عالية الأداء | صعب |

---

## هيكل المستودع

```
Python-GO-RUST-Arabic/
├── python/
│   ├── 01-basics/          # الأساسيات
│   ├── 02-control-flow/    # التحكم في التدفق
│   ├── 03-functions/       # الدوال
│   ├── 04-oop/             # البرمجة كائنية التوجه
│   └── projects/           # مشاريع عملية
├── go/
│   ├── 01-basics/          # الأساسيات
│   ├── 02-control-flow/    # التحكم في التدفق
│   ├── 03-functions/       # الدوال والحزم
│   ├── 04-goroutines/      # التزامن
│   └── projects/           # مشاريع عملية
├── rust/
│   ├── 01-basics/          # الأساسيات
│   ├── 02-ownership/       # الملكية والاستعارة
│   ├── 03-functions/       # الدوال ومعالجة الأخطاء
│   ├── 04-structs/         # الهياكل والسمات
│   └── projects/           # مشاريع عملية
├── docs/                   # التوثيق
├── tests/                  # الاختبارات
└── .github/workflows/      # CI/CD
```

---

## أمثلة سريعة

### 🐍 Python - مرحبا بالعالم

```python
# أبسط برنامج بايثون
print("مرحبا بالعالم!")

# متغيرات
اسم = "أحمد"
عمر = 25
print(f"مرحبا يا {اسم}، عمرك {عمر} سنة")

# دالة بسيطة
def مجموع(أ, ب):
    return أ + ب

print(مجموع(5, 3))  # 8
```

### 🐹 Go - مرحبا بالعالم

```go
package main

import "fmt"

func main() {
    // أبسط برنامج Go
    fmt.Println("مرحبا بالعالم!")

    // متغيرات
    اسم := "أحمد"
    عمر := 25
    fmt.Printf("مرحبا يا %s، عمرك %d سنة\n", اسم, عمر)
}
```

### 🦀 Rust - مرحبا بالعالم

```rust
fn main() {
    // أبسط برنامج Rust
    println!("مرحبا بالعالم!");

    // متغيرات
    let اسم = "أحمد";
    let عمر = 25;
    println!("مرحبا يا {}، عمرك {} سنة", اسم, عمر);
}
```

---

## المقارنة بين اللغات

| المعيار | Python | Go | Rust |
|---------|--------|-----|------|
| سهولة التعلم | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐ |
| الأداء | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| أمان الذاكرة | ⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| التزامن | ⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ |
| حجم المجتمع | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐ |

---

## كيفية البدء

### 1. تثبيت الأدوات

```bash
# Python
# تحميل من: https://www.python.org/downloads/
python --version  # التحقق

# Go
# تحميل من: https://golang.org/dl/
go version  # التحقق

# Rust
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
rustc --version  # التحقق
```

### 2. استنساخ المستودع

```bash
git clone https://github.com/asalomeri/Python-GO-RUST-Arabic.git
cd Python-GO-RUST-Arabic
```

### 3. تشغيل المثال الأول

```bash
# Python
python python/01-basics/hello_world.py

# Go
go run go/01-basics/hello_world.go

# Rust
rustc rust/01-basics/hello_world.rs && ./hello_world
```

---

## المشاريع العملية

### 🐍 مشاريع Python
| المشروع | الوصف | المسار |
|---------|-------|--------|
| آلة حاسبة | حاسبة تفاعلية | `python/projects/calculator/` |
| قائمة المهام | تطبيق Todo مع حفظ البيانات | `python/projects/todo_app/` |
| Web Scraper | استخراج بيانات من الويب | `python/projects/web_scraper/` |

### 🐹 مشاريع Go
| المشروع | الوصف | المسار |
|---------|-------|--------|
| HTTP Server | خادم ويب بسيط | `go/projects/http_server/` |
| مدير الملفات | عمليات الملفات والمجلدات | `go/projects/file_manager/` |
| REST API | واجهة برمجية | `go/projects/rest_api/` |

### 🦀 مشاريع Rust
| المشروع | الوصف | المسار |
|---------|-------|--------|
| CLI Tool | أداة سطر الأوامر | `rust/projects/cli_tool/` |
| System Monitor | مراقب الموارد | `rust/projects/system_monitor/` |
| لعبة | لعبة حجر ورقة مقص | `rust/projects/game/` |

---

## المساهمة

نرحب بمساهماتكم! يرجى قراءة [دليل المساهمة](CONTRIBUTING.md) قبل البدء.

---

## الموارد والمراجع

📚 راجع [docs/RESOURCES.md](docs/RESOURCES.md) للحصول على قائمة شاملة بالموارد التعليمية.

⚙️ راجع [docs/SETUP.md](docs/SETUP.md) لإعداد بيئة التطوير.

---

**صُنع بـ ❤️ للمجتمع العربي البرمجي**
