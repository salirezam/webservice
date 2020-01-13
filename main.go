package main

import (
	"github.com/salirezam/webservice/controllers"
	"net/http"
)

func main() {
	controllers.RegisterController()
	http.ListenAndServe(":3000", nil)
}
