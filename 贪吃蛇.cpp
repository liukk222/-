#include"贪吃蛇.h"




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
	printf("-----1.开始游戏------\n");
	printf("-----2.退出游戏------\n");
	printf("---------------------\n");
	printf("---------------------\n");
}

int main()
{
	srand((unsigned int)time(NULL));
	int input = 0;
	do {
		meat();
		printf("请选择>:");
		scanf("%d", &input);
		getchar();
		if (input == 1)
		{
			game();
		}
		else if (input == 0)
		{
			printf("退出游戏\n");
		}
		else
		{
			printf("请重新选择\n");
		}

	} while (input);

	return 0;
}
