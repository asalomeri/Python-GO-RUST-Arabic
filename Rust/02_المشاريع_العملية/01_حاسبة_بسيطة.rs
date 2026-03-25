fn main() {
    let mut x = String::new();
    let mut y = String::new();
    
    println!("ادخل الرقم الأول:");
    std::io::stdin().read_line(&mut x).expect("فشل في القراءة");
    
    println!("ادخل الرقم الثاني:");
    std::io::stdin().read_line(&mut y).expect("فشل في القراءة");
    
    let num1: i32 = x.trim().parse().expect("يرجى إدخال عدد صحيح");
    let num2: i32 = y.trim().parse().expect("يرجى إدخال عدد صحيح");
    
    println!("النتيجة : {} + {} = {} ", num1, num2, num1 + num2);
    println!("النتيجة : {} - {} = {} ", num1, num2, num1 - num2);
    println!("النتيجة : {} * {} = {} ", num1, num2, num1 * num2);
    println!("النتيجة : {} / {} = {} ", num1, num2, num1 / num2);
}