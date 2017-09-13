package handlers

import "net/http"
import "github.com/martini-contrib/sessions"
//import "github.com/pepe/accesos/models"
import "github.com/pepe/accesos/config"
import "fmt"

type Estado struct {
    ID      uint `gorm:"primary_key"` 
    Nombre   string
}

func (Estado) TableName() string {
    return "estado_usuarios"
}

func EstadoUsuarioListar(res http.ResponseWriter, req *http.Request, session sessions.Session) (int, string) {
	db := database.DB()
	estado := Estado{}
	rs := db.Find(&estado)
	fmt.Printf( rs.Count + "")
    var rpta string = "EstadoUsuarioListar"
	return 200, rpta
}