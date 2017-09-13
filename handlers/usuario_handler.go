package usuario_handler

import "net/http"

import "github.com/martini-contrib/sessions"

func Index(w http.ResponseWriter, r *http.Request, session sessions.Session) (int, string) {
	r.Header.Set("hola","mundo")
    var response string = "usuario/index XD"
	return 200, response
}