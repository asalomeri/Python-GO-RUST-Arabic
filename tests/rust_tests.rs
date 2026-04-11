// اختبارات Rust
#[cfg(test)]
mod اختبارات {
    // ===== دوال مساعدة =====
    fn جمع(أ: i32, ب: i32) -> i32 {
        أ + ب
    }

    fn ضرب(أ: i32, ب: i32) -> i32 {
        أ * ب
    }

    fn قسمة(أ: f64, ب: f64) -> Result<f64, String> {
        if ب == 0.0 {
            Err(String::from("لا يمكن القسمة على صفر"))
        } else {
            Ok(أ / ب)
        }
    }

    fn فيبوناتشي(n: u32) -> u32 {
        match n {
            0 => 0,
            1 => 1,
            _ => فيبوناتشي(n - 1) + فيبوناتشي(n - 2),
        }
    }

    fn هل_عدد_أولي(n: u32) -> bool {
        if n < 2 {
            return false;
        }
        let حد = (n as f64).sqrt() as u32;
        for i in 2..=حد {
            if n % i == 0 {
                return false;
            }
        }
        true
    }

    fn عكس_نص(نص: &str) -> String {
        نص.chars().rev().collect()
    }

    // ===== الاختبارات =====
    #[test]
    fn اختبار_الجمع() {
        assert_eq!(جمع(2, 3), 5);
        assert_eq!(جمع(-1, 1), 0);
        assert_eq!(جمع(0, 0), 0);
        assert_eq!(جمع(-5, -3), -8);
    }

    #[test]
    fn اختبار_الضرب() {
        assert_eq!(ضرب(3, 4), 12);
        assert_eq!(ضرب(-2, 3), -6);
        assert_eq!(ضرب(0, 100), 0);
    }

    #[test]
    fn اختبار_القسمة() {
        assert_eq!(قسمة(10.0, 2.0), Ok(5.0));
        assert_eq!(قسمة(7.0, 2.0), Ok(3.5));
    }

    #[test]
    fn اختبار_القسمة_على_صفر() {
        assert!(قسمة(5.0, 0.0).is_err());
        assert_eq!(
            قسمة(5.0, 0.0),
            Err(String::from("لا يمكن القسمة على صفر"))
        );
    }

    #[test]
    fn اختبار_فيبوناتشي() {
        let متوقعة = vec![0, 1, 1, 2, 3, 5, 8, 13, 21, 34];
        for (i, &متوقع) in متوقعة.iter().enumerate() {
            assert_eq!(فيبوناتشي(i as u32), متوقع);
        }
    }

    #[test]
    fn اختبار_الأعداد_الأولية() {
        let أولية = vec![2, 3, 5, 7, 11, 13];
        let غير_أولية = vec![1, 4, 6, 8, 9, 10];

        for n in أولية {
            assert!(هل_عدد_أولي(n), "{} يجب أن يكون عدداً أولياً", n);
        }

        for n in غير_أولية {
            assert!(!هل_عدد_أولي(n), "{} لا يجب أن يكون عدداً أولياً", n);
        }
    }

    #[test]
    fn اختبار_عكس_النص() {
        assert_eq!(عكس_نص("hello"), "olleh");
        assert_eq!(عكس_نص("rust"), "tsur");
        assert_eq!(عكس_نص(""), "");
        assert_eq!(عكس_نص("a"), "a");
    }
}

fn main() {
    println!("تشغيل: cargo test");
}
