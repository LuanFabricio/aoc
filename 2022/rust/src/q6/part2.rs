pub fn find_start_message(package: &String) -> usize {
    const MESSAGE_START_SIZE: usize = 14;
    let package_len = package.len();

    let mut i = 0;
    while i < package_len {
        // Can be better if update i for the first repeated char position.
        if super::is_start_sequence(package, i, MESSAGE_START_SIZE) {
            break;
        }
        i += 1;
    }

    i + MESSAGE_START_SIZE
}
