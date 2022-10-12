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
	printf("程序知道你不服，再给你一次机会\n");
	printf("1.我要逆天改命  2.我承认这就是我在女神心中的地位\n");
	int f = 0;
	scanf("%d", &f);
	getchar();
	if (f == 1) { game(); }
	else if (f == 2)
	{
		printf("不要卑微的做一条舔狗，站起来，我们去舔一群做一条狼狗\n");
		printf("不要卑微的做一条舔狗，站起来，我们去舔一群做一条狼狗\n");
		printf("不要卑微的做一条舔狗，站起来，我们去舔一群做一条狼狗\n");
		Sleep(3000);
		system("C:\\vs.c\\让程序猜猜在女神心中是什么地位2.0\\as.gif");
	}
	else{ printf("输入错误，请重新输入");}
		return;
	
}





void game() {
	for (int i = 0; i < 4; i++) { printf("程序分析中...\n"); Sleep(1000); }
	int x = 0;
	x = rand() % 10 + 1;
	switch (x) {
	case 1:printf("哦，你连女神家的一条狗都不如\n"); sss(); break;
	case 2:printf("哦，原来你是备胎\n"); sss(); break;
	case 3:printf("哦，很抱歉女神想不起来你是谁了\n"); sss();  break;
	case 4:printf("哦，原来你是女神的提款机\n"); sss(); break;
	case 5:printf("哦，在女神心中你只是一条单身狗\n"); sss(); break;
	case 6:printf("哦，在女神心中你是一个好人\n"); sss();  break;
	case 7:printf("哦，对不起女神有点讨厌你每天来烦她\n"); sss();  break;
	case 8:printf("哦，对不起女神的心中还有另一个他\n"); sss();  break;
	case 9:printf("哦，女神从来没有爱过你\n"); sss(); break;
	case 10:printf("哦, 其实不用再想了，你自己心里也知道你就是一个小丑\n"); sss();  break;
	default:printf("输入错误，请重新输入"); break;
	}


}

int main() {
	int a = 0;
	do
	{
		mciSendString(L"open asd.mp3", 0, 0, 0);
		mciSendString(L"play asd.mp3", 0, 0, 0);
		srand((unsigned int)time(0));
		printf("让程序猜猜你在女神心中是什么地位\n");
		printf("     1.让我看看   2.我暂时还不想知道  3.我的心中没有女神\n");
		scanf("%d", &a);
		getchar();
		switch (a) {
		case 1:game(); break;
		case 2:printf("骚年不要再幻想啦面对现实吧"); game(); break;
		case 3:("没错女人只会影响我们拔剑的速度"); break;
		default:printf("输入错误，请重新输入"); break;
		}
	} while (a);
	return 0;
}