#include<stdio.h>
#include<string.h>
#include<math.h>
int add(char a[])
{
    int n = strlen(a);//�ж��ַ����ĳ��Ⱥ���
    int i, num = 0;
    for (i = 0; i < n; i++)
    {

        if (a[i] == 'A')
            num += 10 * pow(16, n - i - 1); //pow() ���������� x �� y �η���ֵ��
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
    //����֪��0��ASCIIֵ


    char a[10];
    printf("������һ��16���Ƶ���:\n");
    gets_s(a);//gets�ӱ�׼�����豸���ַ����������������޶�ȡ�������ж����ޣ��Իس�������ȡ
    printf("%d\n", add(a));
}
