# pwn | mafioznik

## Information

Помоги Вито Скалетта вернуться домой. Он очень соскуфился.

P.S. Дабы упростить таск здесь отключен ASLR.

## Writeup

Этот таск был на ROP и ret2libc. В нём был отключен ASLR.

### Exploit Preparation

Здесь нам понадобятся два гаджета.

Первый для вызова puts для лика адреса libc и функции system для запуска /bin/sh.
```
0x00000000004011a5 : pop rdi ; ret
```

Второй для выравнивания стека.
```
0x000000000040101a : ret
```

Я использую ROPgadget для поиска гаджетов `ROPgadget --binary <binary>`

В начале получаем все offset'ы из libc.
```
SYSTEM_LIBC_OFFSET = 0x50d70
PUTS_LIBC_OFFSET = 0x80e50
BIN_SH_LIBC_OFFSET = 0x1d8678
```
Полезные команды для этого:
```
readelf -s <libc> - чекнуть адреса фунок в бинаре
strings -a -t x <libc> | grep "/bin/sh" - поиск /bin/sh в libc
```
Также получаем адреса в бинаре.
```
GADGET_ADDR = 0x4011a5 # pop rdi; ret
GADGET_STACK_ALIGNMENT_ADDR = 0x40101a # ret
PUTS_GOT_PLT_ADDR = 0x404018
PUTS_PLT_SEC_ADDR = 0x401070
```
Эти адреса можно глянуть в IDA.

### Exploit

1. Используем первый гаджет. С помощью puts читаем адрес puts в got.plt. Вычитаем от этого адреса наше смещение PUTS_LIBC_OFFSET и получаем LIBC_ADDR - адрес начало библиотеки libc.

2. Далее снова подключаемся, теперь мы уже знаем адрес libc. Остаётся вызвать функцию system с аргументом "/bin/sh". Для этого опять используем первый гаджет.

Во время двух этих шагов необходимо выравнивание стека, для этого используем второй гаджет.

Эксплойт есть [тут](exploit.py).

## Flag
`SgffCTF{m4f1a2_1s_th3_b35t}`
