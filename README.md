# Тестовый запуск Docker-Go

Этот репозиторий содержит пример простого веб-сервера на Go, который запускается в контейнере Docker. Проект демонстрирует базовые навыки работы с Docker и Go, а также основы создания и запуска контейнеризированных приложений.

## Описание файлов

### main.go

Этот файл содержит код простого HTTP-сервера на Go. Сервер отвечает на запросы к корневому URL ("/") сообщением "Тестовый запуск Go Docker".
```
package main

import (
  "fmt"
  "net/http"
)

// handler - это функция-обработчик для HTTP-запросов к корневому URL ("/")
func handler(w http.ResponseWriter, r *http.Request) {
  // Отправляем ответ "Тестовый запуск Go Docker" в браузер клиента
  fmt.Fprint(w, "Тестовый запуск Go Docker")
}

func main() {
  // http.HandleFunc связывает URL ("/") с функцией-обработчиком (handler)
  http.HandleFunc("/", handler)

  // Выводим сообщение в консоль о том, что сервер запускается
  fmt.Println("Server is starting")

  // Запускаем HTTP сервер на порту 8787
  // Если возникнет ошибка, она будет возвращена функцией ListenAndServe
  http.ListenAndServe(":8787", nil)
}
```
### Dockerfile

Этот файл описывает процесс сборки Docker-образа для нашего Go-приложения. Образ строится в два этапа: сборка (build-stage) и финальный минималистичный образ.

# Используем официальный образ Golang версии 1.22.5 в качестве начального этапа сборки
`FROM golang:1.22.5 AS build-stage`

# Устанавливаем рабочую директорию в контейнере
`WORKDIR /app`

# Копируем все файлы из текущей директории на хосте в рабочую директорию контейнера
`COPY . /app`

# Устанавливаем переменную окружения CGO_ENABLED в 0, чтобы отключить использование CGO
`ENV CGO_ENABLED=0`

# Компилируем Go-приложение и сохраняем исполняемый файл с именем server в директорию /app
`RUN go build -o /app/server /app/main.go`

# Используем минималистичный базовый образ для финального контейнера
`FROM scratch`

# Устанавливаем рабочую директорию в корневую директорию контейнера
`WORKDIR /`

# Копируем исполняемый файл из начального этапа сборки в текущую рабочую директорию
`COPY --from=build-stage /app/server server`

# Открываем порт 8787 для доступа к приложению из вне
`EXPOSE 8787`

# Устанавливаем команду по умолчанию для запуска при старте контейнера
`ENTRYPOINT [ "/server" ]`
## Инструкция по использованию

### При работе с Docker и Go

1. Установите Go и его расширение в VSCode.
2. Создайте файл main.go и напишите простенький веб-сервер (как показано выше).
3. Создайте файл Dockerfile (как показано выше).
4. Убедитесь, что Docker установлен на вашем компьютере.
5. Откройте Docker Desktop.
6. Соберите Docker-образ для вашего приложения командой:
  
   `docker build -t sf_server .`
   
   - -t указывает имя, под которым будет использоваться контейнер.
7. Запустите сервер командой:
  
   `docker run -d --rm -p 8788:8787 --name sf_server1 sf_server`
   
   - -d запускает контейнер в фоновом режиме.
   - --rm удаляет контейнер после остановки.
   - -p 8788:8787 перенаправляет порт 8788 на хосте к порту 8787 в контейнере.
   - --name указывает имя контейнера.
8. Для установки и запуска Selenoid с GitHub:
  
   `./selenoid_windows_amd64.exe selenoid start --browsers "chrome" --last-versions 1`
   
Теперь вы можете открыть браузер и перейти по адресу http://localhost:8788, чтобы увидеть сообщение "Тестовый запуск Go Docker".
