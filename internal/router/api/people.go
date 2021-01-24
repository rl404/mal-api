package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper"
	"github.com/rl404/mal-api/internal/pkg/utils"
)

type people struct {
	mal *malscraper.Malscraper
}

func registerPeople(r chi.Router, mal *malscraper.Malscraper) {
	p := people{mal: mal}
	r.Get("/people/{id}", p.getPeople)
	r.Get("/people/{id}/characters", p.getPeopleCharacter)
	r.Get("/people/{id}/staff", p.getPeopleStaff)
	r.Get("/people/{id}/manga", p.getPeopleManga)
	r.Get("/people/{id}/news", p.getPeopleNews)
	r.Get("/people/{id}/article", p.getPeopleArticle)
	r.Get("/people/{id}/pictures", p.getPeoplePicture)
}

// @summary Get people details
// @tags people
// @accept json
// @produce json
// @param id path integer true "People ID"
// @success 200 {object} utils.Response{data=model.People}
// @router /people/{id} [get]
func (c *people) getPeople(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := c.mal.GetPeople(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get people anime character role list
// @tags people
// @accept json
// @produce json
// @param id path integer true "People ID"
// @success 200 {object} utils.Response{data=[]model.PeopleCharacter}
// @router /people/{id}/characters [get]
func (c *people) getPeopleCharacter(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := c.mal.GetPeopleCharacter(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get people anime staff role list
// @tags people
// @accept json
// @produce json
// @param id path integer true "People ID"
// @success 200 {object} utils.Response{data=[]model.Role}
// @router /people/{id}/staff [get]
func (c *people) getPeopleStaff(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := c.mal.GetPeopleStaff(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get people published manga list
// @tags people
// @accept json
// @produce json
// @param id path integer true "People ID"
// @success 200 {object} utils.Response{data=[]model.Role}
// @router /people/{id}/manga [get]
func (c *people) getPeopleManga(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := c.mal.GetPeopleManga(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get people news list
// @tags people
// @accept json
// @produce json
// @param id path integer true "People ID"
// @success 200 {object} utils.Response{data=[]model.NewsItem}
// @router /people/{id}/news [get]
func (c *people) getPeopleNews(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := c.mal.GetPeopleNews(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get people featured article list
// @tags people
// @accept json
// @produce json
// @param id path integer true "People ID"
// @success 200 {object} utils.Response{data=[]model.ArticleItem}
// @router /people/{id}/article [get]
func (c *people) getPeopleArticle(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := c.mal.GetPeopleArticle(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get people picture list
// @tags people
// @accept json
// @produce json
// @param id path integer true "People ID"
// @success 200 {object} utils.Response{data=[]string}
// @router /people/{id}/pictures [get]
func (c *people) getPeoplePicture(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := c.mal.GetPeoplePicture(id)
	utils.ResponseWithJSON(w, code, data, err)
}
