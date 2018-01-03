package main


import (
	"fmt"
	"net/http"
	"github.com/dollarshaveclub/email/pkg/api"

)

func main() {

	fmt.Println("Server starting")
	http.ListenAndServe(":8081", api.Handlers())

}

