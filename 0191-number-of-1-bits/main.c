#include "stdio.h"
#include "stdint.h"

int hammingWeight(uint32_t n) {
    int cnt = 0;
    while(1) {
        if(n == 0) {
            break;
        }
        
        if (n%2 == 1) {
            cnt ++;
        }
        
        n = n/2;
    }
    return cnt;
}

int main() {
    return 0;
}

