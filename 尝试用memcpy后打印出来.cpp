# define _CRT_SECURE_NO_WARNINGS 1
#include<stdio.h>
#include<string>

int main()
{
	int a[10] = { 1,2,3,4,5,6,7,8,9,10 };
	int b[10] = { 0 };
	memcpy(b, a, 20);
	int n = 0;
	for (n = 0; n <=9; n++)
	{
		printf("%d", b[n]);
}
	return 0;
}