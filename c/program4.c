#include <stdio.h>

int main() {
    int num1, num2;

    printf("Enter the first number: ");
    scanf("%d", &num1);

    printf("Enter the second number: ");
    scanf("%d", &num2);

    printf("\nMultiplication Table for %d:\n", num1);
    for (int i = 1; i <= 10; i++) {
        printf("%d x %d = %d\n", num1, i, num1 * i);
    }

    printf("\nMultiplication Table for %d:\n", num2);
    for (int i = 1; i <= 10; i++) {
        printf("%d x %d = %d\n", num2, i, num2 * i);
    }

    return 0;
}
