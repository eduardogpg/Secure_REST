package main

import (
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	"io"
)

const port string = ":8000"

type Response struct{
	Message string `json:"message"`
}

func JSONResponse(w http.ResponseWriter, r *http.Request){
	response := CreateResponse()
	json.NewEncoder(w).Encode(response)
}

func TextResponse(w http.ResponseWriter, r *http.Request){
	io.WriteString(w, "Hello World")
}

func CreateResponse() Response{
	return Response{ "Hola Mundo" }
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/JSON", JSONResponse )
	mux.HandleFunc("/Hello", TextResponse )

	// go run generate_cert.go --host localhost --ca
	http.Handle("/", mux)
	err := http.ListenAndServeTLS(port, "cert.pem", "key.pem", nil)
	if err != nil{
		log.Fatal("ListenAndServe : ", err)
	}
}




