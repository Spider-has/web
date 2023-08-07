package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http" // Подключаем пакет с HTTP сервером
	"os"
	"strconv"
	"strings"
	"time"

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

type session struct {
	userId string
	expiry time.Time
}

type UserIDData struct {
	UserId string `db:"user_id"`
}

var sessions = map[string]session{}

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
	MainImg  string `db:"heroImg1Path"`
	Content  string `db:"content"`
}

type UserDataDb struct {
	UserId   string `db:"user_id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}

type UserIdDb struct {
	UserID int `db:"user_id"`
}

type userData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
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
	HeroImg1Name          string `db:"heroImg1Path"`
	HeroImg2Name          string `db:"heroImg2Path"`
	AuthorImgPath         string `db:"authorImgPath"`

	Tags    []string
	PostURL string // URL ордера, на который мы будем переходить для конкретного поста
}

type createPostRequest struct {
	Title         string `json:"title"`
	Subtitle      string `json:"subtitle"`
	PublishDate   string `json:"publish_date"`
	Author        string `json:"author"`
	AuthorImg     string `json:"authorAvatar"`
	HeroImg1      string `json:"heroImg1"`
	HeroImg2      string `json:"heroImg2"`
	Content       string `json:"content"`
	HeroImg1Name  string `json:"heroImg1Name"`
	HeroImg2Name  string `json:"heroImg2Name"`
	AuthorImgName string `json:"authorAvatarName"`
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

const adminSuccessUrl string = "http://localhost:3000/admin/post-settings"

func adminLogin(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) { // Функция для отдачи страницы
	return func(w http.ResponseWriter, r *http.Request) {

		err := authByCookie(db, w, r)
		if err == nil {
			w.Header().Set("Content-Type", "application/json")
			http.Redirect(w, r, adminSuccessUrl, http.StatusSeeOther)
			//http.Error(w, "No auth cookie passed", 401)
			log.Println(err)
			return
		}

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

const adminFailUrl string = "http://localhost:3000/admin"

func adminPage(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) { // Функция для отдачи страницы
	return func(w http.ResponseWriter, r *http.Request) {

		ts, err := template.ParseFiles("pages/admin_page.html") // Главная страница блога
		if err != nil {
			http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
			log.Println(err)
			return // Не забываем завершить выполнение ф-ии
		}

		err = authByCookie(db, w, r)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			http.Redirect(w, r, adminFailUrl, http.StatusSeeOther)
			//http.Error(w, "No auth cookie passed", 401)
			log.Println(err)
			return
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
			image_url,
			heroImg1Path,
			heroImg2Path,
			authorImgPath
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
			heroImg1Path,
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

func createPost(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqData, err := io.ReadAll(r.Body) // Прочитали тело запроса с reqData в виде массива байт
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		var req createPostRequest

		err = json.Unmarshal(reqData, &req) // Отдали reqData и req на парсинг библиотеке json
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		err = getDecodeImg(req.AuthorImg, req.AuthorImgName)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		err = getDecodeImg(req.HeroImg1, req.HeroImg1Name)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		err = getDecodeImg(req.HeroImg2, req.HeroImg2Name)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		err = savePost(db, req)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println(err)
			return
		}
	}
}

func getDecodeImg(codeImgbase64, imgName string) error {

	b64data := codeImgbase64[strings.IndexByte(codeImgbase64, ',')+1:]
	img, err := base64.StdEncoding.DecodeString(b64data)
	if err != nil {
		return err
	}

	file, err := os.Create("static/img/" + imgName)
	if err != nil {
		return err
	}

	_, err = file.Write(img)
	if err != nil {
		return err
	}
	return nil
}

func savePost(db *sqlx.DB, req createPostRequest) error {
	const imgPath = "./static/img/"
	const query = `
       INSERT INTO
           post
       (
			title,
			subtitle,
			publish_date,
			author,
			authorImgPath,
			featured,
			content,
			heroImg1Path,
			heroImg2Path
       )
       VALUES
       (
           ?,
           ?,
		   ?,
		   ?,
		   ?,
           ?,
		   ?,
		   ?,
		   ?
       );
	   `
	reqDbData := struct {
		Title         string
		Subtitle      string
		PublishDate   string
		AuthorName    string
		AuthorImgPath string
		featured      int
		content       string
		HeroImg1Path  string
		HeroImg2Path  string
	}{
		Title:         req.Title,
		Subtitle:      req.Subtitle,
		PublishDate:   req.PublishDate,
		AuthorName:    req.Author,
		AuthorImgPath: imgPath + req.AuthorImgName,
		featured:      0,
		content:       req.Content,
		HeroImg1Path:  imgPath + req.HeroImg1Name,
		HeroImg2Path:  imgPath + req.HeroImg2Name,
	}
	_, err := db.Exec(query, reqDbData.Title, reqDbData.Subtitle, reqDbData.PublishDate, reqDbData.AuthorName,
		reqDbData.AuthorImgPath, reqDbData.featured, reqDbData.content, reqDbData.HeroImg1Path, reqDbData.HeroImg2Path) // Сами данные передаются через аргументы к ф-ии Exec
	return err
}

const cookieName string = "Fucking_Cookie_For_Fucking_user"

func logination(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reqData, err := io.ReadAll(r.Body) // Прочитали тело запроса с reqData в виде массива байт
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		var req userData
		err = json.Unmarshal(reqData, &req) // Отдали reqData и req на парсинг библиотеке json
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log.Println(err)
			return
		}

		log.Println(req)

		user, err := findUserInDB(db, req)
		if err != nil {
			http.Error(w, "Incorrect password or email", 401)
			return
		}

		//expiresAt := time.Now().Add(120 * time.Second)
		//randomCookieName := randomStroke(15)

		sessions[string(rune(user.UserID))] = session{
			userId: string(rune(user.UserID)),
			expiry: time.Now().AddDate(0, 0, 1),
		}

		http.SetCookie(w, &http.Cookie{
			Name:    cookieName,                  // Устанавливаем имя куки
			Value:   fmt.Sprint(user.UserID),     // Конвертируем userID из user из типа int в string
			Path:    "/",                         // Говорим куке действовать по всем путям сайта
			Expires: time.Now().AddDate(0, 0, 1), // говорим куке протухнуть через день
		})

		w.WriteHeader(200)
	}
}

func findUserInDB(db *sqlx.DB, user userData) (UserIdDb, error) {
	const query = `
		SELECT
			user_id
		FROM
		 ` + "`user`" +
		`WHERE
			email = ? AND
			password = ?;
	`
	var userD UserIdDb

	// Обязательно нужно передать в параметрах orderID
	err := db.Get(&userD, query, user.Email, user.Password)
	if err != nil {
		return UserIdDb{}, err
	}

	return userD, nil
}

// const messageLetters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// func randomStroke(n int) string {
// 	b := make([]byte, n)
// 	for i := range b {
// 		b[i] = messageLetters[rand.Intn(len(messageLetters))]
// 	}
// 	return string(b)
// }

func redirectToAdmin(w http.ResponseWriter, r *http.Request) {

}

func authByCookie(db *sqlx.DB, w http.ResponseWriter, r *http.Request) error {
	// Получаем куку или реагируем на её отсутствие
	cookie, err := r.Cookie(cookieName)
	if err != nil {
		if err == http.ErrNoCookie {
			//http.Error(w, "No auth cookie passed", 401)
			log.Println(err)
			return err
		}
		http.Error(w, "Internal Server Error", 500)
		log.Println(err)
		return err
	}

	userIDStr := cookie.Value

	err = findUserInDBById(db, userIDStr)
	if err != nil {
		http.Error(w, "Incorrect password or email", 401)
		return err
	}

	return nil
}

func loginCheck(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := authByCookie(db, w, r)
		if err != nil {
			return
		}
		return
	}
}

func findUserInDBById(db *sqlx.DB, userId string) error {
	const query = `
		SELECT
			user_id
		FROM
		 ` + "`user`" +
		`WHERE
			user_id = ? 
	`
	var userD UserIdDb

	// Обязательно нужно передать в параметрах orderID
	err := db.Get(&userD, query, userId)
	log.Println(userD)
	log.Println(err)
	if err != nil {
		return err
	}

	return nil
}
