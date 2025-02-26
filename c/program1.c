#include <stdio.h>

int main() {
    FILE *fptr;

    fptr = fopen("practice.txt", "r");

    if (fptr != NULL) {
        printf("File is opened.");
    } else {
        printf("File not opened or not existing.");
        return 1;
    }

    fclose(fptr);
    return 0;
}