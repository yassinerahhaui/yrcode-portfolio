package router

import (
	"net/http"

	"yrcode/controllers"
)

func WebMux() *http.ServeMux {
	webMux := http.NewServeMux()
	webMux.HandleFunc("/", controllers.HomePage)
	return webMux
}
