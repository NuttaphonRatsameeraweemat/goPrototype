package main

import (
	"log"
	"net/http"
	env "redis-cache-api/Helper/ApiHelper/ConfigEnv"
	setRouter "redis-cache-api/Helper/ApiHelper/Routes"
)

func main() {
	env.InitEnv()

	router := setRouter.NewRouter()

	log.Fatal(
		// start on port 8083 by default
		http.ListenAndServe(":80", router),
	)

}
