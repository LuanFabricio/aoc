use super::part1::Pair;

pub fn count_full_overlap(pairs: Vec<Pair>) -> u32 {
    let mut count = 0_u32;
    for pair in pairs {
        if pair.full_overlap() || pair.should_change() {
            count += 1;
        }
    }
    count
}
