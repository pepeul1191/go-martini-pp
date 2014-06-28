package main

import "github.com/go-martini/martini"
import "github.com/martini-contrib/sessions"

import "github.com/IoSync/martini-boilerplate/auth"

func main() {
  m := martini.Classic()
  m.Use(martini.Logger())
  m.Use(martini.Static("/public"))

  store := sessions.NewCookieStore([]byte("secret123"))
  m.Use(sessions.Sessions("monitor_session", store))

  //
  // login/logout handlers
  //
  m.Post("/login", auth.Login)
  m.Post("/logout", auth.Logout)

  m.Run()
}
