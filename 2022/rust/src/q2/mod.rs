use crate::file_reader::get_buf_reader;

#[derive(Debug, Clone)]
pub struct Match {
    p1: String,
    p2: String,
}

impl Match {
    fn get_score(&self) -> u32 {
        let mut total = get_item_point(&self.p2);

        total += match (self.p1.as_str(), self.p2.as_str()) {
            ("A", "X") | ("B", "Y") | ("C", "Z") => 3,
            ("A", "Y") | ("B", "Z") | ("C", "X") => 6,
            (_, _) => 0,
        };

        total
    }

    fn get_encrypted_score(&self) -> u32 {
        let mut score = match self.p2.as_str() {
            "Y" => 3,
            "Z" => 6,
            _ => 0,
        };

        // Getting play based on expected result.
        let p2_play = match (self.p1.as_str(), self.p2.as_str()) {
            ("A", "X") => "Z",
            ("A", "Y") => "X",
            ("A", "Z") => "Y",
            ("B", "X") => "X",
            ("B", "Y") => "Y",
            ("B", "Z") => "Z",
            ("C", "X") => "Y",
            ("C", "Y") => "Z",
            ("C", "Z") => "X",
            (_, _) => "",
        };

        score += get_item_point(p2_play);

        score
    }
}

fn get_item_point(item: &str) -> u32 {
    match item {
        "A" | "X" => 1,
        "B" | "Y" => 2,
        "C" | "Z" => 3,
        _ => 0,
    }
}

pub fn read_data(filepath: &str) -> Vec<Match> {
    let lines = get_buf_reader(filepath);

    let mut matches = Vec::<Match>::new();
    for line in lines {
        if let Ok(line) = line {
            let mut split = line.split(" ");

            let p1 = split.next().expect("Player 1 not found.").to_string();
            let p2 = split.next().expect("Player 1 not found.").to_string();
            matches.push(Match { p1, p2 });
        }
    }

    matches
}

pub fn sum_score(matches: Vec<Match>) -> u32 {
    let mut total_score = 0;

    for m in matches {
        total_score += m.get_score();
    }

    total_score
}

pub fn sum_decrypted(matches: Vec<Match>) -> u32 {
    let mut total_score = 0;

    for m in matches {
        total_score += m.get_encrypted_score();
    }

    total_score
}
