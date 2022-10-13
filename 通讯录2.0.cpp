#include"通讯录2.0.h"
void menu()
{
	printf("**************************\n");
	printf("****1.add  2.del**********\n");
	printf("****3.sear 4.modify*******\n");
	printf("****5.sort 6.print********\n");
	printf("********0.exit************\n");
}
enum optinon
{
	EXIT,
	ADD,
	DEL,
	SEARCH,
	MODIFY,
	SORT,
	PRINT
};


int main()
{
	int input = 0;
	//创建通讯录
	//peoInfo date[MAX];
	////通讯录中当前总共有几个元素
	//int sz = 0;
	Contact con;//通讯录
	//初始化通讯录
	InitContact(&con);
	do
	{
		menu();
		printf("请选择:");
		scanf("%d", &input);
		switch (input)
		{
		case ADD:
			AddContact(&con);
			//增加人的信息
			break;
		case DEL:
			//删除人的信息
			DelContact(&con);
			break;
		case SEARCH:
			SearchContact(&con);
			break;
		case MODIFY:
			ModifContact(&con);
			break;
		case SORT:
			SortContact(&con);
			break;
		case PRINT:
			PrintContact(&con);
			break;
		case EXIT:
			//保存信息到文件
			SaveContact(&con);
			//销毁通讯录
			DestoryContact(&con);
			printf("退出通讯录");
				break;
		default:
			printf("选择错误，请重新选择");
			break;
		}
	} while (input);
return 0;
}