// أداة سطر الأوامر - CLI Tool في Rust
use std::env;
use std::process;

fn طباعة_مساعدة() {
    println!("أداة سطر الأوامر - CLI Tool");
    println!("\nالاستخدام:");
    println!("  cli_tool <أمر> [خيارات]");
    println!("\nالأوامر:");
    println!("  مرحبا <اسم>    - طباعة ترحيب");
    println!("  جمع <أ> <ب>    - جمع عددين");
    println!("  عكس <نص>       - عكس النص");
    println!("  مساعدة         - عرض هذه المساعدة");
}

fn أمر_مرحبا(args: &[String]) {
    if args.len() < 3 {
        println!("مرحبا! (لم يُحدد اسم)");
    } else {
        println!("مرحبا يا {}! 👋", args[2]);
    }
}

fn أمر_جمع(args: &[String]) {
    if args.len() < 4 {
        eprintln!("خطأ: تحتاج إلى عددين");
        process::exit(1);
    }
    
    let أ: f64 = match args[2].parse() {
        Ok(n) => n,
        Err(_) => {
            eprintln!("خطأ: '{}' ليس رقماً صحيحاً", args[2]);
            process::exit(1);
        }
    };
    
    let ب: f64 = match args[3].parse() {
        Ok(n) => n,
        Err(_) => {
            eprintln!("خطأ: '{}' ليس رقماً صحيحاً", args[3]);
            process::exit(1);
        }
    };
    
    println!("{} + {} = {}", أ, ب, أ + ب);
}

fn أمر_عكس(args: &[String]) {
    if args.len() < 3 {
        eprintln!("خطأ: تحتاج إلى نص");
        process::exit(1);
    }
    let نص_معكوس: String = args[2].chars().rev().collect();
    println!("النص الأصلي: {}", args[2]);
    println!("النص المعكوس: {}", نص_معكوس);
}

fn main() {
    let args: Vec<String> = env::args().collect();
    
    if args.len() < 2 {
        طباعة_مساعدة();
        process::exit(0);
    }
    
    match args[1].as_str() {
        "مرحبا" => أمر_مرحبا(&args),
        "جمع" => أمر_جمع(&args),
        "عكس" => أمر_عكس(&args),
        "مساعدة" | "--help" | "-h" => طباعة_مساعدة(),
        أمر_غير_معروف => {
            eprintln!("خطأ: أمر غير معروف '{}'", أمر_غير_معروف);
            طباعة_مساعدة();
            process::exit(1);
        }
    }
}
