use crate::file_reader::get_buf_reader;

pub fn read_data(filepath: &str) -> (Vec<String>, Vec<CommandMove>) {
    let buf_reader = get_buf_reader(filepath);

    let mut table = Vec::<String>::new();
    let mut finished_table = false;
    let mut commands = Vec::<CommandMove>::new();
    for line in buf_reader {
        if let Ok(line) = line {
            match (line.contains("1"), finished_table) {
                (false, false) => {
                    table.push(line);
                }
                (true, false) => {
                    table.push(line);
                    finished_table = true
                }
                (_, _) => {
                    if line.len() > 0 {
                        commands.push(CommandMove::new(line));
                    }
                }
            };
        }
    }

    (table, commands)
}

#[derive(Debug, Clone)]
pub struct Crate {
    stack: Vec<char>,
}

impl Crate {
    pub fn new() -> Self {
        Self { stack: vec![] }
    }

    pub fn add(&mut self, item: char) {
        if item != ' ' {
            self.stack.push(item);
        }
    }

    pub fn string_table_to_crates(mut table: Vec<String>) -> Vec<Self> {
        let mut crates = vec![];

        let quant_crates = table.last().unwrap().split("  ").count();
        table.pop();

        for _ in 0..quant_crates {
            crates.push(Crate::new());
        }

        for line in table.iter().rev() {
            let items = Self::get_items(line, quant_crates);
            for i in 0..items.len() {
                crates[i].add(items[i]);
            }
        }

        crates
    }

    fn get_items(line: &String, quant_crates: usize) -> Vec<char> {
        let mut items = vec![];
        let mut line_chars = line.chars();

        for i in 0..quant_crates {
            let item_position = if i == 0 { 1 } else { 3 };
            if let Some(item) = line_chars.nth(item_position) {
                items.push(item);
            } else {
                items.push(' ');
            }
        }

        items
    }
}

pub struct CommandMove {
    quant: usize,
    source: usize,
    target: usize,
}

impl CommandMove {
    pub fn new(command: String) -> Self {
        let mut split = command.trim().split(" ");

        let quant = split.nth(1).expect("[Command]: Quantity not found!");
        let source = split.nth(1).expect("[Command]: Source not found!");
        let target = split.nth(1).expect("[Command]: Target not found!");

        Self {
            quant: quant.parse().expect("[Command]: Quant parse fail."),
            source: source.parse().expect("[Command]: Source parse fail."),
            target: target.parse().expect("[Command]: Target parse fail."),
        }
    }

    pub fn apply_on_crates(&self, crates: &mut Vec<Crate>, preseve_order: bool) {
        let mut items = Vec::<char>::new();

        if let Some(c) = crates.get_mut(self.source - 1) {
            for _ in 0..self.quant {
                let item = c.stack.pop().expect("Cant pop.");
                items.push(item);
            }
        }

        if let Some(c) = crates.get_mut(self.target - 1) {
            for i in 0..items.len() {
                let index = if preseve_order {
                    items.len() - 1 - i
                } else {
                    i
                };
                c.add(items[index]);
            }
        }
    }
}

pub fn apply_all_commands(
    commands: &Vec<CommandMove>,
    crates: &mut Vec<Crate>,
    preseve_order: bool,
) {
    for command in commands {
        command.apply_on_crates(crates, preseve_order);
    }
}

pub fn get_top_of_stack(crates: &Vec<Crate>) -> String {
    let mut items = String::from("");

    for c in crates.iter() {
        items += c
            .stack
            .last()
            .expect("Cant get last item.")
            .to_string()
            .as_str();
    }

    items
}
