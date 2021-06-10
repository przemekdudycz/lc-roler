package routing

import (
	"github.com/gorilla/mux"

	"livechat.com/lc-roler/handlers"
)

func SetApiRoutes(router *mux.Router) {
	router.HandleFunc("/install", handlers.HandleInstall)
	router.HandleFunc("/newchat", handlers.HandleNewChat)
	router.HandleFunc("/rmpostback", handlers.HandleRMPostback)
	router.HandleFunc("/newevent", handlers.HandleEvent)
}
