#include"ͨѶ¼2.0.h"
//��̬�汾
//void InitContact(Contact* pc)
// {
//
//	pc->sz = 0;
//	memset(pc->data, 0, sizeof(pc->data));
//}
// 
void LoadContact(Contact* pc)
{
	FILE* pf = fopen("contact.txt", "r");
	if(pf==NULL)
	{
		perror("fopen");
		return;
	}
	//���ļ�
	peoInfo tmp = { 0 };
	while(fread(&tmp, sizeof(peoInfo), 1, pf))
	{
		//�Ƿ���Ҫ����
		CheckCapacity(pc); 
		pc->data[pc->sz] = tmp;
		pc->sz++;
	}
	//�ر��ļ�
	fclose(pf);
	pf = NULL;
}
//��̬�汾-��ʼ��
// ��datas����һ�������Ŀռ��ٶ�����
//capacity��ʼ����ǰ����������
void InitContact(Contact* pc)
{
	

	pc->data = (peoInfo*)malloc(DEFAULT_SZ * sizeof(peoInfo));
	if (pc->data == NULL)
	{
		perror("InitContact");
		return;
	}
	pc->sz = 0;//��ʼ����Ĭ��Ϊ0
	pc->capacity = DEFAULT_SZ;
	//�����ļ�
	LoadContact(pc);
}
//����ͨѶ¼
void DestoryContact(Contact* pc)
{
	free(pc->data);
	pc->data = NULL;
	pc->sz = 0;
	pc->capacity=0;
}

//��̬�汾-������ϵ��
//void AddContact(Contact* pc)
//{
//	if (pc->sz == MAX)
//	{
//		printf("ͨѶ¼�������޷����");
//		return;
//	}
//	//����һ���˵���Ϣ
//	printf("����������:>");
//	scanf("%s", pc->data[pc->sz].name);
//	printf("����������:>");
//	scanf("%d", &(pc->data[pc->sz].age));
//	printf("�������Ա�:>");
//	scanf("%s", pc->data[pc->sz].sex);
//	printf("������绰:>");
//	scanf("%s", pc->data[pc->sz].tele);
//	printf("�������ַ:>");
//	scanf("%s", pc->data[pc->sz].addr);
//	pc->sz++;
//	printf("���ӳɹ�\n");
//}
void CheckCapacity(Contact* pc)
{
	if (pc->sz == pc->capacity)
	{
		peoInfo* ptr = (peoInfo*)realloc(pc->data, (pc->capacity + INC_SZ) * sizeof(peoInfo));
		if (ptr != NULL)
		{
			pc->data = ptr;
			pc->capacity + INC_SZ;
			printf("���ݳɹ�\n");
		}
		else {
			perror("Addcontact");
			printf("����ʧ��\n");
			return;
		}

	}
}

//��̬�汾-������ϵ��
void AddContact(Contact* pc)
{
	CheckCapacity(pc);
	
	//����һ���˵���Ϣ
	printf("����������:>");
	scanf("%s", pc->data[pc->sz].name);
	printf("����������:>");
	scanf("%d", &(pc->data[pc->sz].age));
	printf("�������Ա�:>");
	scanf("%s", pc->data[pc->sz].sex);
	printf("������绰:>");
	scanf("%s", pc->data[pc->sz].tele);
	printf("�������ַ:>");
	scanf("%s", pc->data[pc->sz].addr);
	pc->sz++;
	printf("���ӳɹ�\n");
}
void PrintContact(const Contact* pc)
{
	int i = 0;
	printf("%-20s\t%-5s\t%-5s\t%-12s\t%-20s\n", "����", "����", "�Ա�", "�绰", "��ַ");
	for (i = 0; i < pc->sz; i++)
	{
		printf("%-20s\t%-5d\t%-5s\t%-12s\t%-20s\n", pc->data[i].name, pc->data[i].age, pc->data[i].sex, pc->data[i].tele, pc->data[i].addr);
	}
}
static int FindByName(Contact* pc,char name[])
{
	int i= 0;
	for(i=0;i<pc->sz;i++)
	{
	if(strcmp(pc->data[i].name,name)==0)
	{ 
		return i;
	}
	}
	return -1;
}
void DelContact(Contact* pc)
{
	char name[MAX_NAME] = { 0 };
	if (pc->sz == 0)
	{
		printf("ͨѶ¼Ϊ�գ�����ɾ��\n");
		return;
	}
	printf("������Ҫɾ�����˵�����");
	scanf("%s", name);
	//1.����Ҫɾ������
	//�С�û��
int pos=FindByName(pc,name);
if(pos==-1)
{
	printf("Ҫɾ�����˲�����");
	return;
}
	//2.ɾ��
int i = 0;
for (i = pos; i < pc->sz-1; i++)
{
	pc->data[i] = pc->data[i + 1];
}
pc->sz--;
printf("ɾ���ɹ�\n");
}
void SearchContact(Contact* pc)
{
	char name[MAX_NAME] = { 0 };
	printf("������Ҫ���ҵ��˵�����");
	scanf("%s", name);
	int i = FindByName(pc, name);
	if (i == -1)
	{
		printf("����Ҫ���˲�����");
	}
	else
	{
		printf("%-20s\t%-5s\t%-5s\t%-12s\t%-20s\n", "����", "����", "�Ա�", "�绰", "��ַ");
		printf("%-20s\t%-5d\t%-5s\t%-12s\t%-20s\n", pc->data[i].name, pc->data[i].age, pc->data[i].sex, pc->data[i].tele, pc->data[i].addr);
	}
}
void ModifContact(Contact* pc)
{
	char name[MAX_NAME] = { 0 };
	printf("������Ҫ�޸ĵ��˵�����");
	scanf("%s", name);
	int i = FindByName(pc, name);
	if (i == -1)
	{
		printf("Ҫ�޸ĵ��˲�����");
	}
	else
	{
		printf("����������:>");
		scanf("%s", pc->data[i].name);
		printf("����������:>");
		scanf("%d", &(pc->data[i].age));
		printf("�������Ա�:>");
		scanf("%s", pc->data[i].sex);
		printf("������绰:>");
		scanf("%s", pc->data[i].tele);
		printf("�������ַ:>");
		scanf("%s", pc->data[i].addr);
		printf("�޸ĳɹ�\n");
	}
}
void SortContact(Contact* pc)
{
	for (int i = 0; i < pc->sz; i++)
	{
		for (int j = 0; j < pc->sz-1; j++)
		{
			if(pc->data[j].age>pc->data[j+1].age)
			{
				pc->data[pc->sz] = pc->data[j+1];
			pc->data[j + 1] = pc->data[j];
			pc->data[j]= pc->data[pc->sz];

			}
		}
	}
	printf("����ɹ�\n");
}

void SaveContact(Contact* pc)
{
	FILE* pf = fopen("contact.txt", "w");
	if(pf==NULL)
	{
		perror("SaveContact");
		return;
	}
	int i = 0;
	for (i = 0; i < pc->sz; i++)
	{
		fwrite(pc->data + i, sizeof(peoInfo), 1, pf);
	}

	fclose(pf);
	pf = NULL;
}