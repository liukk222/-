
#include"Ã∞≥‘…ﬂ.h"

//void gameif(char board[ROW][COL], int row, int col)
//{
//
//	int i = 0;
//	int j = 0;
//	for (i = 0; i <= row; i++)
//	{
//		for (j = 0; j <= col; j++)
//		{
//			if (board[i][j] == 'o') {}
//			else { printf("”Œœ∑ §¿˚"); break; }
//
//		}
//	}




void InitBoard3(char board[ROW][COL], int row, int col)
{
	int x = 0; int y = 0;
	int q = 5;
	while (q) {
		x = rand() % 10;
		y = rand() % 10;
		if (y > 0) { board[x][y] = 'o'; q--; }

	}

}


void InitBoard2(char board[ROW][COL], int row, int col)//≥ı ºªØ∆Â≈Ã
{
	int i = 0;
	int j = 0;
	int x = 0;
	int y = 0;
	for (i = 0; i <= row; i++)
	{
		for (j = 0; j <= col; j++)
		{
			if (board[i][j] != ' ' && board[i][j] != '*' && board[i][j] != 'c')//o

			{
				for (x = 0; x <= row; x++) {
					for (y = 0; y <= col; y++)
					{
						if (board[x][y] == 'c' || board[x][y] == '*')
						{
							board[x][y] = ' ';
						}


					}
				}
			
				}
			else if (board[0][0] != 'o' &&
				board[0][1] != 'o' &&
				board[0][2] != 'o' &&
				board[0][3] != 'o' &&
				board[0][4] != 'o' &&
				board[0][5] != 'o' &&
				board[0][6] != 'o' &&
				board[0][7] != 'o' &&
				board[0][8] != 'o' &&
				board[0][9] != 'o' &&
				board[1][0] != 'o' &&
				board[1][1] != 'o' &&
				board[1][2] != 'o' &&
				board[1][3] != 'o' &&
				board[1][4] != 'o' &&
				board[1][5] != 'o' &&
				board[1][6] != 'o' &&
				board[1][7] != 'o' &&
				board[1][8] != 'o' &&
				board[1][9] != 'o' &&
				board[2][0] != 'o' &&
				board[2][1] != 'o' &&
				board[2][2] != 'o' &&
				board[2][3] != 'o' &&
				board[2][4] != 'o' &&
				board[2][5] != 'o' &&
				board[2][6] != 'o' &&
				board[2][7] != 'o' &&
				board[2][8] != 'o' &&
				board[2][9] != 'o' &&
				board[3][0] != 'o' &&
				board[3][1] != 'o' &&
				board[3][2] != 'o' &&
				board[3][3] != 'o' &&
				board[3][4] != 'o' &&
				board[3][5] != 'o' &&
				board[3][6] != 'o' &&
				board[3][7] != 'o' &&
				board[3][8] != 'o' &&
				board[3][9] != 'o' &&
				board[4][0] != 'o' &&
				board[4][1] != 'o' &&
				board[4][2] != 'o' &&
				board[4][3] != 'o' &&
				board[4][4] != 'o' &&
				board[4][5] != 'o' &&
				board[4][6] != 'o' &&
				board[4][7] != 'o' &&
				board[4][8] != 'o' &&
				board[4][9] != 'o' &&
				board[5][0] != 'o' &&
				board[5][1] != 'o' &&
				board[5][2] != 'o' &&
				board[5][3] != 'o' &&
				board[5][4] != 'o' &&
				board[5][5] != 'o' &&
				board[5][6] != 'o' &&
				board[5][7] != 'o' &&
				board[5][8] != 'o' &&
				board[5][9] != 'o' &&
				board[6][0] != 'o' &&
				board[6][1] != 'o' &&
				board[6][2] != 'o' &&
				board[6][3] != 'o' &&
				board[6][4] != 'o' &&
				board[6][5] != 'o' &&
				board[6][6] != 'o' &&
				board[6][7] != 'o' &&
				board[6][8] != 'o' &&
				board[6][9] != 'o' &&
				board[7][0] != 'o' &&
				board[7][1] != 'o' &&
				board[7][2] != 'o' &&
				board[7][3] != 'o' &&
				board[7][4] != 'o' &&
				board[7][5] != 'o' &&
				board[7][6] != 'o' &&
				board[7][7] != 'o' &&
				board[7][8] != 'o' &&
				board[7][9] != 'o' &&
				board[8][0] != 'o' &&
				board[8][1] != 'o' &&
				board[8][2] != 'o' &&
				board[8][3] != 'o' &&
				board[8][4] != 'o' &&
				board[8][5] != 'o' &&
				board[8][6] != 'o' &&
				board[8][7] != 'o' &&
				board[8][8] != 'o' &&
				board[8][9] != 'o' &&
				board[9][0] != 'o' &&
				board[9][1] != 'o' &&
				board[9][2] != 'o' &&
				board[9][3] != 'o' &&
				board[9][4] != 'o' &&
				board[9][5] != 'o' &&
				board[9][6] != 'o' &&
				board[9][7] != 'o' &&
				board[9][8] != 'o' &&
				board[9][9] != 'o') {
				printf(" ”Œœ∑ §¿˚ "); break
				;}
			
			
			
			
			
		}
	}
}


	void InitBoard(char board[ROW][COL], int row, int col)//≥ı ºªØ∆Â≈Ã;
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
	int itu = 1;
	int * p =&itu;

	int y = 0;
	int z = 2;

	int f = 1;
	int i = 0;
	int j = 1;
	int r = 2;//1…œ2œ¬
	
	
	while (itu) {
		

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
			 InitBoard2(board, ROW, COL); if (f == 0) { f++; }r = 2;
		}
		else if (i == 1) {
			if (j == 1)
			{
				z = z + 1; board[z][y-1] = 'c';
				board[z - 1][y-1] = '*';
				board[z - 1][y-2] = 'c'; i--; system("cls"); dispalyBoard(board, ROW, COL);
				InitBoard2(board, ROW, COL); y--;
			}//”“
			 if(j==2){
				z = z + 1; board[z-1][y] = 'c';
				board[z-1][y-1] = '*';
				board[z][y-1] = 'c';  i--; system("cls"); dispalyBoard(board, ROW, COL);
				InitBoard2(board, ROW, COL); y--;
			
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
				InitBoard2(board, ROW, COL); if (f == 0) { f++; }r = 1;
			}
			else if (i == 1) {
				if (j == 1) {
					z = z - 1;
					board[z][y - 1] = 'c';
					board[z + 1][y - 1] = '*';
					board[z + 1][y - 2] = 'c'; i--; system("cls"); dispalyBoard(board, ROW, COL);
					InitBoard2(board, ROW, COL);  z = z + 2; y--;
				}
				if (j == 2) {
					z = z - 1;
					board[z][y - 1] = 'c';
					board[z + 1][y - 1] = '*';
					board[z + 1][y] = 'c'; i--; system("cls"); dispalyBoard(board, ROW, COL);
					InitBoard2(board, ROW, COL);  z = z + 2; y--;
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
					InitBoard2(board, ROW, COL);f--; z--; y--;
				}
				if (r == 1) {
					y = y + 1;
					board[z][y - 1] = 'c';
					board[z-1][y - 1] = '*';
					board[z-1][y] = 'c';
					system("cls");	dispalyBoard(board, ROW, COL);
					InitBoard2(board, ROW, COL);  f--; z--;

				}
			}
			else if (f == 0)
			{
				y = y + 1;
				board[z][y - 2] = 'c';
				board[z][y - 1] = '*';
				board[z][y] = 'c';
				system("cls");	dispalyBoard(board, ROW, COL);
				InitBoard2(board, ROW, COL); if (i == 0) { i++; }j = 1;
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
					InitBoard2(board, ROW, COL);  f--; z--; y = y + 2;
				}
				if (r == 1) {
					y = y - 1;
					board[z ][y + 1] = 'c';
					board[z - 1][y + 1] = '*';
					board[z - 1][y] = 'c';
					system("cls"); dispalyBoard(board, ROW, COL);
					InitBoard2(board, ROW, COL); f--; z--; y = y + 2;
				}

			}
			else if (f == 0)
			{
				y = y-1;
				board[z][y - 2] = 'c';
				board[z][y - 1] = '*';
				board[z][y] = 'c';
				system("cls"); dispalyBoard(board, ROW, COL);
				InitBoard2(board, ROW, COL);  if (i == 0) { i++; }j = 2;
			}
		}
		


	}
		
		

	}
