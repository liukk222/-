# define _CRT_SECURE_NO_WARNINGS 1
#include<stdio.h>
#include<string>
#include<stdlib.h>
struct sut
{
	char name[20];
	int age;

};

int sort_ae(const void* as, const void* bs)
{
	
	return ((struct sut*)as)->age-((struct sut*)bs)->age;

}


 void tesd()
 {
	 struct sut s[3] = { { "a",3},{"b",2},{"c",1}};
	 //struct sut s2 = { 1 };
	 int sz = sizeof(s) / sizeof(s[0]);
	 qsort(s, sz, sizeof(s[0]), sort_ae);
	 for (int i = 0; i < 3; i++)
	 {
		 printf("%d", s[i].age);
	 }

 }


int main()
{
	 
	/*struct sut
	{
		int age;

	};*/
 tesd();
	
	
return 0;
}


//int arrs(const void* e1, const void* e2)
//{
//
//	return*(int*)e1 - *(int*)e2;
//}
//
//
//int main()
//{
//	int arr[] = { 9,8,7,6,5,4,3,2,1 };
//	int sz = sizeof(arr) / sizeof(arr[0]);
//	qsort(arr, sz, sizeof(arr[0]), arrs);
//	for (int i = 0; i < sz; i++)
//	{
//		printf("%d", arr[i]);
//	}
//
//	return 0;
//}