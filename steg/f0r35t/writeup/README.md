# steg | f0r35t

## Information
Вам ведь тоже кажется, что фотография обрезана?

## Writeup
Увеличиваем высоту фото.
![alt text](img/size.png)
Считаем crc-32 от этих байт и заменяем его. Фиолетовые байты - crc-32.
![alt text](img/crc-32.png)
Profit!
![alt text](../task/original.png)
Подробнее про структуру png можно узнать [здесь](https://habr.com/ru/articles/130472/).

## Flag
`SgffCTF{4x3_4tt4ck5}`