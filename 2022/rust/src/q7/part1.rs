use std::collections::HashMap;

use crate::file_reader::get_buf_reader;

#[derive(Debug)]
pub struct Command {
    command: String,
    arguments: String,
}

impl Command {
    pub fn from(line: String) -> Self {
        let mut split = line.split(" ");
        split.next();

        let command = split.next().expect("Command not found").to_string();
        let arguments = split.next().unwrap_or("").to_string();
        Self { command, arguments }
    }
}

#[derive(Debug)]
pub enum FolderItem {
    File { file: File },
    Folder { folder: Folder },
}

#[derive(Debug)]
pub struct Folder {
    name: String,
    files: Vec<FolderItem>,
    total_size: u32,
}

impl Folder {
    pub fn new(name: String) -> Self {
        Self {
            name,
            files: vec![],
            total_size: 0,
        }
    }

    pub fn add_file(&mut self, file: File) {
        self.total_size += file.size;
        self.files.push(FolderItem::File { file });
    }

    pub fn add_folder(&mut self, folder: Folder) {
        if self.contains(&folder.name) {
            self.total_size += folder.total_size;
            self.files.push(FolderItem::Folder { folder });
        }
    }

    pub fn contains(&self, path: &String) -> bool {
        for file in self.files.iter() {
            match file {
                FolderItem::Folder { folder } => {
                    println!("Contains? {path} == {}", folder.name);
                    if folder.name.eq(path) {
                        return true;
                    }
                }
                _ => {}
            }
        }

        false
    }

    pub fn update_folder(&mut self, folder: Folder) {
        println!("Folder: {folder:?}");
        let folder_name = folder.name.clone();
        for file in self.files.iter_mut() {
            match file {
                FolderItem::Folder {
                    folder: current_folder,
                } => {
                    if current_folder.name == folder_name {
                        self.total_size -= current_folder.total_size;
                        self.total_size += current_folder.total_size;
                        // let new_folder = Folder::new(folder.name);
                        *file = FolderItem::Folder { folder };
                        return;
                    }
                    continue;
                }
                _ => continue,
            };
        }
    }
}

#[derive(Debug)]
pub struct File {
    name: String,
    size: u32,
}

impl File {
    pub fn new(name: String) -> Self {
        Self { name, size: 0 }
    }

    pub fn from(line: String) -> Self {
        let mut split = line.split(" ");

        let size = split
            .next()
            .expect("File size not found.")
            .parse::<u32>()
            .expect("Parse erro!");
        let name = split.next().expect("File name not found.").to_string();

        Self { name, size }
    }
}

enum LineType {
    Command { command: Command },
    Folder { folder: Folder },
    File { file: File },
}

#[derive(Debug)]
pub struct State {
    current_folder: String,
    command_queue: Vec<Command>,
    folders: HashMap<String, Folder>,
}

impl State {
    pub fn new() -> Self {
        Self {
            current_folder: String::from("~"),
            command_queue: vec![],
            folders: HashMap::new(),
        }
    }

    pub fn use_command(&mut self, command: Command) {
        match command.command.as_str() {
            "cd" => {
                if command.arguments.starts_with("/") {
                    self.current_folder = command.arguments;
                } else if command.arguments.starts_with("..") {
                    let mut count = self.current_folder.split("/").count();
                    let split = self.current_folder.split("/");

                    self.current_folder = split.fold(String::new(), |a, b| {
                        count -= 1;
                        if count > 1 {
                            return a + b + "/";
                        }
                        a
                    });
                } else {
                    self.current_folder += format!("{}/", command.arguments).as_str();
                }
            }
            _ => {}
        }
    }

    pub fn add_folder(&mut self, folder: Folder) {
        let folder_path = format!("{}{}/", self.current_folder, folder.name);

        println!("current_folder: {}", self.current_folder);
        println!("folders  {:?}", self.folders);
        if self.live_in_folder(&folder_path) {
            self.update_in_folder(&folder_path, folder);
        } else if let Some(current_folder) = self.folders.get_mut(&self.current_folder) {
            current_folder.add_folder(folder);
        } else {
            self.folders.insert(folder_path, folder);
        }
    }

    fn live_in_folder(&self, folder_path: &String) -> bool {
        for (_, folder) in self.folders.iter() {
            if folder.contains(folder_path) {
                return true;
            }
        }
        false
    }

    fn update_in_folder(&mut self, folder_path: &String, folder: Folder) {
        for (_, current_folder) in self.folders.iter_mut() {
            if current_folder.contains(folder_path) {
                current_folder.update_folder(folder);
                return;
            }
        }
    }

    fn update_other_folder_size(&mut self, folder_path: &String, folder: &Folder) {
        println!("folder: {folder_path}");
        for (_, current_folder) in self.folders.iter_mut() {
            println!(
                "\ton: {} | {}",
                current_folder.name,
                current_folder.files.len()
            );
            for file in current_folder.files.iter_mut() {
                match file {
                    FolderItem::File { file } => {}
                    FolderItem::Folder { folder } => {
                        println!("\t\tfile: {}", folder.name);
                        if folder.name == *folder_path {
                            folder.total_size = folder.total_size;
                            println!("\t\tUpdate!");
                        }
                    }
                }
            }
        }
    }

    pub fn add_file(&mut self, file: File) {
        if let Some(folder) = self.folders.get_mut(&self.current_folder.to_string()) {
            folder.add_file(file);
        }
    }

    fn parse_line(line: String) -> LineType {
        println!("parse_line: {line}");
        if line.starts_with("$") {
            let command = Command::from(line);
            LineType::Command { command }
        } else if line.starts_with("dir") {
            let mut split = line.split(" ");
            split.next();
            let name = split.next().expect("File name not found");
            let folder = Folder::new(name.to_string());
            LineType::Folder { folder }
        } else {
            let file = File::from(line);
            LineType::File { file }
        }
    }
}

pub fn read_data(filepath: &str) -> State {
    let buf_reader = get_buf_reader(filepath);

    let mut state = State::new();
    for line in buf_reader {
        if let Ok(line) = line {
            let line_type = State::parse_line(line);
            match line_type {
                LineType::Command { command } => state.use_command(command),
                LineType::Folder { folder } => state.add_folder(folder),
                LineType::File { file } => state.add_file(file),
            }
        }
    }

    state
}

pub fn calc_total_size(folders: &mut HashMap<String, Folder>) {
    let mut folders: Vec<(String, Folder)> = vec![];
}
