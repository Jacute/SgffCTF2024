#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <malloc.h>


const char* flag = "SgffCTF{sr4nd_1s_bull5h1t}";


char* encrypt(const char* text) {
    int textLen = strlen(text);
    char* buf = malloc(textLen + 5 * textLen);
    int step;

    int bufIndex = 0;
    for (int i = 0; i < textLen; i++) {
        step = rand() % 5;
        
        int j;
        for (j = bufIndex; j < bufIndex + step; j++) {
            buf[j] = rand() % 128;
        }
        buf[j] = text[i];
        bufIndex = j + 1;
    }

    return buf;
}

void writeFile(const char* filename, const char* text) {
    FILE *f = fopen(filename, "w");
    if (f == NULL) {
        perror("Unable to create file");
        return;
    }

    fprintf(f, "%s\n", text);
    fclose(f);
}

void main() {
    srand(1337);
    
    char* cipherText = encrypt(flag);
    writeFile("output.txt", cipherText);
    free(cipherText);
}
