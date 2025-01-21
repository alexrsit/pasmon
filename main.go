package main

import (
	"fmt"
	"net/http"
	"pasmon/api"
	"pasmon/routes"
)

func main() {
	client, err := api.InitApiClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	routes.InitRoutes(client)

	http.ListenAndServe(":3000", nil)
	fmt.Println("http:localhost:3000")
}
