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

type footerData struct {
	form struct {
		header                string
		inputFieldPlaceholder string
		buttonContent         string
	}
	backgroundImage string
	mainPhrase      string
	navigationList  []string
}

type navigationPanelData struct {
	NavigationPanelElements []string
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
				PostImage:   "",
				Title:       "Still Standing Tall",
				Subtitle:    "Life begins at the end of your comfort zone.",
				Author:      "William Wong",
				AuthorImg:   "article-footer__icon-place_icon_type2",
				PublishDate: "9/25/2015",
				Tags:        []string{},
			},
			{},
		},
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
	}{
		Header:          indexHeader(),
		NavigationPanel: navigationPanel(),
		FeaturedPosts:   featuredPosts(),
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
	data := indexHeader()
	err = ts.Execute(w, data) // Запускаем шаблонизатор для вывода шаблона в тело ответа
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		log.Println(err.Error())
		return
	}
}
