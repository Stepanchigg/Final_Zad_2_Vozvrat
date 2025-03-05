# Final_Zad_2_Vozvrat
Итоговая задача модуля 2 Яндекс лицея.

Этот проект реализует веб-сервис, принимающий выражение через Http запрос и возвращабщий результат вычислений.

Инструкция по запуску:

1)Убедитесь, что у вас установлен Go (версия 1.16 или выше).

2)Скопируйте репозиторий(через git bash ):

```bash
git clone https://github.com/Stepanchigg/Final_Zad_2_Vozvrat
```

```bash
cd Final_Zad_2_Vozvrat
```

Запускаем orchestator:

```bash
export TIME_ADDITION_MS=200
export TIME_SUBTRACTION_MS=200
export TIME_MULTIPLICATIONS_MS=300
export TIME_DIVISIONS_MS=400

go run cmd/orchestrator/orchestrator_start.go
```

Вы получите ответ  Оркестратор запускается на порту 8080.

В новом окне bash:

Опять переходим в репозиторию с проектом:

```bash
cd Final_Zad_2_Vozvrat
```

Затем запускаем agent:

```bash
export COMPUTING_POWER=4
export ORCHESTRATOR_URL=http://localhost:8080

 go run cmd/agent/agent_start.go
```

Вы получите ответ:
Запускаем агент
Стартующий воркер 0
Стартующий воркер 1
Стартующий воркер 2
Стартующий воркер 3


Примеры использования:

Успешный запрос:

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '
{
  "expression": "2*2+2"
}'
```

Ответ:

```bash
{
  "id": "1"
}
```

После можно посмотреть этап выполнения данного запроса и его результат(если уже вычислилось ):

```bash
curl --location 'http://localhost:8080/api/v1/expressions'
```

Если вычисления выполнены то:

```bash
"expressions":{"id":"1","expression":"2*2+2","status":"завершено","result":6}
```

Или узнать точный результат нужного выражения по его точному id:

```bash
curl --location 'http://localhost:8080/api/v1/expressions/id'
```

Ошибки при запросах:

Ошибка 404(отсутствие выражения ):

```bash
{"error":"Выражение не найдено"}
```

Ошибка 422 (невалидное выражение ):

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '
{
  "expression": "2+a"
}'

```
Ответ:

```bash
{
  {"error":"неожиданное число на месте 2"}
}
```

Ошибка 500 (внутренняя ошибка сервера ):

```bash
curl --location 'http://localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '
{
  "expression": "2/0"
}'
```
Ответ(у  меня высвечивается изначально id созданной задачи,а после в git bash где был запущен agent_start можно увидеть что выводится деление на 0 ):

```bash
{
  Воркер n: Ошибка в вычислении задачи n: Деление на ноль
}
```

Тесты для agent запускаются тоже через git bash:

1)Сначала опять переходим в папку с модулем.

```bash
cd Final_Zad_2_Vozvrat
```

2)Затем запускаем тестирование:

```bash
go test ./internal/agent/agent_calculation_test.go
```

3)При успешном прохождение теста должен вывестись ответ:

```bash
ok      command-line-arguments  0.094s
```

4)При ошибке в тестах будет указано где она совершена.
P.S ошибка связанная с не указанным ErrDivivsionByZero появляется так как в функции тестирования я ее не оглашаю,
она создает конфликты в visual studio code так как уже присутствует в самом агенте
