use crate::file_reader::get_buf_reader;

#[derive(Debug, Clone)]
pub struct Pair {
    elf1: u128,
    elf2: u128,
}

impl Pair {
    pub fn new(elf1: (u32, u32), elf2: (u32, u32)) -> Self {
        Self {
            elf1: Self::get_sections(elf1),
            elf2: Self::get_sections(elf2),
        }
    }

    pub fn should_change(&self) -> bool {
        let r = self.elf1 | self.elf2;

        r == self.elf1 || r == self.elf2
    }

    pub fn full_overlap(&self) -> bool {
        let r = self.elf1 & self.elf2;
        r != 0
    }

    fn get_sections(elf: (u32, u32)) -> u128 {
        let mut sections = 0_u128;
        for i in elf.0..elf.1 + 1 {
            sections |= 1 << i;
        }

        sections
    }
}

pub fn read_data(filepath: &str) -> Vec<Pair> {
    let buf_reader = get_buf_reader(filepath);

    let mut pairs = vec![];
    for line in buf_reader {
        if let Ok(line) = line {
            let mut split = line.split(",");

            let mut elf1_split = split.next().unwrap().split("-");
            let elf1 = (
                elf1_split.next().unwrap().parse::<u32>().unwrap(),
                elf1_split.next().unwrap().parse::<u32>().unwrap(),
            );

            let mut elf2_split = split.next().unwrap().split("-");
            let elf2 = (
                elf2_split.next().unwrap().parse::<u32>().unwrap(),
                elf2_split.next().unwrap().parse::<u32>().unwrap(),
            );

            pairs.push(Pair::new(elf1, elf2));
        }
    }

    pairs
}

pub fn count_fully_contains(pairs: Vec<Pair>) -> u32 {
    let mut count = 0_u32;
    for pair in pairs {
        if pair.should_change() {
            count += 1;
        }
    }
    count
}
