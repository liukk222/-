# define _CRT_SECURE_NO_WARNINGS 1
#include<stdio.h>
#include<string>
#include<assert.h>
int my_strlen(const char* arr1)
{
	int a = 0;
	assert(arr1 != NULL);
	while (*arr1 != 0) {
		*arr1++;
		a++;

	}
	return a;


}
int main()
{
	char arr[20] = "1234567890";

	printf("%d\n", my_strlen(arr));
	return 0;

}