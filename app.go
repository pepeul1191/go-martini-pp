package main

import "github.com/go-martini/martini"
import "github.com/martini-contrib/sessions"

import "github.com/pepe/accesos/auth"
import "github.com/pepe/accesos/handlers"

func main() {
  m := martini.Classic()
  m.Use(martini.Logger())
  m.Use(martini.Static("/public"))

  store := sessions.NewCookieStore([]byte("secret123"))
  m.Use(sessions.Sessions("monitor_session", store))
  //
  // login/logout handlers
  //
  m.Get("/usuario", usuario_handler.Index)
  m.Post("/login", auth.Login)
  m.Post("/logout", auth.Logout)

  m.Run()
}
