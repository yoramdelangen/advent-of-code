use std::collections::HashMap;

#[derive(Hash, Eq, PartialEq, Debug)]
enum Colors {
    Blue,
    Red,
    Green,
}

fn main() {
    let input = include_str!("input.txt");

    let total: u32 = input
        .lines()
        .map(|line| {
            // println!("Part: {}", line)
            let parts = line.split(":").collect::<Vec<_>>();
            let game = parts
                .first()
                .unwrap()
                .replace("Game ", "")
                .parse::<u32>()
                .unwrap();

            // println!("last: {:?}", parts.last().unwrap().split(";"));

            let mut color_count: HashMap<Colors, u32> = HashMap::new();
            color_count.insert(Colors::Green, 0);
            color_count.insert(Colors::Red, 0);
            color_count.insert(Colors::Blue, 0);

            parts.last().unwrap().trim().split(";").for_each(|set| {
                set.trim().split(",").into_iter().for_each(|play| {
                    let mut m = play.split_whitespace();
                    let d = m.next().unwrap().parse::<u32>().unwrap();

                    let t = match m.next().unwrap() {
                        "green" => Colors::Green,
                        "red" => Colors::Red,
                        "blue" => Colors::Blue,
                        _ => return (),
                    };

                    // make the hashmap value the bigger value
                    if &d > color_count.get(&t).unwrap() {
                        *color_count.get_mut(&t).unwrap() = d;
                    }
                });
            });

            print!("Game id: {:?} ", game);
            print!("Counting: {:?} ", color_count);

            let game_sum = color_count.get(&Colors::Green).unwrap() *  
                color_count.get(&Colors::Blue).unwrap() * 
                color_count.get(&Colors::Red).unwrap();

            println!("Game sum: {}", game_sum);

            game_sum
        })
        .sum();

    println!("Total: {}", total);
}
