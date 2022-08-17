#include"game.h"

void game()
{
	char mine[ROWS][COLS] = { 0 };
	char show[ROWS][COLS] = { 0 };
	InitBoard(mine, ROWS, COLS, '0');
	InitBoard(show, ROWS, COLS, '*');
	DisplayBoard(mine, ROW, COL);
	DisplayBoard(show, ROW, COL);
	SetMine(mine, ROW, COL);
	DisplayBoard(mine, ROW, COL);
	FindMine(mine, show, ROW, COL);
}



void menu()
{
	printf("********************\n");
	printf("***1.paly  2.esc****\n");
	printf("********************\n");
}


void tset()
{
	int input = 0;

	do
	{
		menu();
		printf("«Î—°‘Ò");
		scanf("%d", &input);
		switch (input)
		{
		case 1:game();
			break;
		case 0:printf("esc");
			break;
		default:printf("esc°£2");
			break;
		}
	} while (input);
}


int main()
{
	srand((unsigned int)time(0));
	tset();
	return 0;
}