package main

import (
	"net/http"

	"github.com/webAPi/controllers"
)

func main() {
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)
}
