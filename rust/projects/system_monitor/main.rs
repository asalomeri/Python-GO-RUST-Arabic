// مراقب النظام - System Monitor في Rust
use std::fs;
use std::time::{SystemTime, UNIX_EPOCH};

fn قراءة_uptime() -> Result<String, Box<dyn std::error::Error>> {
    let محتوى = fs::read_to_string("/proc/uptime")?;
    let ثواني: f64 = محتوى
        .split_whitespace()
        .next()
        .ok_or("ملف uptime فارغ")?
        .parse()?;

    let ساعات = (ثواني / 3600.0) as u64;
    let دقائق = ((ثواني % 3600.0) / 60.0) as u64;
    let ثواني_باقية = (ثواني % 60.0) as u64;

    Ok(format!("{}ساعة {}دقيقة {}ثانية", ساعات, دقائق, ثواني_باقية))
}

fn قراءة_معلومات_الذاكرة() -> Result<Vec<(String, u64)>, Box<dyn std::error::Error>> {
    let محتوى = fs::read_to_string("/proc/meminfo")?;
    let mut معلومات = Vec::new();

    for سطر in محتوى.lines().take(5) {
        let أجزاء: Vec<&str> = سطر.split_whitespace().collect();
        if أجزاء.len() >= 2 {
            let مفتاح = أجزاء[0].trim_end_matches(':').to_string();
            if let Ok(قيمة) = أجزاء[1].parse::<u64>() {
                معلومات.push((مفتاح, قيمة));
            }
        }
    }

    Ok(معلومات)
}

fn تحويل_إلى_ميغابايت(كيلوبايت: u64) -> f64 {
    كيلوبايت as f64 / 1024.0
}

fn الوقت_الحالي() -> String {
    let الوقت = SystemTime::now()
        .duration_since(UNIX_EPOCH)
        .unwrap_or_default();
    format!("Unix timestamp: {}", الوقت.as_secs())
}

fn main() {
    println!("╔══════════════════════════════╗");
    println!("║      مراقب النظام - Rust      ║");
    println!("╚══════════════════════════════╝\n");

    // الوقت الحالي
    println!("🕐 الوقت: {}", الوقت_الحالي());

    // وقت التشغيل
    match قراءة_uptime() {
        Ok(وقت) => println!("⏱️  وقت التشغيل: {}", وقت),
        Err(خطأ) => println!("❌ خطأ في قراءة وقت التشغيل: {}", خطأ),
    }

    // معلومات الذاكرة
    println!("\n📊 معلومات الذاكرة:");
    match قراءة_معلومات_الذاكرة() {
        Ok(معلومات) => {
            for (مفتاح, قيمة) in &معلومات {
                println!("  {}: {:.1} MB", مفتاح, تحويل_إلى_ميغابايت(*قيمة));
            }
        }
        Err(خطأ) => println!("❌ خطأ في قراءة الذاكرة: {}", خطأ),
    }

    println!("\n✅ انتهى فحص النظام");
}
