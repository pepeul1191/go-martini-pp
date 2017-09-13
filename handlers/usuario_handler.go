package handlers

import "net/http"
import "github.com/martini-contrib/sessions"

func UsuarioIndex(res http.ResponseWriter, req *http.Request, session sessions.Session) (int, string) {
    var rpta string = "UsuarioIndex"
	return 200, rpta
}