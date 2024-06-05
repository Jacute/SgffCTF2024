#include <stdio.h>
#include <unistd.h>

void setup() {
    setvbuf(stdout, (char *)NULL, _IONBF, 0); 
    setvbuf(stderr, (char *)NULL, _IONBF, 0); 
}

void becomePneuma() {
    FILE *file = fopen("/flag", "r");

    if (file == NULL) {
        perror("Error opening file");
        return;
    }

    fseek(file, 0, SEEK_END);
    long filesize = ftell(file);
    fseek(file, 0, SEEK_SET);

    char *buffer = malloc(filesize + 1);

    size_t bytesRead = fread(buffer, 1, filesize, file);
    if (bytesRead != filesize) {
        perror("Error reading file");
        free(buffer);
        fclose(file);
        return;
    }
    buffer[filesize] = '\0';

    puts(buffer);
}

void oneFootNailedDown() {
    char buf[64];

    fputs("Child, release the light\nLight: ", stdout);
    gets(buf);
}

int main() {
    alarm(5);
    
    setup();
	oneFootNailedDown();

	return 0; 
}