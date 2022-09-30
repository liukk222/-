# define _CRT_SECURE_NO_WARNINGS 1
#include<stdio.h>
#include<string>

int main()
{
	const int n = 10;
	const int* p = &n;
	int a = 20;
	p = &a;
	printf("%d\n", n);
	return 0;

}

//int main()
//{
//	int n = 0;
//	scanf("%d", &n);
//	int i = 0;
//	int ret = 1;
//	int sum = 0;
//	int j = 0;
//	for (j = 1; j <= n; j++)
//	{
//		ret = 1;
//		for (i = 1; i <= j; i++) {
//			
//			ret *= i;
//
//	}
//		sum += ret;
//	}
//	printf("%d\n ", sum);
//	printf("\\""n");
//	return 0;
//}