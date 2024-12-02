#include <stdio.h>
#include <stdlib.h>

const int LIST_LEN = 1000;

void swap(int* a, int* b)
{
	int tmp = *a;
	*a = *b;
	*b = tmp;
}

int partition(int *arr, int low, int high)
{
	int pivot = arr[high];

	int i = low;
	for (int j = low; j <= high-1; j++) {
		if (arr[j] < pivot) {
			swap(&arr[i], &arr[j]);
			i++;
		}
	}
	swap(&arr[i], &arr[high]);
	return i;
}

void quicksort(int *arr, int low, int high)
{
	if (low >= high || low < 0) return;

	int p = partition(arr, low, high);

	quicksort(arr, low, p-1);
	quicksort(arr, p+1, high);
}

void read_input(int* left_list, int *right_list)
{
	FILE *file = fopen("puzzle.in", "r");

	for (int i = 0; i < LIST_LEN; i++) {
		fscanf(file, "%i %i", &left_list[i], &right_list[i]);
	}

	fclose(file);
}

int repeated(const int *list, const int list_len, int number)
{
	int count = 0;
	for (int i = 0; i < list_len; i++) {
		if (number == list[i]) count ++;
	}
	return count;
}

int main(void)
{

	int *left_list = (int*)calloc(LIST_LEN, sizeof(int));
	int *right_list = (int*)calloc(LIST_LEN, sizeof(int));

	read_input(left_list, right_list);

	quicksort(left_list, 0, LIST_LEN-1);
	quicksort(right_list, 0, LIST_LEN-1);

	int total_diff = 0;
	for (int i = 0; i < LIST_LEN; i++) {
		int diff = abs(left_list[i] - abs(right_list[i]));
		total_diff += diff;
		printf("%i ", diff);
	}
	printf("= %i\n", total_diff);

	char cache_use[1000000] = {0};
	int cache[1000000] = {0};
	int score = 0;
	for (int i = 0; i < LIST_LEN; i++) {
		int number = left_list[i];
		if (cache_use[left_list[i]]) {
			score += cache_use[number];
		} else {
			int res = repeated(right_list, LIST_LEN, number) * number;

			cache_use[number] = 1;
			cache[number] = res;
			score += res;
		}
	}

	printf("Score: %i\n", score);

	free(left_list);
	free(right_list);

	return 0;
}
