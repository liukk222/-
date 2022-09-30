# define _CRT_SECURE_NO_WARNINGS 1
#include<stdio.h>
#include<string>

int main()
{
 int a = 0x11223344;
	//int* pa = &a;
	 char* pc =(char*) & a;
	//printf("%d", pa);
	printf("%d", pc);
	return 0;
}