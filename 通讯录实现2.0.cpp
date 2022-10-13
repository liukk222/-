#include"通讯录2.0.h"
//静态版本
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
	//读文件
	peoInfo tmp = { 0 };
	while(fread(&tmp, sizeof(peoInfo), 1, pf))
	{
		//是否需要増容
		CheckCapacity(pc); 
		pc->data[pc->sz] = tmp;
		pc->sz++;
	}
	//关闭文件
	fclose(pf);
	pf = NULL;
}
//动态版本-初始化
// 给datas申请一块连续的空间再堆区上
//capacity初始化当前的最大的容量
void InitContact(Contact* pc)
{
	

	pc->data = (peoInfo*)malloc(DEFAULT_SZ * sizeof(peoInfo));
	if (pc->data == NULL)
	{
		perror("InitContact");
		return;
	}
	pc->sz = 0;//初始化后默认为0
	pc->capacity = DEFAULT_SZ;
	//加载文件
	LoadContact(pc);
}
//销毁通讯录
void DestoryContact(Contact* pc)
{
	free(pc->data);
	pc->data = NULL;
	pc->sz = 0;
	pc->capacity=0;
}

//静态版本-增加联系人
//void AddContact(Contact* pc)
//{
//	if (pc->sz == MAX)
//	{
//		printf("通讯录已满，无法添加");
//		return;
//	}
//	//增加一个人的信息
//	printf("请输入名字:>");
//	scanf("%s", pc->data[pc->sz].name);
//	printf("请输入年龄:>");
//	scanf("%d", &(pc->data[pc->sz].age));
//	printf("请输入性别:>");
//	scanf("%s", pc->data[pc->sz].sex);
//	printf("请输入电话:>");
//	scanf("%s", pc->data[pc->sz].tele);
//	printf("请输入地址:>");
//	scanf("%s", pc->data[pc->sz].addr);
//	pc->sz++;
//	printf("增加成功\n");
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
			printf("增容成功\n");
		}
		else {
			perror("Addcontact");
			printf("增加失败\n");
			return;
		}

	}
}

//动态版本-增加联系人
void AddContact(Contact* pc)
{
	CheckCapacity(pc);
	
	//增加一个人的信息
	printf("请输入名字:>");
	scanf("%s", pc->data[pc->sz].name);
	printf("请输入年龄:>");
	scanf("%d", &(pc->data[pc->sz].age));
	printf("请输入性别:>");
	scanf("%s", pc->data[pc->sz].sex);
	printf("请输入电话:>");
	scanf("%s", pc->data[pc->sz].tele);
	printf("请输入地址:>");
	scanf("%s", pc->data[pc->sz].addr);
	pc->sz++;
	printf("增加成功\n");
}
void PrintContact(const Contact* pc)
{
	int i = 0;
	printf("%-20s\t%-5s\t%-5s\t%-12s\t%-20s\n", "名字", "年龄", "性别", "电话", "地址");
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
		printf("通讯录为空，无需删除\n");
		return;
	}
	printf("请输入要删除的人的名字");
	scanf("%s", name);
	//1.查找要删除的人
	//有、没有
int pos=FindByName(pc,name);
if(pos==-1)
{
	printf("要删除的人不存在");
	return;
}
	//2.删除
int i = 0;
for (i = pos; i < pc->sz-1; i++)
{
	pc->data[i] = pc->data[i + 1];
}
pc->sz--;
printf("删除成功\n");
}
void SearchContact(Contact* pc)
{
	char name[MAX_NAME] = { 0 };
	printf("请输入要查找的人的名字");
	scanf("%s", name);
	int i = FindByName(pc, name);
	if (i == -1)
	{
		printf("查找要的人不存在");
	}
	else
	{
		printf("%-20s\t%-5s\t%-5s\t%-12s\t%-20s\n", "名字", "年龄", "性别", "电话", "地址");
		printf("%-20s\t%-5d\t%-5s\t%-12s\t%-20s\n", pc->data[i].name, pc->data[i].age, pc->data[i].sex, pc->data[i].tele, pc->data[i].addr);
	}
}
void ModifContact(Contact* pc)
{
	char name[MAX_NAME] = { 0 };
	printf("请输入要修改的人的名字");
	scanf("%s", name);
	int i = FindByName(pc, name);
	if (i == -1)
	{
		printf("要修改的人不存在");
	}
	else
	{
		printf("请输入名字:>");
		scanf("%s", pc->data[i].name);
		printf("请输入年龄:>");
		scanf("%d", &(pc->data[i].age));
		printf("请输入性别:>");
		scanf("%s", pc->data[i].sex);
		printf("请输入电话:>");
		scanf("%s", pc->data[i].tele);
		printf("请输入地址:>");
		scanf("%s", pc->data[i].addr);
		printf("修改成功\n");
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
	printf("排序成功\n");
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