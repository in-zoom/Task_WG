В базе данных есть таблица cats с такой схемой:
          
          Table "public.cats" 
          
|     Column      |       Type        |
|:---             |:---               |
| name            | character varying |
| color           | cat_color         |
| tail_length     | integer           |
| whiskers_length | integer           | 

И она заполнена некоторыми данными, примерно такими:

| name  |     color     | tail_length | whiskers_length |
| :---: |     :---:     |         ---:|             ---:|
| Tihon | red & white   |          15 |              12 |
| Marfa | black & white |          13 |              11 |

Про котов мы знаем некоторую важную информацию, например имя, цвет, длину хвоста и усов.

Цвет котов определен как перечисляемый тип данных:
``` 
CREATE TYPE cat_color AS ENUM (
    'black',
    'white',
    'black & white',
    'red',
    'red & white',
    'red & black & white'
);
``` 
Нужно выяснить, сколько котов каждого цвета есть в базe и записать эту информацию в таблицу cat_colors_info:
```
Table "public.cat_colors_info"
```
                    
| Column |   Type    |
| :---:  |   :---:   |
| color  | cat_color |
| count  | integer   |
```
Indexes:
    "cat_colors_info_color_key" UNIQUE CONSTRAINT, btree (color)
```
Должно получиться примерно так:

|        color        | count |
|:---                 |   ---:|
| black & white       |    1  |
| red & white         |    1  |
