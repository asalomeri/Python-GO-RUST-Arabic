# Web Scraper بسيط - جمع البيانات من الويب

import urllib.request
import urllib.parse
from html.parser import HTMLParser


class مستخرج_الروابط(HTMLParser):
    """كلاس لاستخراج الروابط من صفحة HTML"""

    def __init__(self):
        super().__init__()
        self.روابط = []

    def handle_starttag(self, tag, attrs):
        if tag == "a":
            for اسم, قيمة in attrs:
                if اسم == "href" and قيمة:
                    self.روابط.append(قيمة)


def جلب_صفحة(رابط):
    """جلب محتوى صفحة ويب"""
    try:
        headers = {"User-Agent": "Mozilla/5.0 (Python Web Scraper)"}
        طلب = urllib.request.Request(رابط, headers=headers)
        with urllib.request.urlopen(طلب, timeout=10) as استجابة:
            return استجابة.read().decode("utf-8")
    except Exception as خطأ:
        print(f"خطأ في جلب الصفحة: {خطأ}")
        return None


def استخراج_الروابط(محتوى_html):
    """استخراج الروابط من محتوى HTML"""
    مستخرج = مستخرج_الروابط()
    مستخرج.feed(محتوى_html)
    return مستخرج.روابط


def فلترة_الروابط(روابط, كلمة_مفتاحية=""):
    """فلترة الروابط"""
    روابط_صالحة = [r for r in روابط if r.startswith("http")]
    if كلمة_مفتاحية:
        روابط_صالحة = [r for r in روابط_صالحة if كلمة_مفتاحية in r]
    return list(set(روابط_صالحة))  # إزالة التكرار


def main():
    print("=== Web Scraper بسيط ===\n")
    # مثال باستخدام موقع عام
    رابط = "https://example.com"
    print(f"جلب البيانات من: {رابط}")

    محتوى = جلب_صفحة(رابط)
    if محتوى:
        روابط = استخراج_الروابط(محتوى)
        روابط_مفلترة = فلترة_الروابط(روابط)

        print(f"\nتم العثور على {len(روابط_مفلترة)} رابط فريد:")
        for رابط_وجد in روابط_مفلترة[:10]:
            print(f"  - {رابط_وجد}")
    else:
        print("فشل في جلب البيانات")


if __name__ == "__main__":
    main()
