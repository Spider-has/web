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
	Posts  []featuredPostDataDB
}

type mostResentPostDataDB struct {
	Title                 string `db:"title"`
	Subtitle              string `db:"subtitle"`
	PublishDate           string `db:"publish_date"`
	BackgroundImgModifier string `db:"image_url"`
	Author                string `db:"author"`
	AuthorImg             string `db:"author_url"`
	Tags                  []string
}

type mostResentPostsHeaderData struct {
	Title string
}

type mostResentPostsSectionData struct {
	Header mostResentPostsHeaderData
	Posts  []mostResentPostDataDB
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

type featuredPostDataDB struct {
	Title                 string `db:"title"`
	Subtitle              string `db:"subtitle"`
	PublishDate           string `db:"publish_date"`
	BackgroundImgModifier string `db:"image_url"`
	Author                string `db:"author"`
	AuthorImg             string `db:"author_url"`
	Tags                  []string
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

func postMainContent() postPageMainContentData {
	return postPageMainContentData{
		Paragraphs: []string{"Dark spruce forest frowned on either side the frozen waterway. The trees had been stripped by a recent wind of their white covering of frost, and they seemed to lean towards each other, black and ominous, in the fading light. A vast silencer igned over the land. The land itself was a desolation, lifeless, without movement, so lone and cold that the spirit of it was not even that of sadness. There was ahint in it of laughter, but of a laughter more terrible than an sadness—a laughter that wasmirthless as the smile of the sphinx, a laughter cold as the frost and partaking of the grimness of infallibility. It was the masterful and incommunicablewisdom of eternity laughing at the futility of life and the effort of life. It was the Wild, thesavage, frozen-hearted Northland Wild.",
			"But there was life, abroad in the land and defiant. Down the frozen waterway toiled a string of  wolfish  dogs. Their bristly fur was rimed with frost. Their breath froze in the air as it left their mouths,  spouting forth in spumes of vapour that settled upon the hair of their bodies and formed into  crystals  of  frost. Leather harness was on the dogs, and leather traces attached them to a sled which dragged  along  behind. The sled was without runners. It was made of stout birch-bark, and its full surface rested  on  the  snow. The front end of the sled was turned up, like a scroll, in order to force down and under the  bore  of  soft snow that surged like a wave before it. On the sled, securely lashed, was a long and narrow  oblong  box.  There were other things on the sled—blankets, an axe, and a coffee-pot and frying-pan; but  prominent,  occupying most of the space, was the long and narrow oblong box.",
			"In advance of the dogs, on wide snowshoes, toiled a man. At the rear of the sled toiled a second  man. On  the  sled, in the box, lay a third man whose toil was over,—a man whom the Wild had conquered and beaten  down  until he would never move nor struggle again. It is not the way of the Wild to like movement. Life  is an  offence to it, for life is movement; and the Wild aims always to destroy movement. It freezes the  water  to  prevent it running to the sea; it drives the sap out of the trees till they are frozen to their  mighty  hearts; and most ferociously and terribly of all does the Wild harry and crush into submission  man—man  who  is the most restless of life, ever in revolt against the dictum that all movement must in the end  come  to  the cessation of movement.",
			"But at front and rear, unawed and indomitable, toiled the two men who were not yet dead. Their  bodies  were  covered with fur and soft-tanned leather. Eyelashes and cheeks and lips were so coated with the  crystals  from their frozen breath that their faces were not discernible. This gave them the seeming of  ghostly  masques, undertakers in a spectral world at the funeral of some ghost. But under it all they were  men,  penetrating the land of desolation and mockery and silence, puny adventurers bent on colossal  adventure,  pitting themselves against the might of a world as remote and alien and pulseless as the abysses of  space.",
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

func postHeaderInfo() HeaderPostData {
	return HeaderPostData{
		Title:     "The Road Ahead",
		Subtitle:  "The road ahead might be paved - it might not be.",
		MainImage: "./static/img/road-ahead_main-img.png",
	}
}

func index(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) { // Функция для отдачи страницы
	return func(w http.ResponseWriter, r *http.Request) {
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
			FeaturedPosts   featuredPostsSectionData
			MostResentPosts mostResentPostsSectionData
			Footer          footerData
		}{
			Header:          indexHeader(),
			NavigationPanel: navigationPanel(),
			FeaturedPosts:   featuredPosts(),
			MostResentPosts: mostResentPosts(),
			Footer:          footerinfo(),
		}

		indexData.FeaturedPosts.Posts = FeaturedPosts
		indexData.MostResentPosts.Posts = MostResentposts
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

func theRoadAhead(w http.ResponseWriter, r *http.Request) { // Функция для отдачи страницы
	ts, err := template.ParseFiles("pages/the-road-ahead.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошибки в консоль
		return                                      // Не забываем завершить выполнение ф-ии
	}

	// Подготовим данные для шаблона
	indexData := struct {
		Header      indexHeaderData
		HeaderPost  HeaderPostData
		PostContent postPageMainContentData
		Footer      footerData
	}{
		Header:      indexHeader(),
		HeaderPost:  postHeaderInfo(),
		PostContent: postMainContent(),
		Footer:      footerinfo(),
	}
	err = ts.Execute(w, indexData) // Запускаем шаблонизатор для вывода шаблона в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}

func GetFeaturedPostsDB(db *sqlx.DB) ([]featuredPostDataDB, error) {
	const query = `
		SELECT
			title,
			subtitle,
			publish_date,
			author,
			author_url,
			image_url
		FROM
			post
		WHERE featured = 1
	` // Составляем SQL-запрос для получения записей для секции featured-posts

	var posts []featuredPostDataDB // Заранее объявляем массив с результирующей информацией

	err := db.Select(&posts, query) // Делаем запрос в базу данных
	if err != nil {                 // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return posts, nil
}

func GetMostResentPostsDB(db *sqlx.DB) ([]mostResentPostDataDB, error) {
	const query = `
		SELECT
			title,
			subtitle,
			publish_date,
			author,
			author_url,
			image_url
		FROM
			post
		WHERE featured = 0
	` // Составляем SQL-запрос для получения записей для секции featured-posts

	var posts []mostResentPostDataDB // Заранее объявляем массив с результирующей информацией

	err := db.Select(&posts, query) // Делаем запрос в базу данных
	if err != nil {                 // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	return posts, nil
}

func post(db *sqlx.DB) func(w http.ResponseWriter, r *http.Request) {
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
			post_id,
			title
		FROM 
		  post
		`
	// Такое объединение строк делается только для таблицы order, т.к. это зарезерированное слово в SQL, наряду с SELECT, поэтому его нужно заключить в ``

	var orders []*orderListData // Заранее объявляем массив с результирующей информацией

	err := db.Select(&orders, query) // Делаем запрос в базу данных
	if err != nil {                  // Проверяем, что запрос в базу данных не завершился с ошибкой
		return nil, err
	}

	for _, order := range orders {
		order.OrderURL = "/post/" + order.OrderID // Формируем исходя из ID ордера в базе
	}

	fmt.Println(orders)

	return orders, nil
}

// Получает информацию о конкретном ордере из базы данных
func postByID(db *sqlx.DB, orderID int) (orderData, error) {
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

	var post orderData

	// Обязательно нужно передать в параметрах orderID
	err := db.Get(&post, query, orderID)
	if err != nil {
		return orderData{}, err
	}

	return post, nil
}
