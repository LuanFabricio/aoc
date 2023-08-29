use crate::file_reader::get_buf_reader;

#[derive(Debug)]
pub struct Group {
    sacks: [u64; 3],
}

impl Group {
    pub fn new(sacks: &[String; 3]) -> Self {
        Self {
            sacks: [
                Self::convert_sack(sacks[0].as_str()),
                Self::convert_sack(sacks[1].as_str()),
                Self::convert_sack(sacks[2].as_str()),
            ],
        }
    }

    fn convert_sack(sack: &str) -> u64 {
        let mut final_result = 0_u64;
        for item in sack.chars() {
            let value = super::get_priority(item);
            final_result |= 1 << value;
        }

        final_result
    }

    pub fn group_item(&self) -> Result<u32, String> {
        let binary_value = self.sacks[0] & self.sacks[1] & self.sacks[2];

        for i in 0..52 {
            let char_value = 1 << i & binary_value;
            if char_value > 0 {
                return Ok(i + 1);
            }
        }

        Err("Item not found".to_string())
    }
}

pub fn read_data(filepath: &str) -> Vec<Group> {
    let buf_reader = get_buf_reader(filepath);

    let mut groups = Vec::<Group>::new();
    let mut count = 0_usize;
    let mut current_group: [String; 3] = [String::from(""), String::from(""), String::from("")];
    for line in buf_reader {
        if let Ok(line) = line {
            count += 1;
            current_group[count - 1] = line.clone();
            if count >= 3 {
                groups.push(Group::new(&current_group));
                count = 0;
            }
        }
    }

    groups
}

pub fn sum_groups_item(groups: Vec<Group>) -> u32 {
    let mut sum = 0_u32;
    for group in groups {
        if let Ok(item) = group.group_item() {
            sum += item;
        }
    }

    sum
}
