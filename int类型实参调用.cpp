# define _CRT_SECURE_NO_WARNINGS 1
#include<stdio.h>
#include<string.h>

//int main()
//{
//	char arr[20][20] = {0};
//	for (int i = 0; i <20; i++)
//	{
//		for (int j = 0; j <20; j++) {
//			arr[i][j] = '*';
//			printf("%c", arr[i][j]);
//		}
//		printf("\n");
//	}
//	int x = 0;
//	int y = 0;
//	int r = 0;
//	int e = 0;
//	int s = 0;
//	printf("请输入马儿的坐标");
//	scanf("%d %d",&x,&y);
//	
//	if (r == x + 1 && e == y + 1 || r == x + 1 && e == y || r == x + 1 && e == y - 1 || r == x && e == y + 1 || r == x && e == y || r == x &&
//		y == y - 1 )
//	{
//		printf("%d",s);
//	}
//	printf("请输入B点的坐标");
//	scanf("%d %d", &r, &e);
//	if (r == x + 1 && e == y + 1 || r == x + 1 && e == y || r == x + 1 && e == y - 1 || r == x && e == y + 1 || r == x && e == y || r == x &&
//		y == y - 1 || r == x - 1 && e == y + 1 || r == x - 1 && e == y || r == x - 1 && e == y - 1)
//	{
//		printf("%d", s);
//	}
//	if(arr[x + 2][y + 1]=='*')
//	{
//		s++;
//	}
//	if( arr[x + 1][y + 2]=='*')
//	{
//		s++;
//	}
//if(arr[x -1][ y + 2]=='*')
//	{
//		s++;
//	}
//	if (arr[x - 2][y + 1] == '*')
//	{
//		s++;
//	}
//	if (arr[x - 2][y - 1] == '*')
//	{
//		s++;
//	}
//	if (arr[x - 1][y - 2] == '*')
//	{
//		s++;
//	}
//	if (arr[x +1][y - 2] == '*')
//	{
//		s++;
//	}
//
//if (arr[x + 2][y - 1] == '*')
//	{
//		s++;
//	}
//s = s - 2;
//printf("%d", s);
//return 0;
//}
void if_s( int* i)
{
	int e = 0;
	e = *i;
	int a =8;
	while (e % 10 != 0)
	{
		a--;
		e=e / 10;
	}
	for (int j = 0; j < a; j++)
	{
		printf(" ");
	}
	printf("%d",*i);
	printf(" ");
}


int main()
{
	int x,y,z;
	scanf("%d", &x);
	if_s(&x);
	//printf("%8d %8d %8d", x, y, z);
	/*if_s(x);
	if_s(y);
	if_s(z);*/
	return 0;
}