# المتغيرات في بايثون - Variables

# تعريف المتغيرات
اسم = "أحمد"
عمر = 25
طول = 1.75
متزوج = False

# طباعة المتغيرات
print(f"الاسم: {اسم}")
print(f"العمر: {عمر}")
print(f"الطول: {طول} متر")
print(f"متزوج: {متزوج}")

# تغيير قيمة المتغير
عمر = 26
print(f"العمر بعد التغيير: {عمر}")

# تعدد التعيين
x = y = z = 10
print(f"x={x}, y={y}, z={z}")

# تبادل القيم
a, b = 5, 10
print(f"قبل التبادل: a={a}, b={b}")
a, b = b, a
print(f"بعد التبادل: a={a}, b={b}")

# معرفة نوع المتغير
print(type(اسم))    # <class 'str'>
print(type(عمر))    # <class 'int'>
print(type(طول))    # <class 'float'>
print(type(متزوج))  # <class 'bool'>
