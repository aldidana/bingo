package config

import (
  "github.com/julienschmidt/httprouter"
  "github.com/aldidana/bingo/handler"
)

func Router() *httprouter.Router {
  router := httprouter.New()

  catHandler := handler.NewCatHandler(getSession())

  router.GET("/cat", Logger(catHandler.GetAllCats))
  router.GET("/cat/search", Logger(catHandler.GetCatByName))
  router.POST("/cat/add", Logger(catHandler.AddCat))
  router.PUT("/cat/update/:id", Logger(catHandler.UpdateCat))
  router.DELETE("/cat/delete/:id", Logger(catHandler.DeleteCat))

  return router
}
