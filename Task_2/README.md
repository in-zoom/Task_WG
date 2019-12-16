Продолжим анализ наших котов.
Нужно вычислить некоторые статистические данные о котах:
* средняя длина хвоста,
* медиана длин хвостов,
* мода длин хвостов,
* средняя длина усов,
* медиана длин усов,
* мода длин усов.

И сохранить эту информацию в таблицу cats_stat:
```
Table "public.cats_stat"
         Column         |   Type
------------------------+-----------
 tail_length_mean       | numeric
 tail_length_median     | numeric
 tail_length_mode       | integer[]
 whiskers_length_mean   | numeric
 whiskers_length_median | numeric
 whiskers_length_mode   | integer[]
```
Должно получиться примерно так:
```
 tail_length_mean | tail_length_median | tail_length_mode
------------------+--------------------+------------------
             14.0 |               14.0 | {13,15}

 whiskers_length_mean | whiskers_length_median | whiskers_length_mode
----------------------+------------------------+----------------------
                 11.5 |                   11.5 | {11,12}
```
