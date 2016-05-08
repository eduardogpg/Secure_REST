package main 

import(
	"log"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
)

type Response struct{
	Message string `json:"message"`
}

func Message(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	response := CreateResponse( r.FormValue("message") )
	log.Println("Nuevo mensaje recibido")
	SendJsonMessage(w, response)
}

func CreateResponse(message string) Response{
	return Response{ message }
}

func SendJsonMessage(w http.ResponseWriter, response Response){
	json.NewEncoder(w).Encode(response)	
}

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/message", Message ).Methods("POST")

	http.Handle("/", mux)
	log.Println("Sever listen in the port 8001")
	log.Fatal(  http.ListenAndServe(":8001", nil) )
}