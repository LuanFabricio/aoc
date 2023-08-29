use std::{
    fs::File,
    io::{BufRead, BufReader, Lines},
};

pub fn get_buf_reader(filepath: &str) -> Lines<BufReader<File>> {
    let file = File::open(filepath).expect("File not found.");
    BufReader::new(file).lines()
}
