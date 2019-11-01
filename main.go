package main

import (
	"log"
	"net/http"

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

	utils.Logging.Fatal(http.ListenAndServe(":8081", router))

	// utils.Logging.Printf("port : %s", os.Getenv("PORT"))
	// utils.Logging.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), router))
}

func endpoint() {

	//test
	router.HandleFunc("/api/ping", test).Methods("GET")

	//mobile
	router.HandleFunc("/api/user/register", controller.RegisterUser).Methods("POST")
	router.HandleFunc("/api/user/login", controller.LoginUser).Methods("POST")
	router.HandleFunc("/api/user/get-all", controller.GetAllUser).Methods("GET")
	router.HandleFunc("/api/user/update", controller.EditProfile).Methods("POST")
	router.HandleFunc("/api/user/change-password", controller.ChangePassword).Methods("POST")

	//web
	router.HandleFunc("/api/workshop/register", controller.RegisterBengkel).Methods("POST")
	router.HandleFunc("/api/workshop/login", controller.LoginBengkel).Methods("POST")
	router.HandleFunc("/api/workshop/add-opration", controller.AddOpration).Methods("POST")
	router.HandleFunc("/api/workshop/update", controller.EditBengkel).Methods("POST")
	router.HandleFunc("/api/workshop/update-location", controller.EditLocation).Methods("POST")

	//umum
	router.HandleFunc("/api/workshop/provinsi={provinsi}", controller.GetBengkelProvinsi).Methods("GET")
	router.HandleFunc("/api/workshop/kota={kota}", controller.GetBengkelKota).Methods("GET")
	router.HandleFunc("/api/workshop/id-workshop={id}", controller.GetBengkelId).Methods("GET")
	router.HandleFunc("/api/workshop-all", controller.GetAllBengkel).Methods("GET")
	router.HandleFunc("/api/add-car", controller.AddKendaraan).Methods("POST")
	router.HandleFunc("/api/car/id-mobil={id}", controller.GetKendaraanId).Methods("GET")
	router.HandleFunc("/api/jenis-service", controller.GetJenisService).Methods("GET")
	router.HandleFunc("/api/add-service", controller.AddService).Methods("POST")
	router.HandleFunc("/api/service/id-workshop={id}", controller.GetService).Methods("GET")
	router.HandleFunc("/api/car/update", controller.UpdateKendaraan).Methods("POST")
}

func test(w http.ResponseWriter, r *http.Request) {
	utils.Response(w, utils.Message(true, "pong"))
}
