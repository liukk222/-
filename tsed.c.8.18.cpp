# define _CRT_SECURE_NO_WARNINGS 1
#include<stdio.h>
#include<string>
int main()
{
	short s = 0;
	int a = 5;
	printf("%d\n",sizeof (s = 5 + a));
	printf("%d\n", s);
	return 0;
}