package main

import (
	"log"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request){
	//w.Header().Set("Content-Type", "application/json")
	log.Println("Entro")
	http.ServeFile(w,r,"Front/index.html")
}

func main() {
	cssHandler := http.FileServer(http.Dir("./Front/css/"))
	js_Handler := http.FileServer(http.Dir("./Front/js/"))

	mux := mux.NewRouter()
	mux.HandleFunc("/", Home ).Methods("Get")
	http.Handle("/css/", http.StripPrefix("/css/", cssHandler))
	http.Handle("/js/", http.StripPrefix("/js/", js_Handler))

	handler := cors.Default().Handler(mux)

	log.Println("Sever listen in the port 8000")
	log.Fatal(  http.ListenAndServe(":8000", handler) )	
}




