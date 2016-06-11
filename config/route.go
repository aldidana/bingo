package config

import (
  "github.com/julienschmidt/httprouter"
  "github.com/aldidana/bingo/handler"
)

func Router() *httprouter.Router {
  router := httprouter.New()

  catHandler := handler.NewCatHandler(getSession())

  router.GET("/cat", catHandler.GetAllCats)
  router.GET("/cat/search", catHandler.GetCatByName)
  router.POST("/cat/add", catHandler.AddCat)
  router.PUT("/cat/update/:id", catHandler.UpdateCat)
  router.DELETE("/cat/delete/:id", catHandler.DeleteCat)

  return router
}
