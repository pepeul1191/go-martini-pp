package usuario_handler

import "net/http"
import "github.com/martini-contrib/sessions"

func Index(res http.ResponseWriter, req *http.Request, session sessions.Session) (int, string) {
	res.Header().Set("Content-Type", "text/plain; charset=utf-8")
    var rpta string = "usuario/index XD"
	return 200, rpta
}