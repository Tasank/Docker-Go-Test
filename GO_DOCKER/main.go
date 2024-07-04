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
