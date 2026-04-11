# البرمجة كائنية التوجه في بايثون - OOP

# ===== تعريف كلاس بسيط =====
class حيوان:
    # متغير الكلاس
    عدد_الحيوانات = 0

    def __init__(self, اسم, صوت):
        self.اسم = اسم
        self.صوت = صوت
        حيوان.عدد_الحيوانات += 1

    def يصدر_صوت(self):
        print(f"{self.اسم} يقول: {self.صوت}")

    def __str__(self):
        return f"حيوان: {self.اسم}"

    def __repr__(self):
        return f"حيوان('{self.اسم}', '{self.صوت}')"


# ===== إنشاء كائنات =====
قطة = حيوان("قطة", "مياو")
كلب = حيوان("كلب", "هاو")

قطة.يصدر_صوت()
كلب.يصدر_صوت()
print(f"عدد الحيوانات: {حيوان.عدد_الحيوانات}")
print(قطة)

# ===== كلاس مع خصائص (Properties) =====
class دائرة:
    def __init__(self, نصف_قطر):
        self._نصف_قطر = نصف_قطر

    @property
    def نصف_القطر(self):
        return self._نصف_قطر

    @نصف_القطر.setter
    def نصف_القطر(self, قيمة):
        if قيمة < 0:
            raise ValueError("نصف القطر لا يمكن أن يكون سالباً")
        self._نصف_قطر = قيمة

    @property
    def المساحة(self):
        import math
        return math.pi * self._نصف_قطر ** 2

    @property
    def المحيط(self):
        import math
        return 2 * math.pi * self._نصف_قطر


دائرة1 = دائرة(5)
print(f"نصف القطر: {دائرة1.نصف_القطر}")
print(f"المساحة: {دائرة1.المساحة:.2f}")
print(f"المحيط: {دائرة1.المحيط:.2f}")
