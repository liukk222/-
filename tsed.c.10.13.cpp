# define _CRT_SECURE_NO_WARNINGS 1
#include<stdio.h>
#include<string>
#include<Windows.h>
#define MAX(x,y) (x)>(y)?(x):(y)
int main()
{
	int a = 0;
	int b = 0;
	scanf("%d %d", &a, &b);
	int c = MAX(a, b);
	printf("%d", c);


	FILE* pf = fopen("tesd.txt", "w");
	fputs("abcef", pf);
	Sleep(3000);
	printf("Ë¢ÐÂ");
		fflush(pf);
		//fclose(pf);
		pf = NULL;

return 0;
}