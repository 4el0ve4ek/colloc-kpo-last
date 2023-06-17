
## Описание API

### Добавление физ активности 

- POST
- /activity
- Body example:
 ```json
    {
       "activity" : "running",
       "duration" : 3,
       "calories" : 3
    }
 ```
- curl:
```shell
curl "localhost:8080/activity" --data '{"activity": "running", "duration": 3, "calories": 3}' -v
```

### Добавление питания

- POST
- /nutrition
- Body example:
 ```json
    {
       "dish" : "beef",
       "size" : 3,
       "calories" : 3
    }
 ```
- curl:
```shell
curl "localhost:8080/nutrition" --data '{"dish": "beef", "size": 3, "calories": 3}' -v
```


### Добавление сна

- POST
- /sleep
- Body example:
 ```json
    {
       "duration" : 3
    }
 ```
- curl:
```shell
curl "localhost:8080/sleep" --data '{"duration": 4}' -v
```

### Получение статистики

- Get
- /stats
- curl:
```shell
curl "localhost:8080/stats" | jq .
```
- Response example:
```json
{
  "lost_calories": 12,
  "gained_calories": 6,
  "activity_time_seconds": 12,
  "sleep_time_seconds": 8
}
```