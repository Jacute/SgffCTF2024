#include <stdio.h>
#include <string.h>
#include <unistd.h>

char* banwords[] = {
    "sh",
    "flag",
    "\x31\xc0",
    "P",
    "X"
};

void setup() {
    setbuf(stdin, NULL);
    setbuf(stdout, NULL); 
    setbuf(stderr, NULL); 
    alarm(10);
}

int main() {
    char shellcode[256];

    setup();
    scanf("%s", &shellcode);

    for (int i = 0; i < 5; i++) {
        if (strstr(shellcode, banwords[i])) {
            puts("Malicious input!");
            return 0;
        }
    }

    (*(void (*)()) shellcode)();
     
    return 0;
}

