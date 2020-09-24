# Тестовое задание Авито в юнит Auto

Для запуска сервиса:

`docker-compose up`

Для создания сокращенного url:

`
curl -d '{"url":"https://habr.com/ru/company/nixys/blog/461723/"}' -H "Content-Type: application/json" -X POST http://localhost:9000/url
`

Перейти по сокращенному url:

`
curl http://localhost:9000/1
`

Для создание кастомного сокращенного url:

`
curl -d '{"url":"https://github.com/bxcodec/go-clean-arch","custom_key":"clean"}' -H "Content-Type: application/json" -X POST http://localhost:9000/url
`

Переход по кастомному сокращенному url:

`
curl http://localhost:9000/clean
`