package main;

import "core:fmt"
import "core:strconv"
import "core:strings"
import "core:os"

DFAStep :: struct {
	name: string,
	key: string,
	next: [dynamic]^DFAStep,
}

// TODO: Write a better version
main :: proc() {
	fmt.printf("Hello, world!\n");

	// mul
	step_p2 := DFAStep {"par2", ")", {}};
	step_comma := DFAStep {"comma", ",", {&step_p2}};
	step_num := DFAStep {"num", "0123456789", {&step_p2, &step_comma}};
	append(&step_num.next, &step_num);
	append(&step_comma.next, &step_num);
	step_p1 := DFAStep {"par1", "(", {&step_num}};
	step_l := DFAStep {"l", "l", {&step_p1}};
	step_u := DFAStep {"u", "u", {&step_l}};
	step_m := DFAStep {"m", "m", {&step_u}};
	step_initial := DFAStep {"initial", "", {&step_m}}
	append(&step_p2.next, &step_initial);

	current_step := step_initial;
	total := 0;

	data, ok := os.read_entire_file("test.in", context.allocator);
	defer delete(data, context.allocator);

	it := string(data)
	for x in strings.split_lines_iterator(&it) {
		num_acc := strings.builder_make()
		n1 := 0
		for c, i in x {
			updated_step := false;
			// fmt.printf("\nTesting: %c\n", c);
			for step in current_step.next {
				if strings.contains_rune(step.key, c) {
					current_step = step^;
					fmt.printf("%c", c);
					updated_step = true;
					break;
				}
			}
			if !updated_step {
				current_step = step_initial;
				strings.builder_reset(&num_acc);
				fmt.printf("\n");
			} else {
				switch current_step.name {
					case step_num.name: strings.write_rune(&num_acc, c);
					case step_comma.name: {
						n1 = strconv.atoi(strings.to_string(num_acc));
						strings.builder_reset(&num_acc);
					}
					case step_p2.name: {
						total += n1 * strconv.atoi(strings.to_string(num_acc));
						strings.builder_reset(&num_acc);
						current_step = step_initial;
					}
				}
			};
		}
	}
	fmt.printf("\n");

	fmt.printf("Total: %i\n", total);
}
