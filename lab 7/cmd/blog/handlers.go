package main

import (
	"database/sql"
	"fmt"
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

type featuredPostsHeaderData struct {
	Title string
}

type featuredPostsSectionData struct {
	Header featuredPostsHeaderData
	Posts  []PostDataDb
}

type mostResentPostsHeaderData struct {
	Title string
}

type mostResentPostsSectionData struct {
	Header mostResentPostsHeaderData
	Posts  []PostDataDb
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

type postPageMainContentData struct {
	Paragraphs []string
}

type HeaderPostData struct {
	Title     string
	Subtitle  string
	MainImage string
}

type PostDataDb struct {
	Title                 string `db:"title"`
	Subtitle              string `db:"subtitle"`
	PublishDate           string `db:"publish_date"`
	BackgroundImgModifier string `db:"image_url"`
	Author                string `db:"author"`
	AuthorImg             string `db:"author_url"`
	Content               string `db:"content"`
	Tags                  []string
}

type postData struct {
	Featured   featuredPostsSectionData
	MostResent mostResentPostsSectionData
}

type orderListData struct {
	OrderID  string `db:"order_id"`
	Title    string `db:"title"`
	OrderURL string // URL ордера, на который мы будем переходить для конкретного поста
}

type orderData struct {
	Title   string `db:"title"`
	Content string `db:"content"`
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

func featuredPosts() featuredPostsSectionData {
	return featuredPostsSectionData{
		Header: featuredPostsHeaderData{
			Title: "Featured Posts",
		},
	}
}

func mostResentPosts() mostResentPostsSectionData {
	return mostResentPostsSectionData{
		Header: mostResentPostsHeaderData{
			Title: "Most Recent",
		},
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
		postD, err := posts(db)
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}
		FeaturedPosts, err := GetFeaturedPostsDB(db) // Главная страница блога
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err.Error())                    // Используем стандартный логгер для вывода ошибки в консоль
			return                                      // Не забываем завершить выполнение ф-ии
		}

		MostResentposts, err := GetMostResentPostsDB(db) // Главная страница блога
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err.Error())                    // Используем стандартный логгер для вывода ошибки в консоль
			return                                      // Не забываем завершить выполнение ф-ии
		}

		ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}

		// Подготовим данные для шаблона
		//data := indexHeader()
		indexData := struct {
			Header          indexHeaderData
			NavigationPanel navigationPanelData
			posts           postData
			Footer          footerData
		}{
			Header:          indexHeader(),
			NavigationPanel: navigationPanel(),
			Footer:          footerinfo(),
		}

		indexData.posts.Featured.Posts = FeaturedPosts
		indexData.posts.MostResent.Posts = MostResentposts
		/*data := indexPageData{
			Title:         "Blog for traveling",
			Subtitle:      "My best blog for adventures and burgers",
			FeaturedPosts: posts,
		}*/

		err = ts.Execute(w, indexData) // Запускаем шаблонизатор для вывода шаблона в тело ответа
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err.Error())
			return
		}
	}
}

func GetFeaturedPostsDB(db *sqlx.DB) ([]PostDataDb, error) {
	const query = `
		SELECT
			title,
			subtitle,
			publish_date,
			author,
			author_url,
			image_url, 
			content
		FROM
			post
		WHERE featured = 1
	` // Составляем SQL-запрос для получения записей для секции featured-posts

	var posts []PostDataDb // Заранее объявляем массив с результирующей информацией

	err := db.Select(&posts, query) // Делаем запрос в базу данных
	if err != nil {                 // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return posts, nil
}

func GetMostResentPostsDB(db *sqlx.DB) ([]PostDataDb, error) {
	const query = `
		SELECT
			title,
			subtitle,
			publish_date,
			author,
			author_url,
			image_url,
			content
		FROM
			post
		WHERE featured = 0
	` // Составляем SQL-запрос для получения записей для секции featured-posts

	var posts []PostDataDb // Заранее объявляем массив с результирующей информацией

	err := db.Select(&posts, query) // Делаем запрос в базу данных
	if err != nil {                 // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return posts, nil
}

func postsData(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		orderIDStr := mux.Vars(r)["orderID"] // Получаем orderID в виде строки из параметров урла

		orderID, err := strconv.Atoi(orderIDStr) // Конвертируем строку orderID в число
		if err != nil {
			http.Error(w, "Invalid order id", 403)
			log.Println(err)
			return
		}

		order, err := postByID(db, orderID)
		if err != nil {
			if err == sql.ErrNoRows {
				// sql.ErrNoRows возвращается, когда в запросе к базе не было ничего найдено
				// В таком случае мы возвращем 404 (not found) и пишем в тело, что ордер не найден
				http.Error(w, "Post not found", 404)
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

		err = ts.Execute(w, order)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			log.Println(err)
			return
		}

		log.Println("Request completed successfully")
	}
}

// Возвращаем не просто []orderListData, а []*orderListData - так у нас получится подставить OrderURL в структуре
func posts(db *sqlx.DB) ([]*orderListData, error) {
	const query = `
		SELECT
			title,
			subtitle,
			publish_date,
			author,
			author_url,
			image_url,
			content
		FROM 
		  post
		`
	// Такое объединение строк делается только для таблицы order, т.к. это зарезерированное слово в SQL, наряду с SELECT, поэтому его нужно заключить в ``

	var posts []*orderListData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&posts, query) // Делаем запрос в базу данных
	if err != nil {                 // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	for _, post := range posts {
		post.OrderURL = "/post/" + post.OrderID // Формируем исходя из ID ордера в базе
	}

	fmt.Println(posts)

	return posts, nil
}

// Получает информацию о конкретном ордере из базы данных
func postByID(db *sqlx.DB, orderID int) (PostDataDb, error) {
	const query = `
		SELECT
			title,
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
