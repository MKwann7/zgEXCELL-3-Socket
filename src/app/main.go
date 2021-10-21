package main

import (
	"github.com/MKwann7/zgEXCELL-Socket/src/app/libraries/process"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	"github.com/urfave/negroni"
	"log"
	"net/http"
	"os"
)

var (
	clients = make(map[*websocket.Conn]bool)
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := router()
	middleware := negroni.Classic()
	middleware.UseHandler(router)

	//go process.HandleInboundMessage()

	http.ListenAndServe(":"+os.Getenv("PORT_NUM"), middleware)
}

func router() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	router.
		Methods("GET").
		Path("/health-check").
		HandlerFunc(process.HandleHealthCheck)
	router.
		Methods("GET").
		Path("/socket").
		HandlerFunc(process.HandleConnections)

	return router
}
