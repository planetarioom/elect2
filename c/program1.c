#include <stdio.h>

int main() {
    FILE *fptr;

    fptr = fopen("practice.txt", "r");

    if (fptr != NULL) {
        printf("File is opened.\n");
    } else {
        printf("File not opened or not existing.\n");
        return 1;
    }

    fclose(fptr);
    return 0;
}