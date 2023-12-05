use std::collections::HashMap;

#[derive(Hash, Eq, PartialEq, Debug)]
enum Colors {
    Blue,
    Red,
    Green,
}

const MAX_BLUE: u32 = 14;
const MAX_RED: u32 = 12;
const MAX_GREEN: u32 = 13;

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
            print!("Counting: {:?}", color_count);

            // make sure all colors do not exceed give max numbers
            if color_count.get(&Colors::Green).unwrap() <= &MAX_GREEN
                && color_count.get(&Colors::Blue).unwrap() <= &MAX_BLUE
                && color_count.get(&Colors::Red).unwrap() <= &MAX_RED
            {
                println!(" ✅");
                return game;
            }

            println!("");

            0
        })
        .sum();

    println!("Total: {}", total);
}
