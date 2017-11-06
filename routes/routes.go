package routes

import(

	"net/http"

)

func Route(){

	http.HandleFunc("/", home)

	http.HandleFunc("/welcome", welcome)

	http.HandleFunc("/chat", chat)

}

