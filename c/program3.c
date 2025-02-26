#include <stdio.h>

int main() {
    FILE *fptr;

    fptr = fopen("newFile.txt", "w");

    fputs("Hello\n", fptr);
    fputs("It's me!\n", fptr);

    fclose(fptr);
    return 0;
}