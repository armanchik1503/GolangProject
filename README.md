# GolangProject
Golang project

Технические требования:
1) Ваш сервис должен быть RestFul на базе протокола http(https://restfulapi.net/)
2) Должна быть аутентификация с использованием JWT. Длительность валидности токена должна быть не более 15 мин.(https://jwt.io/introduction)
3) Ресурсы для авторизации, регистрации(c подтверждением почтового ящика) и смены пароля.
4) Минимум 7 ресурсов в зависимости от вашего проекта(эти ресурсы должны валидировать токен и узнавать пользователя при помощи токена).
5) Должна быть использована как минимум одна база.(SQL, Nosql, Key-Value).
6) Проект должен быть покрыт юнит тестами(кроме main.go).
7) Нужно придерживаться Clean Architecture(https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html).
8) Нужно подать проект в виде Docker контейнера(то есть вы должны создать образ для вашего проекта) https://citizix.com/how-to-run-a-golang-revel-app-with-docker-and-docker-compose/
9) Ваш проект должен с самого начала разрабатываться на github, для просмотра истории вашего проекта. 
