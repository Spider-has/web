package main

import (
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql" // Импортируем для возможности подключения к MySQL
)

func main() {
	const port = ":3000" // Выносим значение порта в константу

	mux := http.NewServeMux()      // Сущность Mux, которая позволяет маршрутизировать запросы к определенным обработчикам,
	mux.HandleFunc("/home", index) // Прописываем, что по пути /home выполнится наш index, отдающий нашу страницу
	mux.HandleFunc("/post", theRoadAhead)

	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	log.Println("Start server " + port) // Пишем в консоль о том, что стартуем сервер
	err := http.ListenAndServe(port, mux)
	if err != nil {
		log.Fatal(err) // Падаем с логированием ошибки, в случае если не получилось запустить сервер
	}
}
