#include"ͨѶ¼2.0.h"
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
	//����ͨѶ¼
	//peoInfo date[MAX];
	////ͨѶ¼�е�ǰ�ܹ��м���Ԫ��
	//int sz = 0;
	Contact con;//ͨѶ¼
	//��ʼ��ͨѶ¼
	InitContact(&con);
	do
	{
		menu();
		printf("��ѡ��:");
		scanf("%d", &input);
		switch (input)
		{
		case ADD:
			AddContact(&con);
			//�����˵���Ϣ
			break;
		case DEL:
			//ɾ���˵���Ϣ
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
			//������Ϣ���ļ�
			SaveContact(&con);
			//����ͨѶ¼
			DestoryContact(&con);
			printf("�˳�ͨѶ¼");
				break;
		default:
			printf("ѡ�����������ѡ��");
			break;
		}
	} while (input);
return 0;
}