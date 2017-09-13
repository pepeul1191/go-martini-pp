package app_headers

import "net/http"

func Headers(res http.ResponseWriter, req *http.Request) () {
	res.Header().Set("Server", "Ubuntu")
	res.Header().Set("X-Powered-By", "Go")
	res.Header().Set("Content-Type", "text/html; charset=utf-8")
}