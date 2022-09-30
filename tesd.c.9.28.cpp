# define _CRT_SECURE_NO_WARNINGS 1
#include<stdio.h>
#include<string>
int main()
{
	int b = 0;
	int a = 20;
	while (a)
	{
		if (b < 2)
		{
			a--;
			b++;
			if (b == 2) { b++; }
		}
		else
		{
			a--;
			b = b + 2;
		}

		//1.2 3.4 5.6 7.8 9.10  11.12
 	}//    2   3   4   5    6    7   8   9   10
	/*printf("%d\n", b);
	while (a)
	{
		a--;
		b = b + 2;

	}*/
	printf("%d\n",a);
	printf("%d", b);
	return 0;
}