use crate::file_reader::get_buf_reader;

#[derive(Debug, Clone)]
pub struct ElfBag(Vec<u32>);

pub fn most_calories(bags: Vec<ElfBag>) -> u32 {
    let mut max_calories = 0_u32;
    let mut total_bag;

    for bag in bags.iter() {
        total_bag = 0_u32;
        for item in bag.0.iter() {
            total_bag += *item;
        }

        if total_bag > max_calories {
            max_calories = total_bag;
        }
    }

    max_calories
}

pub fn sum_top_3_calories(bags: Vec<ElfBag>) -> u32 {
    let mut top_3_bags: [u32; 3] = [0_u32; 3];
    let mut total_bag: u32;

    for bag in bags.iter() {
        total_bag = 0;

        for item in bag.0.iter() {
            total_bag += *item;
        }

        match (
            total_bag > top_3_bags[0],
            total_bag > top_3_bags[1],
            total_bag > top_3_bags[2],
        ) {
            (true, true, true) => {
                top_3_bags[2] = top_3_bags[1];
                top_3_bags[1] = top_3_bags[0];
                top_3_bags[0] = total_bag;
            }
            (false, true, true) => {
                top_3_bags[2] = top_3_bags[1];
                top_3_bags[1] = total_bag;
            }
            (false, false, true) => {
                top_3_bags[2] = total_bag;
            }
            (_, _, _) => {}
        }
    }

    total_bag = 0;
    for calories in top_3_bags {
        total_bag += calories;
    }

    total_bag
}

pub fn read_data(filepath: &str) -> Vec<ElfBag> {
    let read_buffer = get_buf_reader(filepath);

    let mut bags = Vec::<ElfBag>::new();
    let mut current_bag = Vec::<u32>::new();
    for line in read_buffer {
        if let Ok(line) = line {
            if line.len() == 0 {
                bags.push(ElfBag(current_bag.clone()));
                current_bag.clear();
            } else {
                current_bag.push(line.trim().parse().unwrap());
            }
        }
    }

    bags
}
