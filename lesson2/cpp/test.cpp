#include "test.h"
#include <iostream>  
int main(int argc,char *argv[])  
{  
    std::cout << "hello, world" << std::endl;
    return(0);  
}  

void Test::call() {
    printf("call from c++ language\n");
}