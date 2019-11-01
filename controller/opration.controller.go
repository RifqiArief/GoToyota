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

var AddOpration = func(w http.ResponseWriter, r *http.Request) {

	body, _ := ioutil.ReadAll(r.Body)
	utils.Logging.Println(string(body))
	defer r.Body.Close()

	var opr []*model.Oprasional
	err := json.Unmarshal(body, &opr)
	if err != nil {
		utils.Response(w, utils.Message(false, err.Error()))
		return
	}

	res, err := model.AddOpration(opr)
	if err != nil {
		utils.Response(w, utils.Message(false, err.Error()))
		return
	}
	resLog, _ := json.Marshal(res)
	utils.Logging.Println(string(resLog))
	utils.Response(w, res)
}

var GetOpration = func(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)
	id, _ := strconv.Atoi(param["id"])

	res := model.GetOpration(id)
	utils.Logging.Println(res)

	resLog, _ := json.Marshal(res)
	utils.Logging.Println(string(resLog))
	utils.Response(w, res)
}
