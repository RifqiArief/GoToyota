package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/GoToyota/object"

	"github.com/GoToyota/model"
	"github.com/GoToyota/utils"
)

type resReg struct {
	Code     bool                   `json:"code"`
	Message  string                 `json:"message"`
	Response object.RegisterBengkel `json:"response"`
}
type resLogin struct {
	Code     bool                `json:"code"`
	Message  string              `json:"message"`
	Response []object.Oprasional `json:"response"`
}

var RegisterBengkel = func(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	utils.Logging.Println(string(body))
	defer r.Body.Close()

	bengkel := &model.Bengkel{}
	err := json.Unmarshal(body, bengkel)
	if err != nil {
		utils.Response(w, utils.Message(false, "Invalid request"))
		return
	}

	location := &model.Lokasi{}
	err = json.Unmarshal(body, location)
	if err != nil {
		utils.Response(w, utils.Message(false, "Invalid request"))
		return
	}

	resBengkel, objBengkel := bengkel.AddBengkel()
	if objBengkel == nil {
		utils.Response(w, resBengkel)
		return
	}

	resLocation, objLocation := location.AddLocation(objBengkel.IdBengkel)
	if objLocation == nil {
		utils.Response(w, resLocation)
		return
	}

	loc := object.Lokasi{
		Alamat:    objLocation.Alamat,
		Kota:      objLocation.Kota,
		Provinsi:  objLocation.Provinsi,
		Longitude: objLocation.Longitude,
		Latitude:  objLocation.Latitude,
	}

	data := object.RegisterBengkel{
		IdBengkel: objBengkel.IdBengkel,
		Nama:      objBengkel.Nama,
		Email:     objBengkel.Email,
		Telepon:   objBengkel.Telepon,
		Gambar:    objBengkel.Gambar,
		Lokasi:    loc,
	}

	res := &resReg{
		Code:     true,
		Message:  "Succes registration",
		Response: data,
	}

	resLog, _ := json.Marshal(resLocation)
	utils.Logging.Println(string(resLog))
	w.Header().Add("Content-Type", "aplication/json")
	json.NewEncoder(w).Encode(res)
}

var LoginBengkel = func(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	utils.Logging.Println(string(body))
	defer r.Body.Close()

	bengkel := &model.Bengkel{}

	err := json.Unmarshal(body, bengkel)
	if err != nil {
		utils.Response(w, utils.Message(false, "Invalid request"))
		return
	}

	res, objBengkel := model.LoginBengkel(bengkel.Email, bengkel.Password)
	utils.Logging.Println(res)
	if res != nil {
		utils.Response(w, res)
		return
	}

	res, objLocation := model.GetLocation(objBengkel.IdBengkel)
	if res != nil {
		utils.Response(w, res)
		return
	}
	res, objOpration := model.GetOpration(objBengkel.IdBengkel)
	utils.Logging.Println(res)
	if res != nil {
		utils.Response(w, res)
		return
	}

	data := object.LoginBengkel{
		IdBengkel:  objBengkel.IdBengkel,
		Nama:       objBengkel.Nama,
		Email:      objBengkel.Email,
		Telepon:    objBengkel.Telepon,
		Gambar:     objBengkel.Gambar,
		Token:      objBengkel.Token,
		IdRole:     objBengkel.IdRole,
		Lokasi:     objLocation,
		Oprasional: objOpration,
	}

	response := utils.Message(true, "Success")
	response["response"] = data

	resLog, _ := json.Marshal(data)
	utils.Logging.Println(string(resLog))
	utils.Response(w, response)
}

var GetBengkelKota = func(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	kota := param["kota"]

	res := model.GetBengkelKota(kota)
	utils.Logging.Println(res)

	resLog, _ := json.Marshal(res)
	utils.Logging.Println(string(resLog))
	utils.Response(w, res)
}

var GetBengkelProvinsi = func(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	provinsi := param["provinsi"]

	res := model.GetBengkelProvinsi(provinsi)
	utils.Logging.Println(res)

	resLog, _ := json.Marshal(res)
	utils.Logging.Println(string(resLog))
	utils.Response(w, res)
}

var GetBengkelId = func(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	id, _ := strconv.Atoi(param["id"])

	res := model.GetBengkelId(id)
	utils.Logging.Println(res)

	resLog, _ := json.Marshal(res)
	utils.Logging.Println(string(resLog))
	utils.Response(w, res)
}

var GetAllBengkel = func(w http.ResponseWriter, r *http.Request) {

	res := model.GetAllBengkel()
	utils.Logging.Println(res)

	resLog, _ := json.Marshal(res)
	utils.Logging.Println(string(resLog))
	utils.Response(w, res)
}

var EditBengkel = func(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	utils.Logging.Println(string(body))
	defer r.Body.Close()

	bengkel := &model.Bengkel{}
	err := json.Unmarshal(body, bengkel)
	if err != nil {
		utils.Response(w, utils.Message(false, "Invalid request "+err.Error()))
		return
	}

	resBengkel := bengkel.EditBengkel()

	resLog, _ := json.Marshal(resBengkel)
	utils.Logging.Println(string(resLog))
	utils.Response(w, resBengkel)
}

var EditLocation = func(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	utils.Logging.Println(string(body))
	defer r.Body.Close()

	loc := &model.Lokasi{}
	err := json.Unmarshal(body, loc)
	if err != nil {
		utils.Response(w, utils.Message(false, "Invalid request "+err.Error()))
		return
	}

	utils.Logging.Println(loc.IdBengkel)

	resLoc := loc.EditLocation()

	resLog, _ := json.Marshal(resLoc)
	utils.Logging.Println(string(resLog))
	utils.Response(w, resLoc)
}
