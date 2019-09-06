package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"github.com/GoToyota/app"
	"github.com/GoToyota/controller"
	"github.com/GoToyota/utils"

	"github.com/GoToyota/model"
)

var router *mux.Router

func main() {

	err := utils.Logger()
	if err != nil {
		log.Fatalln(err)
	}

	err = model.Init()
	if err != nil {
		log.Fatalln(err)
	}

	router = mux.NewRouter()
	endpoint()

	router.Use(app.JwtAuthentication)

	port := os.Getenv("port")
	if port == "" {
		port = "9000"
	}

	utils.Logging.Printf("port : %s", port)
	utils.Logging.Fatal(http.ListenAndServe(":"+port, router))
}

func endpoint() {
	router.HandleFunc("/api/ping", test).Methods("GET")
	//mobile
	router.HandleFunc("/api/user/register", controller.RegisterUser).Methods("POST")
	router.HandleFunc("/api/user/login", controller.LoginUser).Methods("POST")
	router.HandleFunc("/api/user/get-all", controller.GetAllUser).Methods("GET")

	//web
	router.HandleFunc("/api/workshop/register", controller.RegisterBengkel).Methods("POST")
	router.HandleFunc("/api/workshop/login", controller.LoginBengkel).Methods("POST")
	router.HandleFunc("/api/workshop/add-opration", controller.AddOpration).Methods("POST")

	//umum
	router.HandleFunc("/api/workshop/{kota}", controller.GetBengkelKota).Methods("GET")

}

func test(w http.ResponseWriter, r *http.Request) {
	utils.Response(w, utils.Message(true, "pong"))
}
