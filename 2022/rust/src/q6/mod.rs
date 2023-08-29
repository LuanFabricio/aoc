pub mod part1;
pub mod part2;

fn is_start_sequence(package: &String, mut start: usize, package_size: usize) -> bool {
    let mut stack = 0_u32;
    let mut package_chars = package.chars();
    let end = start + package_size;
    if start > 0 {
        package_chars.nth(start - 1);
    }

    while start < end {
        let char_index = package_chars.next().unwrap() as u8 - 'a' as u8;

        let char_bit = 1 << char_index;
        if (stack & char_bit) != 0 {
            return false;
        }
        stack |= char_bit;
        start += 1
    }

    true
}
