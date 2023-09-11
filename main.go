package main

import (
	"fmt"
	"main/routes"
	"net/http"
)

func main() {
	fmt.Println("Server : https://localhost:8080")
	http.ListenAndServe(":8080", routes.Routes())
}
