# دليل الإعداد - SETUP.md

## متطلبات النظام

### Python
```bash
# تحقق من إصدار Python
python --version  # يجب 3.8+

# تثبيت التبعيات
pip install -r requirements.txt
```

### Go
```bash
# تحقق من إصدار Go
go version  # يجب 1.21+

# تنزيل Go من
# https://golang.org/dl/
```

### Rust
```bash
# تثبيت Rust
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh

# تحقق من الإصدار
rustc --version  # يجب 1.70+
cargo --version
```

## إعداد بيئة التطوير

### Visual Studio Code
1. تثبيت الإضافات المقترحة:
   - Python (ms-python.python)
   - Go (golang.go)
   - rust-analyzer (rust-lang.rust-analyzer)

### تشغيل الأمثلة

#### Python
```bash
cd python/01-basics
python hello_world.py
```

#### Go
```bash
cd go/01-basics
go run hello_world.go
```

#### Rust
```bash
cd rust/01-basics
rustc hello_world.rs
./hello_world
```

## تشغيل الاختبارات

### Python
```bash
pytest tests/python_tests.py -v
```

### Go
```bash
go test ./...
```

### Rust
```bash
cargo test
```
