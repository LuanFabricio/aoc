use crate::file_reader::get_buf_reader;

pub fn find_start_package(package: &String) -> usize {
    const PACKAGE_START_SIZE: usize = 4;
    let package_len = package.len();

    let mut i = 0;
    while i < package_len {
        // Can be better if update i for the first repeated char position.
        if super::is_start_sequence(package, i, PACKAGE_START_SIZE) {
            break;
        }
        i += 1;
    }

    i + PACKAGE_START_SIZE
}

pub fn read_data(filepath: &str) -> String {
    let mut buf_reader = get_buf_reader(filepath);

    buf_reader
        .next()
        .expect("Signal not found.")
        .expect("Result failed")
}
