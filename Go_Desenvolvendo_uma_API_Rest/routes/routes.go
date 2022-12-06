package routes

import (
	"log"
	"net/http"

	"github.com/furtadomn/golang-studies/Go_Desenvolvendo_uma_API_Rest/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
