#include <stdio.h>

int main() {
    FILE *fptr;

    fptr = fopen("practice.txt", "r");
    char content[1000];

    if (fptr != NULL) {
        while(fgets(content,1000,fptr)){
            printf("%s", content);
        }
        printf("\n");
    } else {
        printf("File is not opened or not existing.\n");
        return 1;
    }

    fclose(fptr);
    return 0;
}