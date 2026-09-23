#include <malloc.h>
int* getPtrOnStack()
{
    // n 分配在栈上，函数返回即被回收
    int n;
    int* pi = &n;
    return pi;
} 
int* mallocInt()
{
    // 动态分配的内存分配在堆上，需要自行释放
    return (int*)malloc(sizeof(int));
}
int main()
{
    // pi 为悬空指针
    int* pi = getPtrOnStack();
    int* pi2 = mallocInt();
    // 申请的内存没有释放，应该先 free(pi2)
    // free(pi2);
    pi2 = 0;
    // pi2 置零后，失去了对未释放内存的控制，因为地址已经找不回了
    // 短时间内一次过运行的程序内存泄漏问题不大，到程序退出都会释放；
    // 但对于需要持续运行的程序，内存泄漏会造成严重后果。
}
// gcc -Wall -Wextra mall.c -o mall