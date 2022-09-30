# define _CRT_SECURE_NO_WARNINGS 1
#include<stdio.h>
#include<string>
int main()
{
	char arr[100] = { 0 };
	gets_s(arr);
	int i = 0;


	while (arr[i])
	{
		int x = i;
		while (arr[x] != ' ' && arr[x] != '\0')
		{
			x++;
		}

		if (5 <= x - i)
		{
			printf("%c\n", arr[i + 1]);
			printf("%c\n", arr[i + 3]);
			printf("%c\n", arr[i + 5]);
		}
		if (arr[x] == ' ') { i = x + 1; }
		else { i = x; }
	}return 0;
}
