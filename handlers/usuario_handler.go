package usuario_handler

import "net/http"
import "github.com/martini-contrib/sessions"

func Index(res http.ResponseWriter, req *http.Request, session sessions.Session) (int, string) {
    var rpta string = "hola mundo"
	return 200, rpta
}