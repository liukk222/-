#pragma once
# define _CRT_SECURE_NO_WARNINGS 1
#include<stdio.h>
#include<string.h>
#include<stdlib.h>
#include<time.h>
#define ROW 10
#define COL 10
//��ʼ������
void InitBoard(char board[ROW][COL], int row, int col);
//��ӡ����
void dispalyBoard(char board[ROW][COL], int row, int col);
void gamea(char board[ROW][COL], int row, int col);
void gameb(char board[ROW][COL], int row, int col);