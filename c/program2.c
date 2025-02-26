#include <stdio.h>

int main() {
    FILE *fptr;

    fptr = fopen("practice.txt", "r");
    char content[1000];

    if (fptr != NULL) {
        while(fgets(content,1000,fptr)){
            printf("%s", content);
        }
    } else {
        printf("File is not opened or not existing.");
        return 1;
    }

    fclose(fptr);
    return 0;
}