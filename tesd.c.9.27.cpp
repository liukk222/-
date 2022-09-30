# define _CRT_SECURE_NO_WARNINGS 1
#include<stdio.h>
#include<string>
void prints(char as[10], char bs[10])
{
	int f = 0;
	int i = 4;
	while (i) {
		
		bs[8- f]=as[8-f];
		bs[0 + f] = as[0 + f];
		f++;
		i--; printf("%s\n", bs); 
	}
	
}


void print(char as[10], char bs[10])
{
	
	char cs[10]; strcpy(cs, as);
	int f = 1;
	int i = 5;
	while (i) {
		printf("%s\n", as);
		as[4 - f] = bs[4 - f];
		as[4 + f] = bs[4 + f];
		f++;
		i--;
	}
	prints(cs, bs);
	
}

int main()
{
	
	 char a[10] = "    #    ";
	 char b[10] = "#########";
	 printf("%s\n", b);
	print(a,b);
	
	/*char ac[10] = "    #    ";
	char bc[10] = "#########";
	prints(ac, bc);*/
	return 0;
}