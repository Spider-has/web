package main

import (
	"html/template"
	"log"
	"net/http" // Подключаем пакет с HTTP сервером
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

type featuredPostData struct {
	Title                 string
	Subtitle              string
	BackgroundImgModifier string
	Author                string
	AuthorImg             string
	PublishDate           string
	Tags                  []string
}

type featuredPostsSectionData struct {
	Header featuredPostsHeaderData
	Posts  []featuredPostData
}

type mostResentPostData struct {
	PostImage   string
	Title       string
	Subtitle    string
	Author      string
	AuthorImg   string
	PublishDate string
	Tags        []string
}

type mostResentPostsHeaderData struct {
	Title string
}

type mostResentPostsSectionData struct {
	Header mostResentPostsHeaderData
	Posts  []mostResentPostData
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
		Posts: []featuredPostData{
			{
				Title:                 "The road Ahead",
				Subtitle:              "The road ahead might be paved - it might not be.",
				BackgroundImgModifier: "post_img_type2",
				Author:                "Met Vogels",
				AuthorImg:             "article-footer__icon-place_icon_type1",
				PublishDate:           "September 25, 2015",
				Tags:                  []string{},
			},
			{
				Title:                 "From Top Down",
				Subtitle:              "Once a year, go someplace you’ve never been before.",
				BackgroundImgModifier: "post_img_type1",
				Author:                "William Wong",
				AuthorImg:             "article-footer__icon-place_icon_type2",
				PublishDate:           "September 25, 2015",
				Tags:                  []string{"ADVENTURE"},
			},
		},
	}
}

func mostResentPosts() mostResentPostsSectionData {
	return mostResentPostsSectionData{
		Header: mostResentPostsHeaderData{
			Title: "Most Recent",
		},
		Posts: []mostResentPostData{
			{
				PostImage:   "./static/img/most-resent_post1-img.jpg",
				Title:       "Still Standing Tall",
				Subtitle:    "Life begins at the end of your comfort zone.",
				Author:      "William Wong",
				AuthorImg:   "article-footer__icon-place_icon_type2",
				PublishDate: "9/25/2015",
				Tags:        []string{},
			},
			{
				PostImage:   "./static/img/most-resent_post2-img.jpg",
				Title:       "Sunny Side Up",
				Subtitle:    "No place is ever as bad as they tell you it’s going to be.",
				Author:      "Mat Vogels",
				AuthorImg:   "article-footer__icon-place_icon_type1",
				PublishDate: "9/25/2015",
				Tags:        []string{},
			},
			{
				PostImage:   "./static/img/most-resent_post3-img.jpg",
				Title:       "Water Falls",
				Subtitle:    "We travel not to escape life, but for life not to escape us.",
				Author:      "Mat Vogels",
				AuthorImg:   "article-footer__icon-place_icon_type1",
				PublishDate: "9/25/2015",
				Tags:        []string{},
			},
			{
				PostImage:   "./static/img/most-resent_post4-img.jpg",
				Title:       "Through the Mist",
				Subtitle:    "Travel makes you see what a tiny place you occupy in the world.",
				Author:      "William Wong",
				AuthorImg:   "article-footer__icon-place_icon_type2",
				PublishDate: "9/25/2015",
				Tags:        []string{},
			},
			{
				PostImage:   "./static/img/most-resent_post5-img.jpg",
				Title:       "Awaken Early",
				Subtitle:    "Not all those who wander are lost.",
				Author:      "Mat Vogels",
				AuthorImg:   "article-footer__icon-place_icon_type1",
				PublishDate: "9/25/2015",
				Tags:        []string{},
			},
			{
				PostImage:   "./static/img/most-resent_post6-img.jpg",
				Title:       "Try it Always",
				Subtitle:    "The world is a book, and those who do not travel read only one page.",
				Author:      "Mat Vogels",
				AuthorImg:   "article-footer__icon-place_icon_type1",
				PublishDate: "9/25/2015",
				Tags:        []string{},
			},
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

func index(w http.ResponseWriter, r *http.Request) { // Функция для отдачи страницы
	ts, err := template.ParseFiles("pages/index.html") // Главная страница блога
	if err != nil {
		http.Error(w, "Internal Server Error", 500) // В случае ошибки парсинга - возвращаем 500
		log.Println(err.Error())                    // Используем стандартный логгер для вывода ошибки в консоль
		return                                      // Не забываем завершить выполнение ф-ии
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

	err = ts.Execute(w, indexData) // Запускаем шаблонизатор для вывода шаблона в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
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
