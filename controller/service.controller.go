package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/GoToyota/model"
	"github.com/GoToyota/utils"
	"github.com/gorilla/mux"
)

var GetJenisService = func(w http.ResponseWriter, r *http.Request) {

	res := model.GetJenisService()
	utils.Logging.Println(res)

	resLog, _ := json.Marshal(res)
	utils.Logging.Println(string(resLog))
	utils.Response(w, res)
}

var AddService = func(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	utils.Logging.Println(string(body))
	defer r.Body.Close()

	service := &model.Service{}
	err := json.Unmarshal(body, service)
	if err != nil {
		utils.Response(w, utils.Message(false, "Invalid request"))
		return
	}

	resService := service.AddService()

	resLog, _ := json.Marshal(resService)
	utils.Logging.Println(string(resLog))
	utils.Response(w, resService)
}

var GetService = func(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	idBengkel, _ := strconv.Atoi(param["id"])

	res := model.GetService(idBengkel)
	utils.Logging.Println(res)

	resLog, _ := json.Marshal(res)
	utils.Logging.Println(string(resLog))
	utils.Response(w, res)
}
