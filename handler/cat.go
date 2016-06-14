package handler

import (
  "encoding/json"
  "fmt"
  "net/http"
  "log"
  "github.com/aldidana/bingo/model"
  r "gopkg.in/dancannon/gorethink.v2"
  "github.com/julienschmidt/httprouter"
)

type (
  CatHandler struct {
    session *r.Session
  }
)

func NewCatHandler (s *r.Session) *CatHandler {
  return &CatHandler{s}
}

func (ch CatHandler) GetAllCats(res http.ResponseWriter, request *http.Request, p httprouter.Params) {
  result, err := r.Table("cat").Run(ch.session)
  if err != nil {
    log.Fatalln(err.Error())
    return
  }

  defer result.Close()

  var cats []model.Cat
  err = result.All(&cats)
  if err != nil {
      fmt.Printf("Error scanning database result: %s", err)
      return
  }

  res.Header().Set("Content-Type", "application/json; charset=UTF-8")
  res.WriteHeader(http.StatusOK)

  if err := json.NewEncoder(res).Encode(cats); err != nil {
    res.Header().Set("Content-Type", "application/json; charset=UTF-8")
		res.WriteHeader(http.StatusBadRequest)
		return
  }
}

func (ch CatHandler) GetCatByName(res http.ResponseWriter, request *http.Request, p httprouter.Params) {
  q := request.URL.Query()
  name := q.Get("name")
  result, err := r.Table("cat").Filter(map[string]interface{}{
      "name": name,
    }).Run(ch.session)
  if err != nil {
    log.Fatalln(err.Error())
    return
  }

  defer result.Close()

  var cats []model.Cat
  err = result.All(&cats)
  if err != nil {
      fmt.Printf("Error scanning database result: %s", err)
      return
  }

  res.Header().Set("Content-Type", "application/json; charset=UTF-8")
  res.WriteHeader(http.StatusOK)

  if err := json.NewEncoder(res).Encode(cats); err != nil {
    res.Header().Set("Content-Type", "application/json; charset=UTF-8")
		res.WriteHeader(http.StatusBadRequest)
		return
  }
}

func (ch CatHandler) AddCat(res http.ResponseWriter, request *http.Request, p httprouter.Params) {
  newCat := new(model.Cat)
  json.NewDecoder(request.Body).Decode(newCat)

  result, err := r.Table("cat").Insert(newCat).RunWrite(ch.session)
  if err != nil {
    log.Fatalln(err.Error())
  }

  res.Header().Set("Content-Type", "application/json; charset=UTF-8")
  res.WriteHeader(http.StatusCreated)

  if err := json.NewEncoder(res).Encode(result); err != nil {
    res.Header().Set("Content-Type", "application/json; charset=UTF-8")
		res.WriteHeader(http.StatusBadRequest)
		return
  }

}

func (ch CatHandler) UpdateCat(res http.ResponseWriter, request *http.Request, p httprouter.Params) {
  updateCat := new(model.Cat)
  catId := p.ByName("id")
  json.NewDecoder(request.Body).Decode(updateCat)

  result, err := r.Table("cat").Get(catId).Update(map[string]interface{}{
    "name": updateCat.Name,
    "type": updateCat.Type,
    }).RunWrite(ch.session)

  if err != nil {
    log.Fatalln(err.Error())
  }

  res.Header().Set("Content-Type", "application/json; charset=UTF-8")
  res.WriteHeader(http.StatusOK)

  if err := json.NewEncoder(res).Encode(result); err != nil {
    res.Header().Set("Content-Type", "application/json; charset=UTF-8")
		res.WriteHeader(http.StatusBadRequest)
		return
  }
}


func (ch CatHandler) DeleteCat(res http.ResponseWriter, request *http.Request, p httprouter.Params) {
  catId := p.ByName("id")

  result, err := r.Table("cat").Get(catId).Delete().RunWrite(ch.session)
  if err != nil {
    log.Fatalln(err.Error())
  }

  res.Header().Set("Content-Type", "application/json; charset=UTF-8")
  res.WriteHeader(http.StatusOK)

  if err := json.NewEncoder(res).Encode(result); err != nil {
    res.Header().Set("Content-Type", "application/json; charset=UTF-8")
		res.WriteHeader(http.StatusBadRequest)
		return
  }
}
