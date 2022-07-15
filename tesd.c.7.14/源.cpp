#define _CRT_SECURE_NO_WARNINGS 1
#include<stdio.h>
#include<string>
int main()
{
	/*char a = 'q';
	char* b = &a;
	*b = 'c';*/

	int a = 10;
	int* p = &a;
	*p = 20;
	printf("%d\n", a);
	printf("%p\n", p);
	printf("%p\n", &a);
	printf("%d\n", sizeof(p));
	return 0;
}