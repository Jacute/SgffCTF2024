BITS 64

section .data
    flag db "SgffCTF{0nly_p41n_4ll_1n_v41n}", 0
    filename db "encrypted.txt", 0

section .text
    global _start

; i   [rbp-72] qword
; buf [rbp-64] 64 bytes
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


; copy rdi to rsi
; i [rbp-8] qword
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
        ; mov rax, qword [rbp-8]
        leave
        ret

; rdi - ciphertext
; i [rbp-8] qword
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

; rdi - filename, rsi - buf
; fd [rbp-8] qword
write_file:
    push rbp
    mov rbp, rsp
    sub rsp, 8
    mov qword [rbp-8], 0

    push rsi ; save rsi

    ; open file
    mov rax, 2
    mov rsi, 0x241
    mov rdx, 0x1A4
    syscall
    mov qword [rbp-8], rax ; save fd

    pop rsi
    mov rdi, rsi
    call strlen

    ; write file
    mov rdx, rax
    mov rax, 1
    mov rdi, qword [rbp-8]
    syscall
    
    ; close file
    mov rax, 3
    mov rdi, qword [rbp-8]
    syscall

    leave
    ret
    
; rdi - str ptr
; return string length
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
