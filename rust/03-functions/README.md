# الدوال ومعالجة الأخطاء في Rust

## المحتويات

| الملف | الموضوع |
|-------|---------|
| `functions.rs` | الدوال والإغلاقات |
| `error_handling.rs` | معالجة الأخطاء |

## تشغيل الأمثلة

```bash
rustc functions.rs && ./functions
rustc error_handling.rs && ./error_handling
```

## المفاهيم المغطاة

- تعريف الدوال
- التعبيرات مقابل الجمل
- الإغلاقات (Closures)
- map و filter و collect
- `Option<T>` للقيم الاختيارية
- `Result<T, E>` لمعالجة الأخطاء
- معامل `?` للتوزيع السريع
- أنواع الأخطاء المخصصة
