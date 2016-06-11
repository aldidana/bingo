package config

import (
  "log"
  r "gopkg.in/dancannon/gorethink.v2"
)

func getSession() *r.Session {
  session, err := r.Connect(r.ConnectOpts{
    Address: "localhost:28015",
    Database: "bingo",
  })

  if err != nil {
    log.Fatalln(err.Error())
  }

  return session
}
