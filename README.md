<h1>Тестовое задание</h1>
<h2>Задание:</h2>
Нужно разработать HTTP сервис для быстрого поиска анаграмм в словаре.
Два слова считаются анаграммами, если одно можно получить из другого перестановкой букв (без учета регистра).
Примеры анаграмм:<br>
["foobar", "barfoo", "boofar"]<br>
["живу", "вижу"]<br>
["Abba", "BaBa"]<br>
Примеры строк, не являющихся анаграммами:<br>
["abba", "bba"] - во второй строке только одна буква "а"<br>
Сервис должен предоставлять эндпойнт для загрузки списка слов в формате json. Пример использования:<br>
curl localhost:8080/load -d '["foobar", "aabb", "baba", "boofar", "test"]'<br>
И эндпойнт для поиска анаграмм по слову в загруженном словаре. Примеры использования:<br>
curl 'localhost:8080/get?word=foobar' => ["foobar","boofar"]<br>
curl 'localhost:8080/get?word=raboof' => ["foobar","boofar"]<br>
curl 'localhost:8080/get?word=abba' => ["aabb","baba"]<br>
curl 'localhost:8080/get?word=test' => ["test"]<br>
curl 'localhost:8080/get?word=qwerty' => null<br>
<h2>Решение:</h2>
Сервер, принимающий и отдающий массивы анаграмм в течение жизни процесса
Работает с POST и GET запросами, отдает и принимает массивы анаграмм. При запросе несуществующего слова возвращает
пустой массив строк.
<h2>Недостатки:</h2>
- избыточная сложность цикла в функциях обработки анаграмм - файл word_handle
