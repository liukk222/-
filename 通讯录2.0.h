#pragma once
# define _CRT_SECURE_NO_WARNINGS 1
#include<stdio.h>
#include<string>
#include<stdlib.h>
#define MAX_NAME 20
#define MAX_SEX 10
#define MAX_TELE 12
#define MAX_ADDR 30
#define MAX 1000
#define DEFAULT_SZ 3
#define INC_SZ 2
typedef struct peoInfo
{

	char name[MAX_NAME];
	char sex[MAX_SEX];
	int age;
	char tele[MAX_TELE];
	char addr[MAX_ADDR];
}peoInfo;
//ͨѶ¼-��̬�汾
//typedef struct Contact
//{
//	peoInfo data[MAX];//�����ӽ������˵���Ϣ
//	int sz;//��¼�ǵ�ǰͨѶ¼����Ч��Ϣ����
//} Contact;
//��̬�汾
typedef struct Contact
{
	peoInfo* data;//ָ��̬����Ŀռ䣬����������ϵ�˵���Ϣ
	int sz;//��¼�ǵ�ǰͨѶ¼����Ч��Ϣ����
	int capacity;//��¼��ǰͨѶ¼���������
} Contact;
//��ʼ��ͨѶ¼
void InitContact(Contact* pc);
//������ϵ��
void AddContact(Contact* pc);
//��ӡͨѶ¼
void PrintContact(const Contact* pc);
//ɾ����ϵ��
void DelContact(Contact* pc);
//������ϵ��
void SearchContact(Contact* pc);
//�޸���ϵ��
void ModifContact(Contact* pc);
//������ϵ��
void SortContact(Contact* pc);
//����ͨѶ¼
void DestoryContact(Contact* pc);
//������Ϣ��ͨѶ¼
void SaveContact(Contact* pc);
//�����ļ����ݵ�ͨѶ¼
void LoadContact(Contact* pc);
//�Ƿ���Ҫ����
void CheckCapacity(Contact* pc);