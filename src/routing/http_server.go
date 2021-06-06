package routing

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func StartHttpServer() {
	port := 8080
	publicRouter := mux.NewRouter()
	SetApiRoutes(publicRouter)
	if err := http.ListenAndServe(fmt.Sprint(":", port), publicRouter); err != nil {
		os.Exit(1)
	}
}
