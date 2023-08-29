mod file_reader;
mod q1;
mod q2;
mod q3;
mod q4;
mod q5;
mod q6;
mod q7;

fn main() {
    resolve_01();
    println!();
    resolve_02();
    println!();
    resolve_03();
    println!();
    resolve_04();
    println!();
    resolve_05();
    println!();
    resolve_06();
    println!();
    resolve_07();
    println!();
}

fn resolve_01() {
    println!("Calories");
    let bags = q1::read_data("assets/q1/input.in");

    let most_calories = q1::most_calories(bags.clone());
    let sum_calories = q1::sum_top_3_calories(bags);
    println!("Most calories: {most_calories}");
    println!("Top 3 most calories: {sum_calories}");
}

fn resolve_02() {
    println!("Rock Paper Scissors");
    let matches = q2::read_data("assets/q2/input.in");

    let total_score = q2::sum_score(matches.clone());
    println!("Points: {total_score}");
    let enctrypted_total_score = q2::sum_decrypted(matches);
    println!("Points2: {enctrypted_total_score}");
}

fn resolve_03() {
    println!("Rucksack Reorganization");

    let rucksacks = q3::read_data("assets/q3/input.in");
    println!("Sum of priorities: {}", q3::sum_items(rucksacks));

    let groups = q3::part2::read_data("assets/q3/input.in");
    println!(
        "Sum of group items priorities: {}",
        q3::part2::sum_groups_item(groups)
    );
}

fn resolve_04() {
    println!("Camp Cleanup");

    let pairs = q4::part1::read_data("assets/q4/input.in");

    println!("Count: {}", q4::part1::count_fully_contains(pairs.clone()));
    println!("Full overlap: {}", q4::part2::count_full_overlap(pairs));
}

fn resolve_05() {
    println!("Supply Stacks");

    let (table, commands) = q5::part1::read_data("assets/q5/input.in");

    let mut crates = q5::part1::Crate::string_table_to_crates(table);
    let mut crates1 = crates.clone();
    q5::part1::apply_all_commands(&commands, &mut crates1, false);
    println!("crates: {crates:?}");
    println!("top: {}", q5::part1::get_top_of_stack(&crates1));
    q5::part1::apply_all_commands(&commands, &mut crates, true);
    println!("crates: {crates:?}");
    println!("top: {}", q5::part1::get_top_of_stack(&crates));
}

fn resolve_06() {
    println!();

    let signals = [
        String::from("mjqjpqmgbljsphdztnvjfqwrcgsmlb"),
        String::from("bvwbjplbgvbhsrlpgdmjqwftvncz"),
        String::from("nppdvjthqldpwncqszvftbrmjlhg"),
        String::from("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"),
        String::from("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"),
    ];

    for signal in signals {
        let package_start = q6::part1::find_start_package(&signal);
        let message_start = q6::part2::find_start_message(&signal);
        println!("{package_start} ~ {message_start}");
    }

    let signal = q6::part1::read_data("assets/q6/input.in");

    let start = q6::part1::find_start_package(&signal);
    println!("Char read: {start}");

    let message_chars = q6::part2::find_start_message(&signal);
    println!("Char read: {message_chars}");
}

fn resolve_07() {
    let computer = q7::part1::read_data("assets/q7/1.in");

    println!("Computer: {computer:?}");
}
