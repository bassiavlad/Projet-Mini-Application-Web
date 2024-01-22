package main

import (
	"fmt"
	"golanta/route"
	"net/http"
)

func main() {
	route.SetRoutes()

	//Chargement des fichiers statics à ajouter
	http.Handle("/static/asset/img", http.StripPrefix("/static/controller/route", http.FileServer(http.Dir("static"))))
	http.Handle("/static/asset/css", http.StripPrefix("/static/controller/route", http.FileServer(http.Dir("static"))))

	fmt.Println("\nLien vers le site : http://localhost:8080/index (CTRL+CLICK)")
	http.ListenAndServe("localhost:8080/", nil)
}
