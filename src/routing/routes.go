package routing

import (
	"github.com/gorilla/mux"

	"livechat.com/lc-roler/handlers"
)

func SetApiRoutes(router *mux.Router) {
	router.HandleFunc("/install", handlers.HandleInstall)
}
