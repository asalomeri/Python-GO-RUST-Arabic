# Goroutines والقنوات في Go

## المحتويات

| الملف | الموضوع |
|-------|---------|
| `goroutines.go` | Goroutines والتزامن |
| `channels.go` | القنوات للتواصل |

## تشغيل الأمثلة

```bash
go run goroutines.go
go run channels.go
```

## المفاهيم المغطاة

- إنشاء goroutines بكلمة `go`
- `sync.WaitGroup` للانتظار
- `sync.Mutex` لحماية البيانات
- القنوات البسيطة
- القنوات ذات السعة (Buffered)
- `select` للاستماع لعدة قنوات
