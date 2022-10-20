# define _CRT_SECURE_NO_WARNINGS 1
#include <stdio.h>
#include <stdlib.h>
//int main()
//{
//	for(int i=0;i<3;i++)
//	{
//		;
//	}
//
//}
// 

//int main()
//{
//    float a[4][5], sum1, sum2;
//    int i, j;
//    for (i = 0; i < 3; i++)
//        for (j = 0; j < 4; j++)
//            scanf("%f", &a[i][j]);
//    for (i = 0; i < 3; i++)
//    {
//        sum1 = 0;
//        for (j = 0; j < 4; j++)
//            sum1 += a[i][j];
//        a[i][4] = sum1 / 4;
//    }
//    for (j = 0; j < 5; j++)
//    {
//        sum2 = 0;
//        for (i = 0; i < 3; i++)
//            sum2 += a[i][j];
//        a[3][j] = sum2 / 3;
//    }
//    for (i = 0; i < 4; i++)
//    {
//        for (j = 0; j < 5; j++)
//            printf("%6.2f", a[i][j]);
//        printf("\n");
//    }
//}
struct Student
{
	int num;
	float score;
	struct Student* next;
};
int main()
{
	struct Student a, b, c, * head, * p;
	a.num = 10101; a.score = 89.5;
	b.num = 10103; b.score = 90;
	c.num = 10107; c.score = 85;
	head = &a;
	a.next = &b;
	b.next = &c;
	c.next = NULL;
	p = head;
	do
	{
		printf("num=%ld   score=%5.1f\n", p->num, p->score);
		p = p->next;
	} while (p != NULL);
	return 0;
}
