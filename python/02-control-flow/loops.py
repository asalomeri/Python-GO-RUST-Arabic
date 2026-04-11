# الحلقات في بايثون - Loops

# ===== حلقة for =====
print("=== حلقة for ===")
for i in range(5):
    print(f"العدد: {i}")

# التكرار على قائمة
فواكه = ["تفاح", "موز", "برتقال"]
for فاكهة in فواكه:
    print(f"الفاكهة: {فاكهة}")

# التكرار مع الفهرس
for فهرس, فاكهة in enumerate(فواكه):
    print(f"{فهرس}: {فاكهة}")

# ===== حلقة while =====
print("\n=== حلقة while ===")
عداد = 0
while عداد < 5:
    print(f"العداد: {عداد}")
    عداد += 1

# ===== break و continue =====
print("\n=== break و continue ===")
for i in range(10):
    if i == 3:
        continue  # تخطي 3
    if i == 7:
        break     # إيقاف عند 7
    print(i)

# ===== حلقة for مع else =====
print("\n=== for مع else ===")
for i in range(5):
    print(i)
else:
    print("انتهت الحلقة بدون break!")

# ===== list comprehension =====
print("\n=== list comprehension ===")
مربعات = [x**2 for x in range(1, 6)]
print(f"المربعات: {مربعات}")

أعداد_زوجية = [x for x in range(20) if x % 2 == 0]
print(f"الأعداد الزوجية: {أعداد_زوجية}")

# ===== الحلقات المتداخلة =====
print("\n=== الحلقات المتداخلة ===")
for i in range(1, 4):
    for j in range(1, 4):
        print(f"{i} × {j} = {i*j}")
