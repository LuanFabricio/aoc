#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define ANSI_COLOR_RED     "\x1b[31m"
#define ANSI_COLOR_GREEN   "\x1b[32m"
#define ANSI_COLOR_GRAY   "\x1b[90m"
#define ANSI_COLOR_RESET   "\x1b[0m"

void print_number(int number, int erros_amount, char current_error)
{
	if (erros_amount == 0) printf(ANSI_COLOR_GREEN "%i " ANSI_COLOR_RESET, number);
	else if (erros_amount == 1 && current_error)
		printf(ANSI_COLOR_GRAY "%i " ANSI_COLOR_RESET, number);
	else if (erros_amount == 1) printf(ANSI_COLOR_GREEN "%i " ANSI_COLOR_RESET, number);
	else printf(ANSI_COLOR_RED "%i " ANSI_COLOR_RESET, number);
}

char* move_to_next_number(char* ptr)
{
	while (ptr[0] != ' ' && ptr[0] != '\n') ptr++;
	return ++ptr;
}

void read_lines(const char* file_path, const char* output_path)
{
	FILE *file = fopen(file_path, "r");

	FILE *output = fopen(output_path, "w");

	int safe_amount = 0;
	char buffer[256];
	int try = 0;
	while (fgets(buffer, 256, file) != NULL) {
		try++;
		char *buffer_ptr = buffer;
		int last = 0;
		int erros_amount = 0;
		char is_decrease = -1;
		char on_error = 0;

		int i = 0;
		sscanf(buffer_ptr, "%d", &last);
		buffer_ptr = move_to_next_number(buffer_ptr);
		do {
			int number;
			sscanf(buffer_ptr, "%d", &number);

			int diff = last - number;
			int abs_diff = abs(diff);

			if (i == 0) is_decrease = diff < 0 ? 1 : 0;

			char not_safe_test_diff = (abs_diff < 1)
				|| (abs_diff > 3);
			char not_safe_test_udpate = (diff < 0 && !is_decrease)
				|| (diff > 0 && is_decrease);
			char not_safe_tests = not_safe_test_diff || not_safe_test_udpate;

			if (not_safe_tests) on_error = 1;
			else last = number;

			buffer_ptr = move_to_next_number(buffer_ptr);
			i++;
		} while(strlen(buffer_ptr) > 0);

		if (!on_error) {
			safe_amount++;
		}
	}

	printf("Amount: %i\n", safe_amount);

	fclose(file);
	fclose(output);
}

int main(void)
{
	read_lines("input.in", "input.out");
	// read_lines("example.in", "example.out");
	return 0;
}
