# forensics | m4g1c

## Information
Медальон дрожит, это место силы. Отсюда изливается магия...

## Writeup
Видим образ файловой системы ext4.
```
$ file ../image.img 
../image.img: Linux rev 1.0 ext4 filesystem data, UUID=3437a200-7677-4567-88be-eb49b5ac2e15 (extents) (64bit) (large files) (huge files)
```

Монтируем её в любое удобное место.
```
sudo mount image.img /mnt/mount/
```
Внутри какой-то мусор: одинаковые папки, архивы, фото. Среди них есть самородок, и он лежит тут `<путь куда монтировали>/Pictures/photo_2023-12-20_20-12-59 (копия).jpg`

## Flag
`SgffCTF{p1vk4_dly4_r1vk4}`
