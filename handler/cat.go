package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/aldidana/bingo/model"
	"github.com/aldidana/bingo/render" //utils consist of render
	"github.com/julienschmidt/httprouter"
	r "gopkg.in/dancannon/gorethink.v2"
)

const (
	xml = "xml"
)

type (
	//CatHandler struct
	CatHandler struct {
		session *r.Session
	}
)

//NewCatHandler CathHandler session
func NewCatHandler(s *r.Session) *CatHandler {
	return &CatHandler{s}
}

//GetAllCats return all cats
func (ch CatHandler) GetAllCats(res http.ResponseWriter, request *http.Request, p httprouter.Params) error {
	q := request.URL.Query()

	result, err := r.Table("cat").Run(ch.session)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}

	defer result.Close()

	var cats []model.Cat
	err = result.All(&cats)
	if err != nil {
		fmt.Printf("Error scanning database result: %s", err)
		return err
	}

	if q.Get("render") == xml {
		render.ToXML(res, request, p, cats)
		return nil
	}
	render.ToJSON(res, request, p, cats)
	return nil
}

//GetCatByName return cat by name
func (ch CatHandler) GetCatByName(res http.ResponseWriter, request *http.Request, p httprouter.Params) error {
	q := request.URL.Query()
	name := q.Get("name")

	result, err := r.Table("cat").Filter(map[string]interface{}{
		"name": name,
	}).Run(ch.session)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}

	defer result.Close()

	var cats []model.Cat
	err = result.All(&cats)
	if err != nil {
		fmt.Printf("Error scanning database result: %s", err)
		return err
	}

	if q.Get("render") == xml {
		render.ToXML(res, request, p, cats)
		return nil
	}
	render.ToJSON(res, request, p, cats)
	return nil
}

//AddCat insert cat
func (ch CatHandler) AddCat(res http.ResponseWriter, request *http.Request, p httprouter.Params) error {
	q := request.URL.Query()
	newCat := new(model.Cat)
	json.NewDecoder(request.Body).Decode(newCat)

	result, err := r.Table("cat").Insert(newCat).RunWrite(ch.session)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}

	if q.Get("render") == xml {
		render.ToXML(res, request, p, result)
		return nil
	}
	render.ToJSON(res, request, p, result)
	return nil
}

//UpdateCat update cat
func (ch CatHandler) UpdateCat(res http.ResponseWriter, request *http.Request, p httprouter.Params) error {
	q := request.URL.Query()
	updateCat := new(model.Cat)
	catID := p.ByName("id")
	json.NewDecoder(request.Body).Decode(updateCat)

	result, err := r.Table("cat").Get(catID).Update(map[string]interface{}{
		"name": updateCat.Name,
		"type": updateCat.Type,
	}).RunWrite(ch.session)

	if err != nil {
		log.Fatalln(err.Error())
		return err
	}

	if q.Get("render") == xml {
		render.ToXML(res, request, p, result)
		return nil
	}
	render.ToJSON(res, request, p, result)
	return nil
}

//DeleteCat delete cat
func (ch CatHandler) DeleteCat(res http.ResponseWriter, request *http.Request, p httprouter.Params) error {
	q := request.URL.Query()
	catID := p.ByName("id")

	result, err := r.Table("cat").Get(catID).Delete().RunWrite(ch.session)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}

	if q.Get("render") == xml {
		render.ToXML(res, request, p, result)
		return nil
	}
	render.ToJSON(res, request, p, result)
	return nil
}
