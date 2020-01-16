Теперь нужен метод для получения списка котов. На запрос:

curl -X GET http://localhost:8080/cats
Должен возвращаться список котов в формате JSON:

[
  {"name": "Tihon", "color": "red & white", "tail_length": 15, "whiskers_length": 12},
  {"name": "Marfa", "color": "black & white", "tail_length": 13, "whiskers_length": 11}
]
Должна работать сортировка по заданному атрибуту, по возрастанию или убыванию:

curl -X GET http://localhost:8080/cats?attribute=name&order=asc
curl -X GET http://localhost:8080/cats?attribute=tail_length&order=desc
Так же клиент должен иметь возможность запросить подмножество данных, указав offset и limit:

curl -X GET http://localhost:8080/cats?offset=10&limit=10
Разумеется, клиент может указать и сортировку, и лимит одновременно:

curl -X GET http://localhost:8080/cats?attribute=color&order=asc&offset=5&limit=2
Подумайте, что должен возвращать сервер, если указан несуществующий атрибут? Неправильный order? Offset больший, чем имеется данных в базе? Какие еще могут быть варианты невалидных запросов?

Обработайте такие запросы так, как считаете правильным.
В этом задании не лишними будут юнит-тесты, проверяющие, что ваша программа корректно обрабатывает валидные и невалидные входящие данные.