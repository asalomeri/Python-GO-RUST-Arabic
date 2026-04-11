# اختبارات Python
import sys
import os

# إضافة المسار للوحدات
sys.path.insert(0, os.path.join(os.path.dirname(__file__), '..', 'python', 'projects', 'calculator'))

import pytest


# ===== اختبارات الآلة الحاسبة =====
def test_جمع():
    from calculator import جمع
    assert جمع(2, 3) == 5
    assert جمع(-1, 1) == 0
    assert جمع(0, 0) == 0


def test_طرح():
    from calculator import طرح
    assert طرح(5, 3) == 2
    assert طرح(0, 0) == 0
    assert طرح(-1, -1) == 0


def test_ضرب():
    from calculator import ضرب
    assert ضرب(3, 4) == 12
    assert ضرب(-2, 3) == -6
    assert ضرب(0, 100) == 0


def test_قسمة():
    from calculator import قسمة
    assert قسمة(10, 2) == 5.0
    assert قسمة(7, 2) == 3.5


def test_قسمة_على_صفر():
    from calculator import قسمة
    with pytest.raises(ValueError, match="لا يمكن القسمة على صفر"):
        قسمة(5, 0)


# ===== اختبارات الدوال الأساسية =====
def test_أنواع_البيانات():
    """اختبار أنواع البيانات الأساسية"""
    اسم = "أحمد"
    عمر = 25
    طول = 1.75
    
    assert isinstance(اسم, str)
    assert isinstance(عمر, int)
    assert isinstance(طول, float)


def test_القوائم():
    """اختبار عمليات القوائم"""
    قائمة = [1, 2, 3, 4, 5]
    assert len(قائمة) == 5
    assert قائمة[0] == 1
    assert قائمة[-1] == 5
    
    قائمة.append(6)
    assert len(قائمة) == 6
    assert 6 in قائمة


def test_القاموس():
    """اختبار عمليات القواميس"""
    شخص = {"اسم": "محمد", "عمر": 30}
    assert شخص["اسم"] == "محمد"
    assert شخص.get("عمر") == 30
    assert شخص.get("وظيفة") is None


def test_الحلقات():
    """اختبار الحلقات"""
    مجموع = sum(range(1, 6))
    assert مجموع == 15
    
    مربعات = [x**2 for x in range(1, 6)]
    assert مربعات == [1, 4, 9, 16, 25]


def test_الدوال_lambda():
    """اختبار دوال lambda"""
    ضعف = lambda x: x * 2
    assert ضعف(5) == 10
    
    مربع = lambda x: x ** 2
    assert مربع(4) == 16


if __name__ == "__main__":
    pytest.main([__file__, "-v"])
