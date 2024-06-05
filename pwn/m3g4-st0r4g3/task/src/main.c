#include <stdio.h>
#include <unistd.h>
#include <stdbool.h>
#include <malloc.h>
#include <string.h>

#define STORAGE_SIZE 5
#define TEXT_SIZE 32

char* FLAG = "SgffCTF{b1m_b1m_b4m_b4m}";

void setup() {
    setvbuf(stdout, (char *)NULL, _IONBF, 0); 
    setvbuf(stderr, (char *)NULL, _IONBF, 0); 
}

void banner() {
    puts("           _____       _  _             _    ___       _  _         _____ ");
    puts(" _ __ ___ |___ /  __ _| || |        ___| |_ / _ \\ _ __| || |   __ _|___ / ");
    puts("| '_ ` _ \\  |_ \\ / _` | || |_ _____/ __| __| | | | '__| || |_ / _` | |_ \\ ");
    puts("| | | | | |___) | (_| |__   _|_____\\__ \\ |_| |_| | |  |__   _| (_| |___) |");
    puts("|_| |_| |_|____/ \\__, |  |_|       |___/\\__|\\___/|_|     |_|  \\__, |____/ ");
    puts("                 |___/                                        |___/       ");
    puts("\n");
}

void menu() {
    puts("=== m3g4-st0r4g3 ===");
    puts("1. PUT");
    puts("2. DELETE");
    puts("3. INFO");
    puts("4. ALLOC FLAG");
    puts("5. EXIT");
}

int main() {
    alarm(90);
    setup();

    char** storage = malloc(STORAGE_SIZE * sizeof(char*));
    int count = 0;

    banner();

    int choice;
    while (true) {
        menu();

        putchar('>');
        scanf("%d", &choice);
        getchar();

        int index;
        switch (choice) {
            case 1: // PUT
                if (count >= STORAGE_SIZE) {
                    puts("Error! Storage is full!");
                    continue;
                }

                storage[count] = malloc(TEXT_SIZE);

                fputs("Text: ", stdout);
                fgets(storage[count], TEXT_SIZE, stdin);
                int length = strlen(storage[count]);
                storage[count][length - 1] = '\0';
                count++;

                if (length == 31) while (getchar() != '\n'); // clear the buffer
                break;
            case 2: // DELETE
                fputs("Number: ", stdout);
                scanf("%d", &index);                
                getchar();
                
                if ((index < 1 || index >= STORAGE_SIZE) || index > count) {
                    puts("Error! Number out of range!");
                    continue;
                }
                index--;
                
                free(storage[index]);
                count--;
                break;
            case 3: // INFO
                puts("=== INFO===");
                for (int i = 0; i < STORAGE_SIZE; i++) {
                    printf("%d. %s\n", i + 1, storage[i]);
                }
                break;
            case 4:
                char* flagPtr = malloc(strlen(FLAG) + 1);
                strcpy(flagPtr, FLAG);
                flagPtr[strlen(FLAG)] = '\0';
                break;
            case 5:
                return 0;
            default:
                break;
        }
    }

	return 0; 
}