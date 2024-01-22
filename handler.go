package controller

import (
	"html/template"
	"net/http"
)

type Request struct {
	Page    string
	Message string
	Adventurer
}

type Adventurer struct {
	Name  string
	Class string
	Level int
	Age   int
	Genre string
}

var appState Request

func FormHandler(w http.ResponseWriter, r *http.Request) {

	appState.Page = "index"
	appState.Message = ""
	renderPage(w, "index", appState)
}

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	// name := strings.TrimSpace(r.FormValue("name"))
	// class := strings.TrimSpace(r.FormValue("class"))
	// levelStr := r.FormValue("level")

	// if name == "" || class == "" || levelStr == "" {
	// 	appState.Message = "Please fill in all fields."
	// 	renderPage(w, "template.html", appState)
	// 	return
	// }

	// level, err := strconv.Atoi(levelStr)
	// if err != nil {
	// 	appState.Message = "Invalid level value."
	// 	renderPage(w, "template.html", appState)
	// 	return
	// }

	// appState.Adventurer = Adventurer{
	//	Name:  "Dodo",
	//	Class: "A",
	//	Level: 1,
	//	Age:   16,
	//	Genre: "Feminin",
	// }

	// appState.Message = "Adventurer created successfully!"
	// appState.Page = "profile"
	renderPage(w, "create", appState)
}

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	//charge les données du profil à passer dans appState
	renderPage(w, "profile", appState)
}

func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	appState.Adventurer = Adventurer{
		Name:  "Dodo",
		Class: "A",
		Age:   16,
		Genre: "Feminin",
	}

	appState.Message = "Adventurer deleted successfully!"
	appState.Page = " "
	renderPage(w, "delete", appState)
}

func renderPage(w http.ResponseWriter, tmpl string, data interface{}) {
	var templates = template.Must(template.ParseFiles("tmp/" + tmpl + ".html"))

	templates.ExecuteTemplate(w, tmpl, data)
}
