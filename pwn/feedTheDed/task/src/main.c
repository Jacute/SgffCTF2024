#include <stdio.h>
#include <stdbool.h>
#include <unistd.h>


void setup() {
    setvbuf(stdout, (char *)NULL, _IONBF, 0); 
    setvbuf(stderr, (char *)NULL, _IONBF, 0); 
}

int main() {
	char name[64];
    bool isFeeding = false;

    alarm(5);
    setup();

    puts("Who will feed the ded?");
    fputs("Say your name: ", stdout);
    scanf("%s", &name);

    printf("Ok. Your name is %s.\n", name);

    if (isFeeding) {
        puts("Ded is fed! :P");
        puts("SgffCTF{4m___4m_4m}");
    } else {
        puts("Ded isn't fed! :(");
    }

	return 0; 
}