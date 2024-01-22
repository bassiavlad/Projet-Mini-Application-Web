package route

import (
	"golanta/controller"
	"net/http"

	"github.com/gorilla/mux"
)

func SetRoutes() {
	r := mux.NewRouter()
	http.Handle("/", r)
	r.HandleFunc("/index", controller.FormHandler).Methods("GET")
	r.HandleFunc("/create", controller.CreateHandler).Methods("GET")
	r.HandleFunc("/profile", controller.ProfileHandler).Methods("GET")
	r.HandleFunc("/delete", controller.DeleteHandler).Methods("GET")
}
