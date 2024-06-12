BITS 64

section .data
    flag db "SgffCTF{REDACTED}", 0
    filename db "encrypted.txt", 0

section .text
    global _start

_start:
    push rbp
    mov rbp, rsp
    sub rsp, 72
    mov qword [rbp-72], 0

    mov rdi, flag
    lea rsi, [rbp-64]
    call strcpy

    lea rdi, [rbp-64]
    call encrypt

    mov rdi, filename
    lea rsi, [rbp-64]
    call write_file

    call _exit

strcpy:
    push rbp
    mov rbp, rsp
    sub rsp, 8
    xor rbx, rbx
    mov qword [rbp-8], 0

    _loop_strcpy:
        mov rcx, qword [rbp-8]
        mov bl, byte [rdi+rcx]
        cmp rbx, 0
        je _exit_loop_strcpy
        mov [rsi+rcx], bl
        add qword [rbp-8], 1
        jmp _loop_strcpy
    _exit_loop_strcpy:
        leave
        ret

encrypt:
    push rbp
    mov rbp, rsp
    sub rsp, 8
    xor rbx, rbx
    mov qword [rbp-8], 0

    _loop_encrypt:
        mov rcx, qword [rbp-8]
        mov bl, byte [rdi+rcx]
        cmp rbx, 0
        je _exit_loop_encrypt
        add bl, cl
        mov byte [rdi+rcx], bl
        add qword [rbp-8], 1
        jmp _loop_encrypt
    _exit_loop_encrypt:
        leave
        ret

write_file:
    push rbp
    mov rbp, rsp
    sub rsp, 8
    mov qword [rbp-8], 0

    push rsi

    mov rax, 2
    mov rsi, 0x241
    mov rdx, 0x1A4
    syscall
    mov qword [rbp-8], rax ; save fd

    pop rsi
    mov rdi, rsi
    call strlen

    mov rdx, rax
    mov rax, 1
    mov rdi, qword [rbp-8]
    syscall
    
    mov rax, 3
    mov rdi, qword [rbp-8]
    syscall

    leave
    ret
    
strlen:
    xor rax, rax
    xor rbx, rbx
    _strlen_loop:
        mov bl, byte [rdi + rax]
        inc rax
        cmp rbx, 0
        jne _strlen_loop
    dec rax
    ret

_exit:
    mov rax, 60
    xor rdi, rdi
    syscall
