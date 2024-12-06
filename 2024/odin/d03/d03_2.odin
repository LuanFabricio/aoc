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

State :: struct {
	step: DFAStep,
	num_acc: strings.Builder,
	str_acc: strings.Builder,
	n1: int,
	total: int,
	active: bool,
}

DFAState :: struct {
	initial: DFAStep,
	m: DFAStep, u: DFAStep, l: DFAStep,
	d: DFAStep, o: DFAStep, n: DFAStep, qm: DFAStep, t: DFAStep,
	p1: DFAStep, p2: DFAStep,
	comma: DFAStep,
	num: DFAStep,
}

dfa_transition :: proc(state: ^State, dfa_states: DFAState, c: rune) {
	updated_step := false;
	for step in state.step.next {
		if strings.contains_rune(step.key, c) {
			state.step = step^;
			fmt.printf("%c", c);
			updated_step = true;
			break;
		}
	}

	if !updated_step {
		state.step = dfa_states.initial;
		strings.builder_reset(&state.num_acc);
		strings.builder_reset(&state.str_acc);
		fmt.printf("\n");
	} else {
		strings.write_rune(&state.str_acc, c);
		switch state.step.name {
			case dfa_states.num.name: strings.write_rune(&state.num_acc, c);
			case dfa_states.comma.name: {
				state.n1 = strconv.atoi(strings.to_string(state.num_acc));
				strings.builder_reset(&state.num_acc);
			}
			case dfa_states.p2.name: {
				command := strings.to_string(state.str_acc);
				fmt.printfln("Command: %s(%b)", command, state.active);
				fmt.printfln("Total: %i", state.total);
				if strings.compare(command, "don't()") == 0{
					state.active = false;
				} else if strings.compare(command, "do()") == 0 {
					state.active = true;
				} else if state.active {
					n2 := strconv.atoi(strings.to_string(state.num_acc))
					state.total += state.n1 * n2;
				}
				strings.builder_reset(&state.num_acc);
				strings.builder_reset(&state.str_acc);
				state.step = dfa_states.initial;
			}
		}
	};
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
	step_p1 := DFAStep {"par1", "(", {&step_num, &step_p2}};
	step_l := DFAStep {"l", "l", {&step_p1}};
	step_u := DFAStep {"u", "u", {&step_l}};
	step_m := DFAStep {"m", "m", {&step_u}};

	// do(n't)
	step_t := DFAStep{"t", "t", {&step_p1}};
	step_qm := DFAStep{"\'", "\'", {&step_t}};
	step_n := DFAStep{"n", "n", {&step_qm}};
	step_o := DFAStep{"o", "o", {&step_p1, &step_n}};
	step_d := DFAStep{"d", "d", {&step_o}};

	step_initial := DFAStep {"initial", "", {&step_m, &step_d}}
	append(&step_p2.next, &step_initial);

	dfa_state := DFAState{
		step_initial,
		step_m, step_u, step_l,
		step_d, step_o, step_n, step_qm, step_t,
		step_p1, step_p2,
		step_comma,
		step_num,
	};

	data, ok := os.read_entire_file("test.in", context.allocator);
	defer delete(data, context.allocator);

	total := 0;
	it := string(data);
	for x in strings.split_lines_iterator(&it) {
		state := State{
		       dfa_state.initial,
		       strings.builder_make(),
		       strings.builder_make(),
		       0, 0, true
		};
		for c in x {
			dfa_transition(&state, dfa_state, c);
		}
		total += state.total;
		fmt.printf("Total: %i\n", state.total);
	}
	fmt.printf("\n");

	fmt.printf("Total: %i\n", total);
}
