#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define ANSI_COLOR_RED     "\x1b[31m"
#define ANSI_COLOR_GREEN   "\x1b[32m"
#define ANSI_COLOR_GRAY   "\x1b[90m"
#define ANSI_COLOR_RESET   "\x1b[0m"

#define MIN(x, y) (x) < (y) ? (x) : (y)

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

void read_lines(const char* file_path, int **output, int *lens, int *lines)
{
	FILE *file = fopen(file_path, "r");

	int safe_amount = 0;
	char buffer[256];
	int try = 0;

	int line = 0;
	while (fgets(buffer, 256, file) != NULL) {
		try++;
		char *buffer_ptr = buffer;

		int col = 0;
		do {
			int number;
			sscanf(buffer_ptr, "%d", &output[line][col]);

			buffer_ptr = move_to_next_number(buffer_ptr);
			col++;
		} while(strlen(buffer_ptr) > 0);

		lens[line] = col;
		line++;
	}

	*lines = line;
	fclose(file);
}

int count_erros(int *report, int report_len, int start, int end, int increase_factor, int error_count, int last_offset)
{
	if (start > end) return error_count;

	int diff = report[start] - report[start-last_offset];
	int diff_abs = abs(diff);

	char test_amount_diff = diff_abs < 1 || diff_abs > 3;
	char test_diff_signal = (diff < 0 && increase_factor == 1)
		|| (diff > 0 && increase_factor == -1);

	if (test_amount_diff || test_diff_signal) printf(ANSI_COLOR_RED);
	else printf(ANSI_COLOR_GREEN);
	printf("[%i] (%i, %i)(%i) => %i ", start, report[start-last_offset], report[start], increase_factor, diff);
	printf(ANSI_COLOR_RESET);
	if (test_amount_diff || test_diff_signal) {
		error_count = count_erros(report, report_len, start+1, end, increase_factor, error_count+1, 2);
	} else {
		error_count = count_erros(report, report_len, start+1, end, increase_factor, error_count, 1);
	}

	return error_count;
}

int main(void)
{
	int lines = 0;
	int lens[1000] = {0};
	int **output = malloc(sizeof(int*)*1000);
	for (int i = 0; i < 1000; i++) {
		output[i] = malloc(sizeof(int) * 30);
	}
	read_lines("input.in", output, lens, &lines);
	// read_lines("example.in", "example.out");

	int amount_safe = 0;
	for (int i = 0; i < lines; i++) {
		printf("[%i] ", i+1);
		for (int j = 0; j < lens[i]; j++) {
			printf("%i ", output[i][j]);
		}
		int increase_factor = (output[i][1] - output[i][0]) < 0 ? -1 : 1;
		printf("\n\t");
		int error_count_1 = count_erros(output[i], lens[i], 1, lens[i]-1, increase_factor, 0, 1);
		printf("\n\t");
		int error_count_2 = count_erros(output[i], lens[i], 2, lens[i]-1, -increase_factor, 0, 1);

		int min_erros = MIN(error_count_1, error_count_2);
		printf("\n\t%i, %i | %i", error_count_1, error_count_2, min_erros);
		printf("\n");

		amount_safe += min_erros <= 1;
	}
	printf("Amount safe: %i\n", amount_safe);

	for (int i = 0; i < 1000; i++) {
		free(output[i]);
	}
	free(output);

	return 0;
}
