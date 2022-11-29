package routers

import (
	"deliveroo/controller"
	"deliveroo/middlewares"
	"net/http"

	mux "github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	//api get data
	subRouterGetData := router.PathPrefix("/deilveroo").Subrouter()
	subRouterGetData.HandleFunc("/grouptype", controller.GetGroupTypeClientInfo).Methods(http.MethodGet)
	subRouterGetData.HandleFunc("/restaurantinfo", controller.GetRestaurantClientInfo).Methods(http.MethodGet)
	subRouterGetData.Use(middlewares.AuthenTokenDev)

	//api login

	subRouterLogin := router.PathPrefix("/user").Subrouter()
	subRouterLogin.HandleFunc("/login", controller.UserLogin).Methods(http.MethodPost)
	subRouterLogin.Use(middlewares.AuthenUserLoginMiddleware)

	return router
}
