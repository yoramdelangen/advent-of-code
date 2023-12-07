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
}

impl Parser {
    fn has_adjacent_symbols(&self, char: &char, index: usize) -> bool {
        // TODO: determine which positions we need to validate
        // in case we find a direct symbol, we can determine its true
        let list = &self.content;

        println!("Checking: {} on position: {}", char, index);

        let mut sides: Vec<Positions> = vec![];

        // TODO: check horizontal directly
        // do the left side check
        if index > 0 && index % self.line_size != 0 {
            sides.push(Positions::Left);
        }
        // are we at the end of a line?
        if (index % self.line_size) != (self.line_size - 1) {
            sides.push(Positions::Right);
        }

        // TODO: verticle checking for symbols
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
                Positions::Top => self.is_symbol_char(&list[index - self.line_size]),
                Positions::Bottom => self.is_symbol_char(&list[index + self.line_size]),
                Positions::Right => self.is_symbol_char(&list[index + 1]),
                Positions::Left => self.is_symbol_char(&list[index - 1]),
                // direct diagnal checks
                Positions::TopLeft => self.is_symbol_char(&list[index - self.line_size - 1]),
                Positions::TopRight => self.is_symbol_char(&list[index - self.line_size + 1]),
                Positions::BottomLeft => self.is_symbol_char(&list[index + self.line_size - 1]),
                Positions::BottomRight => self.is_symbol_char(&list[index + self.line_size + 1]),
            };

            print!("{:?}: {} ", s, found);

            if found {
                print!(" Found it! ");
                return true;
            }
        }

        false
    }

    fn is_symbol_char(&self, char: &char) -> bool {
        !char.is_numeric() && char != &'.'
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

    let parser = &Parser {
        content: content.clone(),
        line_size: file_content.split_once("\n").unwrap().0.chars().count(),
    };

    let line_size = parser.line_size as usize;

    for (i, char) in content.iter().enumerate() {
        // only the first 2 lines...

        if char.is_digit(10) {
            println!("");
            print!("Char: {:?} ", char);
            is_tracking_digit = true;
            tracking_digit += &char.to_string();
            if parser.has_adjacent_symbols(char, i) {
                print!("✅");
                has_adjacent = true;
            }
        }

        // TODO: is next a digit as well, otherwise we should save it..
        if is_tracking_digit
            && ((i % line_size == (&content.len() - 1)) || !&content[i + 1].is_digit(10))
        {
            is_tracking_digit = false;

            // keep track of all digits
            if has_adjacent {
                digits.push(tracking_digit.parse::<i32>().unwrap());
            }
            tracking_digit.clear();
            has_adjacent = false;
            println!("");
        }

        // match char {
        //     '*' | '$' | '#' | '+' => continue,
        //     '.' | _ => continue,
        //     '0'..='9' => has_adjacent_symbols(char, i, &content),
        // };
    }

    println!("");
    println!("achiefed: {:?}", digits);

    println!("Sum of adjacent numbers: {}", digits.iter().sum::<i32>());
}
