package router

import "net/http"

func APIMux() *http.ServeMux {
	apiMux := http.NewServeMux()
	return apiMux
}
