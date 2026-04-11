// معالجة الأخطاء في Rust - Error Handling
use std::fs;
use std::num::ParseIntError;
use std::fmt;

// ===== نوع خطأ مخصص =====
#[derive(Debug)]
enum خطأ_حاسبة {
    قسمة_على_صفر,
    عدد_سالب(f64),
}

impl fmt::Display for خطأ_حاسبة {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        match self {
            خطأ_حاسبة::قسمة_على_صفر => write!(f, "لا يمكن القسمة على صفر!"),
            خطأ_حاسبة::عدد_سالب(n) => write!(f, "العدد {} سالب!", n),
        }
    }
}

fn قسمة_آمنة(أ: f64, ب: f64) -> Result<f64, خطأ_حاسبة> {
    if ب == 0.0 {
        Err(خطأ_حاسبة::قسمة_على_صفر)
    } else {
        Ok(أ / ب)
    }
}

fn جذر_تربيعي(x: f64) -> Result<f64, خطأ_حاسبة> {
    if x < 0.0 {
        Err(خطأ_حاسبة::عدد_سالب(x))
    } else {
        Ok(x.sqrt())
    }
}

fn main() {
    // ===== Option =====
    println!("=== Option ===");
    let أعداد = vec![1, 2, 3, 4, 5];
    
    match أعداد.first() {
        Some(عدد) => println!("أول عدد: {}", عدد),
        None => println!("القائمة فارغة"),
    }

    // unwrap_or
    let فارغة: Vec<i32> = vec![];
    let أول = فارغة.first().unwrap_or(&0);
    println!("أول عدد (افتراضي): {}", أول);

    // ===== Result =====
    println!("\n=== Result ===");
    match قسمة_آمنة(10.0, 2.0) {
        Ok(نتيجة) => println!("10 ÷ 2 = {}", نتيجة),
        Err(خطأ) => println!("خطأ: {}", خطأ),
    }

    match قسمة_آمنة(5.0, 0.0) {
        Ok(نتيجة) => println!("النتيجة: {}", نتيجة),
        Err(خطأ) => println!("خطأ: {}", خطأ),
    }

    match جذر_تربيعي(-4.0) {
        Ok(نتيجة) => println!("الجذر: {}", نتيجة),
        Err(خطأ) => println!("خطأ: {}", خطأ),
    }

    // ===== معالج ? =====
    println!("\n=== معالج ? ===");
    match قراءة_رقم("42") {
        Ok(رقم) => println!("الرقم: {}", رقم),
        Err(خطأ) => println!("خطأ في التحويل: {}", خطأ),
    }
}

fn قراءة_رقم(نص: &str) -> Result<i32, ParseIntError> {
    let رقم = نص.trim().parse::<i32>()?;
    Ok(رقم * 2)
}
