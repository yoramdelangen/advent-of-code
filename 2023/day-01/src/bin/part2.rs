fn main() {
    let inpt = include_str!("input.txt");

    let total: i32 = inpt
        .split("\n")
        .map(|l| {
            if l.len() == 0 {
                return 0;
            }

            print!("Line: {}", l);
            let mut build = String::from("");
            let mut nums: Vec<String> = vec![];
            for c in l.chars() {
                build.push(c);
                if c.is_digit(10) {
                    // get all numbers from the line
                    nums.push(c.to_string());
                    continue;
                } else if let Some(found) = get_spelled_digit(&build) {
                    // the last character of the previous spelled digit
                    // could be the first letter of the next one...
                    let last = build.chars().last().unwrap();
                    build.clear();
                    build.push(last);
                    nums.push(found);
                    continue;
                }
            }

            // if there is some chars left...
            if build.len() > 0 {
                if let Some(found) = get_spelled_digit(&build) {
                    nums.push(found);
                }
            }

            print!(" |{}|", nums.join(","));

            if nums.len() == 0 {
                println!("Skip line because its empty");
                return 0;
            }

            let grab = nums.first().unwrap().to_owned() + nums.last().unwrap();

            println!(" == found: {}", grab);

            grab.parse::<i32>().unwrap()
        })
        .sum();

    println!("Summed number: {}", total);
}

// convert written numbers into actual numbers and return the new line
fn get_spelled_digit(line: &String) -> Option<String> {
    let digits = vec![
        "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
    ];

    for (i, digit) in digits.iter().enumerate() {
        let num = (i + 1).to_string();
        if line.contains(digit) {
            return Some(num);
        }
    }

    return None;
}
