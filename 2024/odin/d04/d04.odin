package main;

import "core:fmt"
import "core:strings"
import "core:os"

WordsToFind :: struct {
	words: [dynamic]string,
	bit_map: [dynamic][dynamic]bool,
	total_xmas: int,
	width: int,
	height: int,
}

xmas := "XMAS"
xmas_len := len(xmas)-1

read_input :: proc(filepath: string) -> WordsToFind{
	data, ok := os.read_entire_file(filepath, context.allocator);
	defer delete(data, context.allocator);

	it, _ := strings.clone_from_bytes(data);

	words := make([dynamic]string, 0, 0);
	i := 0
	for s in strings.split_lines_iterator(&it) {
		append(&words, s);
	}

	for line, i in words {
		fmt.printfln("[%i]: %s", i, line);
	}
	width := len(words[0]);
	height := len(words);
	bit_map := make([dynamic][dynamic]bool, height, height);
	for i in 0..<height {
		bit_map[i] = make([dynamic]bool, width, width);
		for j in 0..<width {
			bit_map[i][j] = false;
		}
	}

	return WordsToFind{words,bit_map,0,width,height};
}

test_cross :: proc(words_to_find: WordsToFind, x, y: int) -> int {
	fail := false;
	count := 0;

	if x+xmas_len < words_to_find.width && y+xmas_len < words_to_find.height {
		for c, i in xmas {
			if auto_cast words_to_find.words[y+i][x+i] != c {
				fail = true;
				break;
			}
		}
		if !fail {
			for c, i in xmas {
				words_to_find.bit_map[y+i][x+i] = true;
			}
			count += 1;
		}
	}

	fail = false;
	if x+xmas_len < words_to_find.width && y-xmas_len >= 0 {
		for c, i in xmas {
			if auto_cast words_to_find.words[y-i][x+i] != c {
				fail = true;
				break;
			}
		}
		if !fail {
			for c, i in xmas {
				words_to_find.bit_map[y-i][x+i] = true;
			}
			count += 1;
		}
	}

	fail = false;
	if x-xmas_len >= 0 && y+xmas_len < words_to_find.height {
		for c, i in xmas {
			if auto_cast words_to_find.words[y+i][x-i] != c {
				fail = true;
				break;
			}
		}
		if !fail {
			for c, i in xmas {
				words_to_find.bit_map[y+i][x-i] = true;
			}
			count += 1;
		}
	}

	fail = false;
	if x-xmas_len >= 0 && y-xmas_len >= 0{
		for c, i in xmas {
			if auto_cast words_to_find.words[y-i][x-i] != c {
				fail = true;
				break;
			}
		}
		if !fail {
			for c, i in xmas {
				words_to_find.bit_map[y-i][x-i] = true;
			}
			count += 1;
		}
	}

	return count;
}

test_vertical :: proc(words_to_find: WordsToFind, x, y: int) -> int {
	fail := false;
	count := 0;

	if y+xmas_len < words_to_find.height {
		for c, i in xmas {
			if auto_cast words_to_find.words[y+i][x] != c {
				fail = true;
				break;
			}
		}
		if !fail {
			for c, i in xmas {
				words_to_find.bit_map[y+i][x] = true;
			}
			count += 1;
		}
	}

	fail = false;
	if y-xmas_len >= 0 {
		for c, i in xmas {
			if auto_cast words_to_find.words[y-i][x] != c {
				fail = true;
				break;
			}
		}
		if !fail {
			for c, i in xmas {
				words_to_find.bit_map[y-i][x] = true;
			}
			count += 1;
		}
	}

	return count;
}

test_horizontal :: proc(words_to_find: WordsToFind, x, y: int) -> int {
	fail := false;
	count := 0;

	if x+xmas_len < words_to_find.width {

		for c, i in xmas {
			if auto_cast words_to_find.words[y][x+i] != c {
				fail = true;
				break;
			}
		}
		if !fail {
			for c, i in xmas {
				words_to_find.bit_map[y][x+i] = true;
			}
			 count += 1;
		}
	}

	fail = false;
	if x-xmas_len >= 0 {
		for c, i in xmas {
			if auto_cast words_to_find.words[y][x-i] != c {
				fail = true;
				break;
			}
		}
		if !fail {
			for c, i in xmas {
				words_to_find.bit_map[y][x-i] = true;
			}
			count += 1;
		}
	}

	return count;
}

find_words :: proc(words_to_find: WordsToFind, x, y: int, test_map: ^[dynamic][dynamic]bool) -> int {
	if x < 0 || x >= words_to_find.width do return 0;
	if y < 0 || y >= words_to_find.height do return 0;
	if test_map[y][x] do return 0;

	test_map[y][x] = true;
	count := 0;

        count += test_vertical(words_to_find, x, y);
        count += test_horizontal(words_to_find, x, y);
        count += test_cross(words_to_find, x, y);

	count += find_words(words_to_find, x+1, y, test_map);
	count += find_words(words_to_find, x+1, y+1, test_map);
	count += find_words(words_to_find, x, y+1, test_map);

	return count;
}

main :: proc() {
	// words_to_find := read_input("example.in");
	words_to_find := read_input("test.in");
	defer delete(words_to_find.words);
	defer delete(words_to_find.bit_map);

	test_map := make([dynamic][dynamic]bool, words_to_find.height);
	for line, i in words_to_find.words {
		test_map[i] = make([dynamic]bool, words_to_find.width);
		for cell, j in line {
			fmt.printf("%c ", cell);
			test_map[i][j] = false;
		}
		fmt.println();
	}
	fmt.println("========================");

	words_to_find.total_xmas = find_words(words_to_find, 0, 0, &test_map);

	for line, i in words_to_find.bit_map {
		for cell, j in line {
			if cell do fmt.printf("%c ", words_to_find.words[i][j]);
			else do fmt.printf(". ");
		}
		fmt.println();
	}

	fmt.printfln("Total xmas: %i", words_to_find.total_xmas);
}
