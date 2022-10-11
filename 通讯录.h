#pragma once
# define _CRT_SECURE_NO_WARNINGS 1
#include<stdio.h>
#include<string>
#define MAX_NAME 20
#define MAX_SEX 10
#define MAX_TELE 12
#define MAX_ADDR 30
#define MAX 1000
typedef struct peoInfo
{

	char name[MAX_NAME];
	char sex[MAX_SEX];
	int age;
	char tele[MAX_TELE];
	char addr[MAX_ADDR];
}peoInfo;
//通讯录
typedef struct Contact
{
	peoInfo data[MAX];//存放添加进来的人的信息
	int sz;//记录是当前通讯录中有效信息个数
} Contact;

//初始化通讯录
void InitContact(Contact* pc);
//增加联系人
void AddContact(Contact* pc);
//打印通讯录
void PrintContact(const Contact* pc);
//删除联系人
void DelContact(Contact* pc);
//查找联系人
void SearchContact(Contact* pc);
//修改联系人
void ModifContact(Contact* pc);
//排序联系人
void SortContact(Contact* pc);