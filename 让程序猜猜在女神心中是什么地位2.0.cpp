# define _CRT_SECURE_NO_WARNINGS 1
#include<stdio.h>
#include<string>
#include<windows.h>
#include<time.h>
#include<mmsystem.h>
#pragma comment(lib,"winmm.lib")
void game();
void sss();

void sss()
{
	printf("����֪���㲻�����ٸ���һ�λ���\n");
	printf("1.��Ҫ�������  2.�ҳ������������Ů�����еĵ�λ\n");
	int f = 0;
	scanf("%d", &f);
	getchar();
	if (f == 1) { game(); }
	else if (f == 2)
	{
		printf("��Ҫ��΢����һ���򹷣�վ����������ȥ��һȺ��һ���ǹ�\n");
		printf("��Ҫ��΢����һ���򹷣�վ����������ȥ��һȺ��һ���ǹ�\n");
		printf("��Ҫ��΢����һ���򹷣�վ����������ȥ��һȺ��һ���ǹ�\n");
		Sleep(3000);
		system("C:\\vs.c\\�ó���²���Ů��������ʲô��λ2.0\\as.gif");
	}
	else{ printf("�����������������");}
		return;
	
}





void game() {
	for (int i = 0; i < 4; i++) { printf("���������...\n"); Sleep(1000); }
	int x = 0;
	x = rand() % 10 + 1;
	switch (x) {
	case 1:printf("Ŷ������Ů��ҵ�һ����������\n"); sss(); break;
	case 2:printf("Ŷ��ԭ�����Ǳ�̥\n"); sss(); break;
	case 3:printf("Ŷ���ܱ�ǸŮ���벻��������˭��\n"); sss();  break;
	case 4:printf("Ŷ��ԭ������Ů�������\n"); sss(); break;
	case 5:printf("Ŷ����Ů��������ֻ��һ������\n"); sss(); break;
	case 6:printf("Ŷ����Ů����������һ������\n"); sss();  break;
	case 7:printf("Ŷ���Բ���Ů���е�������ÿ��������\n"); sss();  break;
	case 8:printf("Ŷ���Բ���Ů������л�����һ����\n"); sss();  break;
	case 9:printf("Ŷ��Ů�����û�а�����\n"); sss(); break;
	case 10:printf("Ŷ, ��ʵ���������ˣ����Լ�����Ҳ֪�������һ��С��\n"); sss();  break;
	default:printf("�����������������"); break;
	}


}

int main() {
	int a = 0;
	do
	{
		mciSendString(L"open asd.mp3", 0, 0, 0);
		mciSendString(L"play asd.mp3", 0, 0, 0);
		srand((unsigned int)time(0));
		printf("�ó���²�����Ů��������ʲô��λ\n");
		printf("     1.���ҿ���   2.����ʱ������֪��  3.�ҵ�����û��Ů��\n");
		scanf("%d", &a);
		getchar();
		switch (a) {
		case 1:game(); break;
		case 2:printf("ɧ�겻Ҫ�ٻ����������ʵ��"); game(); break;
		case 3:("û��Ů��ֻ��Ӱ�����ǰν����ٶ�"); break;
		default:printf("�����������������"); break;
		}
	} while (a);
	return 0;
}