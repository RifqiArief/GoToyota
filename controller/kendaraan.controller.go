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

var AddKendaraan = func(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	utils.Logging.Println(string(body))
	defer r.Body.Close()

	kendaraan := &model.Kendaraan{}
	err := json.Unmarshal(body, kendaraan)
	if err != nil {
		utils.Response(w, utils.Message(false, "Invalid request "+err.Error()))
		return
	}

	resKendaraan, ok := kendaraan.AddKendaraan()
	if !ok {
		utils.Response(w, resKendaraan)
		return
	}

	resLog, _ := json.Marshal(resKendaraan)
	utils.Logging.Println(string(resLog))
	utils.Response(w, resKendaraan)
}

var GetKendaraanId = func(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	id, _ := strconv.Atoi(param["id"])

	res := model.GetKendaraanId(id)
	utils.Logging.Println(res)

	resLog, _ := json.Marshal(res)
	utils.Logging.Println(string(resLog))
	utils.Response(w, res)
}

var UpdateKendaraan = func(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	utils.Logging.Println(string(body))
	defer r.Body.Close()

	kendaraan := &model.Kendaraan{}
	err := json.Unmarshal(body, kendaraan)
	if err != nil {
		utils.Response(w, utils.Message(false, "Invalid request "+err.Error()))
		return
	}

	resKendaraan := kendaraan.UpdateKendaraan()

	resLog, _ := json.Marshal(resKendaraan)
	utils.Logging.Println(string(resLog))
	utils.Response(w, resKendaraan)
}
