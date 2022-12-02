# define _CRT_SECURE_NO_WARNINGS 1
#include<stdio.h>
#include<string.h>
int main() {
	int a;
	for (int i = 0; i < 1000; i++) {
		if (i < 10) {
			if ((i==5) && (i % 3 == 0)) {
				printf("%d\n", i);
			}
		}
		 if (i < 100) {
			a = i % 10;
			if ((a==5) && (i % 3 == 0)) {
				printf("%d\n", i);
			}

		}
		else {
			a = i % 100;
			if ((a==5) && (i % 3 == 0)) {
				printf("%d\n", i);
			}
		}


		
	}return 0;
}