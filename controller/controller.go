package controller

import (
	"deliveroo/db"
	"deliveroo/model"
	"deliveroo/services"
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func GetGroupTypeClientInfo(response http.ResponseWriter, resquest *http.Request) {
	services.GetGroupType(response, resquest)
}

func GetRestaurantClientInfo(response http.ResponseWriter, resquest *http.Request) {
	services.GetRestaurantInfo(response, resquest)
}

func UserLogin(response http.ResponseWriter, resquest *http.Request) {
	userName := resquest.Header.Get("Username")
	log.Info("Start get login tocken")
	token, err := services.GetLoginToken(userName)
	if err != nil {
		log.Error(err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Info("Send to client login info: ", userName)

	groupName, err := db.GetUserGroup(userName)
	if err != nil {
		log.Error(err)
		response.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseInfo := model.UserLoginInfo{UserName: userName, GroupName: groupName, Token: token}
	response.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(response).Encode(responseInfo)
	if err != nil {
		http.Error(response, err.Error(), 500)
		log.Error(err, "endcode json")
	}
}
