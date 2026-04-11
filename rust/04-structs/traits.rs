// السمات في Rust - Traits
use std::fmt;

// ===== تعريف سمة =====
trait صوت {
    fn يصدر_صوت(&self) -> String;
    
    fn وصف(&self) -> String {
        format!("أنا {}", self.يصدر_صوت())
    }
}

trait حيوان: صوت {
    fn اسم(&self) -> &str;
}

// ===== تطبيق السمات =====
struct قطة {
    اسم: String,
}

struct كلب {
    اسم: String,
}

impl صوت for قطة {
    fn يصدر_صوت(&self) -> String {
        String::from("مياو!")
    }
}

impl صوت for كلب {
    fn يصدر_صوت(&self) -> String {
        String::from("هاو!")
    }
}

impl حيوان for قطة {
    fn اسم(&self) -> &str {
        &self.اسم
    }
}

impl حيوان for كلب {
    fn اسم(&self) -> &str {
        &self.اسم
    }
}

impl fmt::Display for قطة {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "قطة اسمها {}", self.اسم)
    }
}

// ===== الدوال العامة (Generics) مع السمات =====
fn اجعله_يصوت<T: صوت>(حيوان: &T) {
    println!("الصوت: {}", حيوان.يصدر_صوت());
}

fn main() {
    let قطتي = قطة { اسم: String::from("ميسي") };
    let كلبي = كلب { اسم: String::from("ريكس") };

    println!("=== الصوت ===");
    اجعله_يصوت(&قطتي);
    اجعله_يصوت(&كلبي);

    println!("\n=== الوصف ===");
    println!("{}", قطتي.وصف());
    println!("{}", كلبي.وصف());

    println!("\n=== Display ===");
    println!("{}", قطتي);

    // ===== trait objects =====
    println!("\n=== trait objects ===");
    let حيوانات: Vec<Box<dyn صوت>> = vec![
        Box::new(قطة { اسم: String::from("كيتي") }),
        Box::new(كلب { اسم: String::from("باكي") }),
    ];

    for حيوان in &حيوانات {
        println!("{}", حيوان.يصدر_صوت());
    }
}
