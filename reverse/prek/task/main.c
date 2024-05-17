#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <malloc.h>


const char* flag = "SgffCTF{REDACTED}";


char* encrypt(const char* text) {
    char* buf = malloc(strlen(text) + 5 * strlen(text));
    int step;

    int bufIndex = 0;
    for (int i = 0; i < strlen(text); i++) {
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
    //char* cipherText = encrypt(flag);
}
