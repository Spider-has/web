package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http" // Подключаем пакет с HTTP сервером
	"strconv"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

type indexHeaderData struct {
	Title           string
	Subtitle        string
	MainPhrase      string
	NavigationList  []string
	BackgroundImage string
	ButtonContent   string
}

type navigationPanelData struct {
	NavigationPanelElements []string
}

type formData struct {
	Title                 string
	InputFieldPlaceholder string
	ButtonContent         string
}

type footerData struct {
	Form            formData
	BackgroundImage string
	MainPhrase      string
	NavigationList  []string
}

type HeaderPostData struct {
	Title     string
	Subtitle  string
	MainImage string
}

type PostDataDb struct {
	Title    string `db:"title"`
	Subtitle string `db:"subtitle"`
	MainImg  string `db:"image_path"`
	Content  string `db:"content"`
}

type PostListData struct {
	PostID                string `db:"post_id"`
	Title                 string `db:"title"`
	Subtitle              string `db:"subtitle"`
	PublishDate           string `db:"publish_date"`
	Author                string `db:"author"`
	AuthorImg             string `db:"author_url"`
	Featured              int    `db:"featured"`
	BackgroundImgModifier string `db:"image_url"`

	Tags    []string
	PostURL string // URL ордера, на который мы будем переходить для конкретного поста
}

func indexHeader() indexHeaderData {
	return indexHeaderData{
		Title:           "Let's do it together",
		Subtitle:        "We travel the world in search of stories. Come along for the ride.",
		MainPhrase:      "Escape.",
		NavigationList:  []string{"HOME", "CATEGORIES", "ABOUT", "CONTACT"},
		BackgroundImage: "header_main-page",
		ButtonContent:   "View Latests Posts",
	}
}

func navigationPanel() navigationPanelData {
	return navigationPanelData{
		NavigationPanelElements: []string{"Nature", "Photography", "Relaxation", "Vacation", "Travel", "Adventure"},
	}
}

func footerinfo() footerData {
	return footerData{
		Form: formData{
			Title:                 "Stay in Touch",
			InputFieldPlaceholder: "Enter your email address",
			ButtonContent:         "Submit",
		},
		BackgroundImage: "footer_background-img",
		MainPhrase:      "Escape.",
		NavigationList:  []string{"HOME", "CATEGORIES", "ABOUT", "CONTACT"},
	}
}

func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) { // Функция для отдачи страницы
	return func(w http.ResponseWriter, r *http.Request) {
		PostList, err := posts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}

		ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}

		indexData := struct {
			Header          indexHeaderData
			NavigationPanel navigationPanelData
			Posts           []*PostListData
			Footer          footerData
		}{
			Header:          indexHeader(),
			NavigationPanel: navigationPanel(),
			Posts:           PostList,
			Footer:          footerinfo(),
		}

		err = ts.Execute(w, indexData) // Запускаем шаблонизатор для вывода шаблона в тело ответа
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}
	}
}

func adminLogin(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) { // Функция для отдачи страницы
	return func(w http.ResponseWriter, r *http.Request) {

		ts, err := template.ParseFiles("pages/admin.html") // Главная страница блога
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}

		adminData := struct{}{}

		err = ts.Execute(w, adminData) // Запускаем шаблонизатор для вывода шаблона в тело ответа
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}
	}
}

func adminPage(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) { // Функция для отдачи страницы
	return func(w http.ResponseWriter, r *http.Request) {

		ts, err := template.ParseFiles("pages/admin_page.html") // Главная страница блога
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}

		adminData := struct{}{}

		err = ts.Execute(w, adminData) // Запускаем шаблонизатор для вывода шаблона в тело ответа
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}
	}
}

func post(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		postIDStr := mux.Vars(r)["postID"] // Получаем orderID в виде строки из параметров урла

		postID, err := strconv.Atoi(postIDStr) // Конвертируем строку orderID в число
		if err != nil {
			http.Error(w, "Invalid order id", 403)
			log.Println(err)
			return
		}

		postData := struct {
			Header indexHeaderData
			Post   PostDataDb
			Footer footerData
		}{
			Header: indexHeader(),
			Footer: footerinfo(),
		}

		post, err := postByID(db, postID)
		if err != nil {
			if err == sql.ErrNoRows {
				// sql.ErrNoRows возвращается, когда в запросе к базе не было ничего найдено
				// В таком случае мы возвращем 404 (not found) и пишем в тело, что ордер не найден
				http.Error(w, "Order not found", 404)
				log.Println(err)
				return
			}

			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		ts, err := template.ParseFiles("pages/post.html")
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}
		postData.Post = post
		err = ts.Execute(w, postData)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		log.Println("Request completed successfully")
	}
}

// Возвращаем не просто []orderListData, а []*orderListData - так у нас получится подставить OrderURL в структуре
func posts(db *sqlx.DB) ([]*PostListData, error) {
	const query = `
		SELECT
			post_id,	
			title,
			subtitle,
			publish_date,
			author,
			author_url,
			featured,
			image_url
		FROM 
		  post
		`
	// Такое объединение строк делается только для таблицы order, т.к. это зарезерированное слово в SQL, наряду с SELECT, поэтому его нужно заключить в ``

	var posts []*PostListData // Заранее объявляем массив с результирующей информацией
	log.Println(posts)
	err := db.Select(&posts, query) // Делаем запрос в базу данных
	if err != nil {                 // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	for _, post := range posts {
		post.PostURL = "/post/" + post.PostID // Формируем исходя из ID ордера в базе
	}

	return posts, nil
}

// Получает информацию о конкретном ордере из базы данных
func postByID(db *sqlx.DB, orderID int) (PostDataDb, error) {
	const query = `
		SELECT
			title,
			subtitle,
			image_path,
			content
		FROM
			post
		WHERE
			post_id = ?
	`
	// В SQL-запросе добавились параметры, как в шаблоне. ? означает параметр, который мы передаем в запрос ниже

	var post PostDataDb

	// Обязательно нужно передать в параметрах orderID
	err := db.Get(&post, query, orderID)
	if err != nil {
		return PostDataDb{}, err
	}

	return post, nil
}
