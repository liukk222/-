#include"̰����.h"




void game()
{
	char board[ROW][COL] = { 0 };
	 InitBoard(board, ROW, COL);
	dispalyBoard(board, ROW, COL);
	gamea(board, ROW, COL);
	dispalyBoard(board, ROW, COL);
	InitBoard(board, ROW, COL);
	gameb(board, ROW, COL);
	
}



	void meat()
	{
		
	printf("---------------------\n");
	printf("-----1.��ʼ��Ϸ------\n");
	printf("-----2.�˳���Ϸ------\n");
	printf("---------------------\n");
	printf("---------------------\n");
}

int main()
{
	srand((unsigned int)time(NULL));
	int input = 0;
	do {
		meat();
		printf("��ѡ��>:");
		scanf("%d", &input);
		getchar();
		if (input == 1)
		{
			game();
		}
		else if (input == 0)
		{
			printf("�˳���Ϸ\n");
		}
		else
		{
			printf("������ѡ��\n");
		}

	} while (input);

	return 0;
}
