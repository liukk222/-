
#include"Ã∞≥‘…ﬂ.h"

void InitBoard(char board[ROW][COL], int row, int col)//≥ı ºªØ∆Â≈Ã
{
	int i = 0;
	int j = 0;
	for (i = 0; i <= row; i++)
	{
		for (j = 0; j <= col; j++)
		{
			board[i][j] = ' ';
		}
	}
}
void dispalyBoard(char board[ROW][COL], int row, int col)//¥Ú”°∆Â≈Ã
{
	int j = 0;
	int i = 0;
	for (i = 0; i < row; i++)
	{

		for (j = 0; j < col; j++) {
			printf(" %c ", board[i][j]);
			if (j < col - 1) { printf("|"); }
		}printf("\n");



		if (i < row - 1) {
			for (j = 0; j < col; j++) {
				printf("---");
				if (j < col - 1) { printf("|"); }
			}printf("\n");
		}
	}
	printf("\n");
}

void gamea(char board[ROW][COL], int row, int col)
{
	int x = 0;
	int y = 0;
	int z = 2;
	board[z - 2][y] = 'c';
	board[z - 1][y] = '*';
		board[z][y] = 'c';
		
	}

void gameb(char board[ROW][COL], int row, int col)
{
	int y = 0;
	int z = 2;

	int f = 1;
	int i = 0;
	int j = 1;
	int r = 2;//1…œ2œ¬
	
	
	while (1) {
		

		char vs[] = "o";
		char bs[] = "s";
		char bw[] = "w";
		char bd[] = "d";
		char ba[] = "a";
		scanf("%s", &vs);

		if (strcmp(vs, bs) == 0) {
		if(i==0)	{z = z + 1; board[z][y] = 'c';
			 board[z-1][y] = '*';
			 board[z-2][y] = 'c'; system("cls"); dispalyBoard(board, ROW, COL);
			 InitBoard(board, ROW, COL); if (f == 0) { f++; }r = 2;
		}
		else if (i == 1) {
			if (j == 1)
			{
				z = z + 1; board[z][y-1] = 'c';
				board[z - 1][y-1] = '*';
				board[z - 1][y-2] = 'c'; i--; system("cls"); dispalyBoard(board, ROW, COL);
				InitBoard(board, ROW, COL); y--;
			}//”“
			 if(j==2){
				z = z + 1; board[z-1][y] = 'c';
				board[z-1][y-1] = '*';
				board[z][y-1] = 'c';  i--; system("cls"); dispalyBoard(board, ROW, COL);
				InitBoard(board, ROW, COL); y--;
			
			}//◊Û
		}
		}
		if (strcmp(vs, bw) == 0)
		{
			if (i == 0) {
				z = z - 1;
				board[z][y] = 'c';
				board[z - 1][y] = '*';
				board[z - 2][y] = 'c'; system("cls"); dispalyBoard(board, ROW, COL);
				InitBoard(board, ROW, COL); if (f == 0) { f++; }r = 1;
			}
			else if (i == 1) {
				if (j == 1) {
					z = z - 1;
					board[z][y-1] = 'c';
					board[z + 1][y-1] = '*';
					board[z + 1][y - 2] = 'c'; i--; system("cls"); dispalyBoard(board, ROW, COL);
					InitBoard(board, ROW, COL); z = z + 2; y--;
				}
				if (j == 2) {
					z = z - 1;
					board[z][y - 1] = 'c';
					board[z + 1][y - 1] = '*';
					board[z + 1][y] = 'c'; i--; system("cls"); dispalyBoard(board, ROW, COL);
					InitBoard(board, ROW, COL); z = z + 2; y--;
				}

				
			}
		
	}
		
		if (strcmp(vs, bd) == 0) {
			if (f == 1)
			{
				if (r == 2)
				{
					y = y + 1;
					board[z - 2][y - 1] = 'c';
					board[z-1][y - 1] = '*';
					board[z-1][y] = 'c';
					system("cls"); dispalyBoard(board, ROW, COL);
					InitBoard(board, ROW, COL); f--; z--;
				}
				if (r == 1) {
					y = y + 1;
					board[z][y - 1] = 'c';
					board[z-1][y - 1] = '*';
					board[z-1][y] = 'c';
					system("cls");	dispalyBoard(board, ROW, COL);
					InitBoard(board, ROW, COL); f--; z--;

				}
			}
			else if (f == 0)
			{
				y = y + 1;
				board[z][y - 2] = 'c';
				board[z][y - 1] = '*';
				board[z][y] = 'c';
				system("cls");	dispalyBoard(board, ROW, COL);
				InitBoard(board, ROW, COL); if (i == 0) { i++; }j = 1;
			}
		}
		if (strcmp(vs, ba) == 0) {
			if (f == 1)
			{
				if (r == 2)
				{
					y = y - 1;
					board[z - 2][y + 1] = 'c';
					board[z - 1][y + 1] = '*';
					board[z - 1][y] = 'c';
					system("cls");	dispalyBoard(board, ROW, COL);
					InitBoard(board, ROW, COL); f--; z--; y = y + 2;
				}
				if (r == 1) {
					y = y - 1;
					board[z ][y + 1] = 'c';
					board[z - 1][y + 1] = '*';
					board[z - 1][y] = 'c';
					system("cls"); dispalyBoard(board, ROW, COL);
					InitBoard(board, ROW, COL); f--; z--; y = y + 2;
				}

			}
			else if (f == 0)
			{
				y = y-1;
				board[z][y - 2] = 'c';
				board[z][y - 1] = '*';
				board[z][y] = 'c';
				system("cls"); dispalyBoard(board, ROW, COL);
				InitBoard(board, ROW, COL); if (i == 0) { i++; }j = 2;
			}
		}
		


	}
		
		

	}
