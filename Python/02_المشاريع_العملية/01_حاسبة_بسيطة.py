"""
# ================================
# حاسبة متقدمة مع سجل العمليات
# Advanced Calculator with History
# ================================

import os
from datetime import datetime

class CalculatorEnhanced:
    """حاسبة متقدمة مع ميزات إضافية"""
    
    def __init__(self, history_file='calculator_history.txt'):
        self.history = []
        self.history_file = history_file
        self.load_history()
    
    # ==================== العمليات الأساسية ====================
    def add(self, x, y):
        """جمع رقمين"""
        return x + y
    
    def subtract(self, x, y):
        """طرح رقمين"""
        return x - y
    
    def multiply(self, x, y):
        """ضرب رقمين"""
        return x * y
    
    def divide(self, x, y):
        """قسمة رقمين مع معالجة الأخطاء"""
        if y == 0:
            return "❌ خطأ: لا يمكن القسمة على صفر"
        return x / y
    
    # ==================== عمليات إضافية ====================
    def power(self, x, y):
        """حساب قوة (أس)"""
        return x ** y
    
    def modulo(self, x, y):
        """حساب الباقي من القسمة"""
        if y == 0:
            return "❌ خطأ: لا يمكن حساب الباقي من القسمة على صفر"
        return x % y
    
    def square_root(self, x):
        """حساب الجذر التربيعي"""
        if x < 0:
            return "❌ خطأ: لا يمكن حساب الجذر التربيعي لعدد سالب"
        return x ** 0.5
    
    def percentage(self, x, percent):
        """حساب النسبة المئوية"""
        return (x * percent) / 100
    
    # ==================== إدارة السجل ====================
    def add_to_history(self, operation, num1, num2, result):
        """إضافة عملية للسجل"""
        timestamp = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
        entry = f"[{timestamp}] {operation}: {num1} {num2} = {result}"
        self.history.append(entry)
    
    def display_history(self):
        """عرض سجل العمليات"""
        if not self.history:
            return "📭 لا توجد عمليات حتى الآن"
        
        result = "\n" + "="*60 + "\n"
        result += "📋 سجل العمليات\n"
        result += "="*60 + "\n"
        for i, entry in enumerate(self.history, 1):
            result += f"{i}. {entry}\n"
        result += "="*60 + "\n"
        return result
    
    def clear_history(self):
        """مسح السجل"""
        self.history = []
        self.save_history()
        return "✅ تم مسح السجل"
    
    # ==================== حفظ وتحميل البيانات ====================
    def save_history(self):
        """حفظ السجل في ملف"""
        try:
            with open(self.history_file, 'w', encoding='utf-8') as f:
                for entry in self.history:
                    f.write(entry + '\n')
            return f"✅ تم حفظ السجل في {self.history_file}"
        except Exception as e:
            return f"❌ خطأ في الحفظ: {str(e)}"
    
    def load_history(self):
        """تحميل السجل من ملف"""
        try:
            if os.path.exists(self.history_file):
                with open(self.history_file, 'r', encoding='utf-8') as f:
                    self.history = [line.strip() for line in f if line.strip()]
        except Exception as e:
            print(f"⚠️  تحذير: {str(e)}")

def display_menu():
    """عرض القائمة الرئيسية"""
    print("\n" + "="*60)
    print("🧮 حاسبة متقدمة - Advanced Calculator")
    print("="*60)
    print("العمليات الأساسية:")
    print("  1️⃣  جمع (+)")
    print("  2️⃣  طرح (-)")
    print("  3️⃣  ضرب (*)")
    print("  4️⃣  قسمة (/)")
    print("\nعمليات إضافية:")
    print("  5️⃣  أس / قوة (^)")
    print("  6️⃣  الباقي من القسمة (%)")
    print("  7️⃣  جذر تربيعي (√)")
    print("  8️⃣  نسبة مئوية")
    print("\nخيارات أخرى:")
    print("  9️⃣  عرض السجل")
    print("  0️⃣  مسح السجل")
    print("  Q ⏏️  خروج")
    print("="*60)

def get_input_safe(prompt):
    """الحصول على رقم من المستخدم بشكل آمن"""
    while True:
        try:
            return float(input(prompt))
        except ValueError:
            print("❌ خطأ: يرجى إدخال رقم صحيح")

def main():
    """البرنامج الرئيسي"""
    calculator = CalculatorEnhanced()
    
    print("\n🎉 مرحباً بك في الحاسبة المتقدمة!")
    print("📁 السجل يتم حفظه تلقائياً في: calculator_history.txt\n")
    
    while True:
        display_menu()
        choice = input("\n👉 اختر عملية (1-9, 0, Q): ").strip().upper()
        
        if choice == 'Q':
            calculator.save_history()
            print("\n👋 شكراً لاستخدام الحاسبة! تم حفظ السجل.")
            break
        
        elif choice == '9':
            print(calculator.display_history())
        
        elif choice == '0':
            print(calculator.clear_history())
        
        elif choice in ['1', '2', '3', '4', '5', '6', '8']:
            num1 = get_input_safe("📥 أدخل الرقم الأول: ")
            num2 = get_input_safe("📥 أدخل الرقم الثاني: ")
            
            operations = {
                '1': ('+', calculator.add),
                '2': ('-', calculator.subtract),
                '3': ('*', calculator.multiply),
                '4': ('/', calculator.divide),
                '5': ('^', calculator.power),
                '6': ('%', calculator.modulo),
                '8': ('نسبة', calculator.percentage),
            }
            
            op_symbol, operation = operations[choice]
            result = operation(num1, num2)
            
            print(f"\n✨ النتيجة: {num1} {op_symbol} {num2} = {result}")
            calculator.add_to_history(op_symbol, num1, num2, result)
        
        elif choice == '7':
            num = get_input_safe("📥 أدخل الرقم: ")
            result = calculator.square_root(num)
            print(f"\n✨ النتيجة: √{num} = {result}")
            calculator.add_to_history('√', num, '', result)
        
        else:
            print("❌ اختيار غير صحيح، حاول مرة أخرى")
        
        # حفظ السجل بعد كل عملية
        calculator.save_history()
        
        cont = input("\n❓ هل تريد إجراء عملية أخرى؟ (نعم/yes/y أو أي مفتاح آخر للخروج): ").strip().lower()
        if cont not in ['نعم', 'yes', 'y']:
            calculator.save_history()
            print("\n👋 شكراً لاستخدام الحاسبة!")
            break

if __name__ == "__main__":
    main()
"""