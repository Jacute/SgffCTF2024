#include <stdio.h>


void gift() {
    __asm__ volatile (
        "pop %%r12\n\t"
        "pop %%r13\n\t"
        "pop %%r14\n\t"
        "pop %%r15\n\t"
        "ret\n\t"
        :
        :
        :
    );
}

void setup() {
    setbuf(stdin, NULL);
    setbuf(stdout, NULL); 
    setbuf(stderr, NULL); 
}

void vuln() {
    char buf[48];
    puts("My name is Vito Scaletta. I was born in Sicily in 1925. That little guy's me. I'm standing there with my parents and my sister Francesca in front...");
    fgets(buf, 256, stdin);
}

int main() {
    alarm(10);
    setup();
	vuln();
	return 0; 
}