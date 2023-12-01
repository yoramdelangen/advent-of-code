fn main() {
    let inpt = include_str!("input1.txt");

    let total: i32 = inpt
        .split("\n")
        .map(|line| {
            print!("Line: {}", line);

            let mut nums: Vec<String> = vec![];
            for c in line.chars() {
                if c.is_digit(10) {
                    nums.push(c.to_string())
                }
            }

            if nums.len() == 0 {
                println!("Skip line because its empty");
                return 0;
            }

            // let grab = format!("{:?}{:?}", nums.first().unwrap(), nums.last().unwrap());
            let grab = nums.first().unwrap().to_owned() + nums.last().unwrap();
            
            println!(", found: {}", grab);

            grab.parse::<i32>().unwrap()
        })
        .sum();

    println!("Summed number: {}", total);
}
