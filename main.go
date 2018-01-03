package main


import (
	"fmt"
	"net/http"
	"github.com/mike001005/emailApi/pkg/api"

	"os"
)

func main() {

	os.Setenv("MG_URL","https://api.mailgun.net/v3")

	fmt.Println("Server starting")
	http.ListenAndServe(":8081", api.Handlers())

}

