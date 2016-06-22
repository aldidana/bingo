package config

import (
	"github.com/aldidana/bingo/handler"
	"github.com/julienschmidt/httprouter"
)

//Router API Router
func Router() *httprouter.Router {
	router := httprouter.New()

	catHandler := handler.NewCatHandler(getSession())

	// router.GET("/cat", Logger(catHandler.GetAllCats))
	router.GET("/cat", Middleware(Logger, catHandler.GetAllCats))
	router.GET("/cat/search", Middleware(catHandler.GetCatByName))
	router.POST("/cat/add", Middleware(catHandler.AddCat))
	router.PUT("/cat/update/:id", Middleware(catHandler.UpdateCat))
	router.DELETE("/cat/delete/:id", Middleware(catHandler.DeleteCat))

	return router
}
