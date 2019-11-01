package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/GoToyota/model"
	"github.com/GoToyota/utils"
)

var RegisterUser = func(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	utils.Logging.Println(string(body))
	defer r.Body.Close()

	user := &model.User{}
	err := json.Unmarshal(body, user)
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

	opration := &model.Oprasional{}
	err = json.Unmarshal(body, opration)
	if err != nil {
		utils.Response(w, utils.Message(false, "Invalid request"))
		return
	}

	resUser, ok := user.AddUser()
	if !ok {
		utils.Response(w, resUser)
		return
	}

	// resLocation, ok := location.AddLocation(data)
	// if !ok {
	// 	utils.Response(w, resLocation)
	// 	return
	// }

	resLog, _ := json.Marshal(resUser)
	utils.Logging.Println(string(resLog))
	utils.Response(w, resUser)
}

var LoginUser = func(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	utils.Logging.Println(string(body))
	defer r.Body.Close()

	user := &model.User{}

	err := json.Unmarshal(body, user)
	if err != nil {
		utils.Response(w, utils.Message(false, "Invalid request"))
		return
	}

	res, ok := model.LoginUser(user.Email, user.Password)
	utils.Logging.Println(res)
	if !ok {
		utils.Response(w, res)
		return
	}

	resLog, _ := json.Marshal(res)
	utils.Logging.Println(string(resLog))
	utils.Response(w, res)
}

var GetAllUser = func(w http.ResponseWriter, r *http.Request) {

	res := model.GetAllUser()
	utils.Logging.Println(res)

	resLog, _ := json.Marshal(res)
	utils.Logging.Println(string(resLog))
	utils.Response(w, res)
}

var EditProfile = func(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	utils.Logging.Println(string(body))
	defer r.Body.Close()

	user := &model.User{}
	err := json.Unmarshal(body, user)
	if err != nil {
		utils.Response(w, utils.Message(false, "Invalid request "+err.Error()))
		return
	}

	resUser := user.UpdateUser()

	resLog, _ := json.Marshal(resUser)
	utils.Logging.Println(string(resLog))
	utils.Response(w, resUser)
}

var ChangePassword = func(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	utils.Logging.Println(string(body))
	defer r.Body.Close()

	user := &model.ChangePassword{}
	err := json.Unmarshal(body, user)
	if err != nil {
		utils.Response(w, utils.Message(false, "Invalid request "+err.Error()))
		return
	}

	resUser := user.ChangePassword()

	resLog, _ := json.Marshal(resUser)
	utils.Logging.Println(string(resLog))
	utils.Response(w, resUser)
}
