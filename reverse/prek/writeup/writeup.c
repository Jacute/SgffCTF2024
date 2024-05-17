#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <malloc.h>


char* decrypt(char* cipherText) {
    char* buf = malloc(strlen(cipherText) + 1);
    int step;

    int bufIndex = 0;
    for (int i = 0; i < strlen(cipherText); i++) {
        step = rand() % 5;
        for (int j = 0; j < step; j++) rand();
        bufIndex += step;
        buf[i] = cipherText[bufIndex];
        bufIndex++;
    }

    buf[strlen(cipherText)] = '\0';

    return buf;
}


char* readFile(const char* filename) {
    FILE *f = fopen(filename, "r");
    if (f == NULL) {
        perror("Unable to read file");
        return;
    }

    fseek(f, 0, SEEK_END);
    int fileLen = ftell(f);

    fseek(f, 0, SEEK_SET);

    char* buf = malloc(fileLen + 1);

    fgets(buf, fileLen, f);

    buf[fileLen] = '\x00';

    fclose(f);

    return buf;
}


int main() {
    srand(1337);
    char* cipherText = readFile("output.txt");
    if (cipherText == NULL) {
        return 1; // Выходим с ошибкой, если чтение файла не удалось
    }
    char* text = decrypt(cipherText);
    printf("%s\n", text);

    free(cipherText);
    free(text);

    return 0;
}
