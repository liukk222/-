#include<stdio.h>
#include<string.h>
#include<math.h>
int add(char a[])
{
    int n = strlen(a);//判断字符串的长度函数
    int i, num = 0;
    for (i = 0; i < n; i++)
    {

        if (a[i] == 'A')
            num += 10 * pow(16, n - i - 1); //pow() 函数用来求 x 的 y 次方的值。
        else if (a[i] == 'B')
            num += 11 * pow(16, n - i - 1);
        else if (a[i] == 'C')
            num += 12 * pow(16, n - i - 1);
        else if (a[i] == 'D')
            num += 13 * pow(16, n - i - 1);
        else if (a[i] == 'E')
            num += 14 * pow(16, n - i - 1);
        else if (a[i] == 'F')
            num += 15 * pow(16, n - i - 1);
        else {
            int j =0;
            for(j=48;j<=57;j++){
                if(a[i]==j)
                {
                    num +=(j-48) * pow(16, n - i - 1);
                }
            }

        }
    


    }
    return num;
}

void main()
{
   /* char b[1] ={'0'};
    printf("%d", b[0]);*/
    //让你知道0的ASCII值


    char a[10];
    printf("请输入一个16进制的数:\n");
    gets_s(a);//gets从标准输入设备读字符串函数。可以无限读取，不会判断上限，以回车结束读取
    printf("%d\n", add(a));
}
