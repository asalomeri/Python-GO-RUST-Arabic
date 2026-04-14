# دليل المساهمة — Python-GO-RUST-Arabic

شكرًا لرغبتك في المساهمة! هذا الدليل يوضح أسلوب العمل المعتمد بشكل منظم.

---

## 1) المتطلبات الأساسية

تأكد من تثبيت الأدوات التالية قبل البدء:

| الأداة | الإصدار المقترح |
|--------|----------------|
| Git | أي إصدار حديث |
| Python | 3.10+ |
| Go | 1.21+ |
| Rust | stable (عبر `rustup`) |

---

## 2) طريقة Fork والاستنساخ

```bash
# 1) اضغط Fork على GitHub لنسخ المستودع إلى حسابك
# 2) استنسخ نسختك المحلية
git clone https://github.com/<اسم-المستخدم>/Python-GO-RUST-Arabic.git
cd Python-GO-RUST-Arabic

# 3) أضف المستودع الأصلي كـ upstream
git remote add upstream https://github.com/asalomeri/Python-GO-RUST-Arabic.git
```

---

## 3) تسمية الفروع (Branch Naming)

أنشئ فرعًا جديدًا لكل مهمة وسمّه بصيغة واضحة:

| البادئة | الاستخدام |
|---------|-----------|
| `feat/<وصف>` | ميزة جديدة |
| `fix/<وصف>` | إصلاح خطأ |
| `docs/<وصف>` | تعديل توثيق |
| `chore/<وصف>` | صيانة / تنظيف |

```bash
# مثال
git checkout -b feat/add-go-goroutines-example
```

---

## 4) اصطلاحات رسائل الـ Commit

استخدم صيغة **Conventional Commits** المبسطة:

```
<نوع>: <وصف قصير>
```

| النوع | المعنى |
|-------|--------|
| `feat` | ميزة جديدة |
| `fix` | إصلاح خطأ |
| `docs` | تعديل توثيق فقط |
| `chore` | صيانة، تحديث اعتمادات |
| `refactor` | إعادة هيكلة بدون تغيير وظيفي |
| `test` | إضافة / تعديل اختبارات |

**أمثلة:**
```
feat: add goroutines example to Go section
fix: correct off-by-one error in calculator.py
docs: update SETUP.md with Rust installation steps
chore: bump dependencies in requirements.txt
```

---

## 5) تشغيل الاختبارات محليًا

### Python
```bash
# تثبيت الاعتمادات
pip install -r requirements.txt

# تشغيل الاختبارات
python -m pytest tests/ -v

# تنسيق الكود
black python/

# فحص الجودة (Lint)
pylint python/ --fail-under=7
```

### Go
```bash
cd go/
go test ./...
gofmt -w .
go vet ./...
```

### Rust
```bash
cd rust/
cargo test
cargo fmt
cargo clippy -- -D warnings
```

---

## 6) التنسيق والفحص (Formatting & Linting)

- **Python:** `black` للتنسيق، `pylint` للفحص.
- **Go:** `gofmt` للتنسيق، `go vet` للفحص.
- **Rust:** `cargo fmt` للتنسيق، `cargo clippy` للفحص.

تأكد من اجتياز جميع الفحوصات قبل فتح PR.

---

## 7) فتح Pull Request

1. ادفع فرعك إلى GitHub:
   ```bash
   git push -u origin feat/<اسم-الفرع>
   ```
2. افتح Pull Request إلى فرع `main`.
3. اتبع قالب الـ PR وأجب على جميع الأسئلة.
4. تأكد من نجاح جميع فحوصات الـ CI قبل طلب المراجعة.

---

## 8) قواعد السلوك (Code of Conduct)

نلتزم بإنشاء بيئة ترحيبية وآمنة للجميع:
- تعامل باحترام مع جميع المساهمين.
- قدم ملاحظات بنّاءة وتجنب النقد الشخصي.
- لا تتردد في طرح الأسئلة؛ كل سؤال له قيمة.

---

*شكرًا لمساهمتك في دعم المحتوى البرمجي العربي!* 🚀
