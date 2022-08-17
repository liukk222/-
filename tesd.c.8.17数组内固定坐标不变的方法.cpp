# define _CRT_SECURE_NO_WARNINGS 1
#include<stdio.h>
#include<string>
#include<time.h>
#define a 3
#define b 3

//int main()
//{
//	int a = 0;
//	int b = 0;
//	printf("a=\n");
//	scanf("%d", &a);
//	while(a)
//	{
//		a=a& (a - 1);
//		b++;
//	
//	}printf("%d", b);
//	return 0;
//}

int main()
{
	srand((unsigned int)time(0));
	int h = 0;
	int x = 0;
	int y = 0;
	char s[a][b] = {'0'};
	for (x = 0; x < a; x++) {
		
		for (y = 0; y < b; y++) {
			s[x][y] = { '*' };
			printf(" %c ", s[x][y]);
		}printf("\n");
	}
	scanf("%d", &h);
	while (h) {
		int x = rand() % 3;
		int y = rand() % 3; if ((x!=1||y!=1) && s[x][y] == '*') {
			s[x][y] = { '0' };
			h--;
			
		}
	}//(x!=1||y!=1)数组内固定坐标不变的方法
	printf("\n");
	for (x = 0; x < a; x++) {

		for (y = 0; y < b; y++) {
			printf(" %c ", s[x][y]);
		}printf("\n");
	}
	if (s[0][0] == '0') { printf("s[0][0]\n"); }
	 if (s[1][0] == '0') { printf("s[1][0]\n"); }
	 if (s[2][0] == '0') { printf("s[2][0]\n"); }
	 if (s[0][1] =='0') { printf("s[0][1]\n"); }
	 if (s[0][2] =='0') { printf("s[0][2]\n"); }
	 if (s[1][2] == '0') { printf("s[1][2]\n"); }
	 if (s[2][1] == '0') { printf("s[2][1]\n"); }
	 if(s[2][2]=='0') { printf("s[2][2]\n"); }
	return 0;
}