pub mod part2;

use crate::file_reader::get_buf_reader;

pub struct Rucksack {
    compartment1: u64,
    compartment2: u64,
}

impl Rucksack {
    pub fn new(value: &str) -> Self {
        let (compartment1, compartment2) = Self::value_to_content(value);

        Self {
            compartment1,
            compartment2,
        }
    }

    fn value_to_content(value: &str) -> (u64, u64) {
        let mut compartment1 = 0_u64;
        let mut compartment2 = 0_u64;

        let mut chars = value.chars();
        let compartment_size = value.len() / 2;
        for _ in 0..compartment_size {
            let c = chars.next().unwrap();
            let value = get_priority(c);
            compartment1 |= 1 << value;

            let c = chars.next_back().unwrap();
            let value2 = get_priority(c);
            compartment2 |= 1 << value2;
        }

        (compartment1, compartment2)
    }

    pub fn apears_in_both(&self) -> u64 {
        self.compartment1 & self.compartment2
    }
}

fn get_priority(item: char) -> u64 {
    let a_lowercase_value = 'a' as u64;
    let a_uppercase_value = 'A' as u64;
    item as u64
        - if item.is_uppercase() {
            a_uppercase_value - 26
        } else {
            a_lowercase_value
        }
}

pub fn sum_items(rucksacks: Vec<Rucksack>) -> u32 {
    let mut sum = 0_u32;
    for sack in rucksacks {
        let sack_items = sack.apears_in_both();
        sum += sum_sack_items(sack_items);
    }

    sum
}

fn sum_sack_items(sack: u64) -> u32 {
    let mut sum = 0_u32;
    for i in 0..52 {
        if (1 << i & sack) > 0 {
            sum += i + 1;
        }
    }
    sum
}

pub fn read_data(filepath: &str) -> Vec<Rucksack> {
    let buf_reader = get_buf_reader(filepath);

    let mut rucksacks = Vec::<Rucksack>::new();
    for line in buf_reader {
        if let Ok(line) = line {
            rucksacks.push(Rucksack::new(&line));
        }
    }

    rucksacks
}
