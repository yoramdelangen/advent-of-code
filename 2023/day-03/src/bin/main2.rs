use std::collections::HashMap;

#[derive(Debug)]
enum Positions {
    Top,
    Bottom,
    Left,
    Right,
    TopLeft,
    TopRight,
    BottomLeft,
    BottomRight,
}

struct Parser {
    content: Vec<char>,
    line_size: usize,
    // key: index location of the potential gear
    // value: vector of values
    gears: HashMap<usize, Vec<i32>>,
}

impl Parser {
    fn adjacent_symbols(&self, char: &char, index: usize) -> Option<usize> {
        // determine which positions we need to validate
        // in case we find a direct symbol, we can determine its true
        let list = &self.content;

        println!("Checking: {} on position: {}", char, index);

        let mut sides: Vec<Positions> = vec![];

        // Check horizontal directly
        // do the left side check
        if index > 0 && index % self.line_size != 0 {
            sides.push(Positions::Left);
        }
        // are we at the end of a line?
        if (index % self.line_size) != (self.line_size - 1) {
            sides.push(Positions::Right);
        }

        // Verticle checking for symbols
        // can we check the above?
        if index > self.line_size {
            sides.push(Positions::Top);

            if index % self.line_size != 0 {
                sides.push(Positions::TopLeft);
            }
            if (index % self.line_size) != (self.line_size - 1) {
                sides.push(Positions::TopRight);
            }
        }

        if index < (list.len() - self.line_size) {
            sides.push(Positions::Bottom);
            if index % self.line_size != 0 {
                sides.push(Positions::BottomLeft);
            }
            if (index % self.line_size) != (self.line_size - 1) {
                sides.push(Positions::BottomRight);
            }
        }

        println!("Positions to check: {:?}", sides);

        for s in sides {
            let found = match s {
                Positions::Top => self.potential_gear_index(index - self.line_size),
                Positions::Bottom => self.potential_gear_index(index + self.line_size),
                Positions::Right => self.potential_gear_index(index + 1),
                Positions::Left => self.potential_gear_index(index - 1),
                // direct diagnal checks
                Positions::TopLeft => self.potential_gear_index(index - self.line_size - 1),
                Positions::TopRight => self.potential_gear_index(index - self.line_size + 1),
                Positions::BottomLeft => self.potential_gear_index(index + self.line_size - 1),
                Positions::BottomRight => self.potential_gear_index(index + self.line_size + 1),
            };

            print!("{:?}: {:?} ", s, found);

            if found.is_some() {
                print!("Found it! ");
                return found;
            }
        }

        None
    }

    fn potential_gear_index(&self, index: usize) -> Option<usize> {
        let char = self.content[index];

        if &char == &'*' {
            return Some(index);
        }

        // !char.is_numeric() && char != &'.'
        None
    }

    fn add_gear(&mut self, gear_index: usize, digit: i32) {
        let _ = &self.gears.entry(gear_index).or_insert(vec![]).push(digit);
    }

    fn get_gears_sum(&self) -> HashMap<usize, i32> {
        self.gears
            .iter()
            // only gears with only 2 numbers allowed
            .filter(|g| g.1.len() == 2)
            .map(|g| -> (usize, i32) { (*g.0, g.1[0] * g.1[1]) })
            .collect()
    }
}

fn main() {
    let file_content = include_str!("input.txt");

    let content: Vec<char> = file_content.trim().replace('\n', "").chars().collect();

    // println!("Line length: {:?}", LINE_SIZE);
    // println!("Lines: {:?}", content);

    let mut tracking_digit = "".to_string();
    let mut is_tracking_digit = false;
    let mut has_adjacent = false;
    let mut digits: Vec<i32> = vec![];
    let mut gear: Option<usize> = None;

    let mut parser: Parser = Parser {
        content: content.clone(),
        line_size: file_content.split_once("\n").unwrap().0.chars().count(),
        gears: HashMap::new(),
    };

    let line_size = parser.line_size as usize;

    for (i, char) in content.iter().enumerate() {
        // only the first 2 lines...

        if char.is_digit(10) {
            println!("");
            print!("Char: {:?} ", char);
            is_tracking_digit = true;
            tracking_digit += &char.to_string();
            if let Some(found_gear) = parser.adjacent_symbols(char, i) {
                print!("✅");
                has_adjacent = true;
                gear = Some(found_gear);
            }
        }

        // is next a digit as well, otherwise we should save it..
        if is_tracking_digit
            && ((i % line_size == (&content.len() - 1)) || !&content[i + 1].is_digit(10))
        {
            is_tracking_digit = false;

            // keep track of all digits
            if has_adjacent {
                let n = tracking_digit.parse::<i32>().unwrap();
                digits.push(n);
                // push into the gears list
                if gear.is_some() {
                    parser.add_gear(gear.unwrap(), n);
                }
            }
            tracking_digit.clear();
            has_adjacent = false;
            println!("");
        }
    }

    let gears = parser.get_gears_sum();

    println!("");
    println!("Gears: {:?}", parser.gears);
    println!("Gears summed: {:?}", gears);
    println!("Gears totals: {:?}", gears.iter().map(|g| g.1).sum::<i32>());
}
