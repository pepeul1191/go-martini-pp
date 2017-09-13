package main

import "github.com/go-martini/martini"
import "github.com/martini-contrib/sessions"

import "github.com/pepe/accesos/auth"
import "github.com/pepe/accesos/middlewares"
import "github.com/pepe/accesos/handlers"

func main() {
    m := martini.Classic()
    m.Use(martini.Logger())
    m.Use(martini.Static("/public"))

    m.Handlers(
        app_headers.Headers,
    )

    store := sessions.NewCookieStore([]byte("secret123"))
    m.Use(sessions.Sessions("monitor_session", store))

    m.Get("/estado_usuario/listar", handlers.EstadoUsuarioListar)
    m.Get("/usuario", handlers.UsuarioIndex)

    m.Post("/login", auth.Login)
    m.Post("/logout", auth.Logout)

    m.Run()
}
