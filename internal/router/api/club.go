package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper"
	"github.com/rl404/mal-api/internal/pkg/utils"
)

type club struct {
	mal *malscraper.Malscraper
}

func registerClub(r chi.Router, mal *malscraper.Malscraper) {
	c := club{mal: mal}
	r.Get("/clubs", c.getClubs)
	r.Get("/club/{id}", c.getClub)
	r.Get("/club/{id}/members", c.getClubMember)
	r.Get("/club/{id}/pictures", c.getClubPicture)
	r.Get("/club/{id}/related", c.getClubRelated)
}

// @summary Get club list
// @tags club
// @accept json
// @produce json
// @param page query integer false "Page"
// @success 200 {object} utils.Response{data=[]model.ClubSearch}
// @router /clubs [get]
func (c *club) getClubs(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := c.mal.GetClubs(page)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get club details
// @tags club
// @accept json
// @produce json
// @param id path integer true "Club ID"
// @success 200 {object} utils.Response{data=model.Club}
// @router /club/{id} [get]
func (c *club) getClub(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := c.mal.GetClub(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get club member list
// @tags club
// @accept json
// @produce json
// @param id path integer true "Club ID"
// @param page query integer false "Page"
// @success 200 {object} utils.Response{data=[]model.ClubMember}
// @router /club/{id}/members [get]
func (c *club) getClubMember(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := c.mal.GetClubMember(id, page)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get club picture list
// @tags club
// @accept json
// @produce json
// @param id path integer true "Club ID"
// @success 200 {object} utils.Response{data=[]string}
// @router /club/{id}/pictures [get]
func (c *club) getClubPicture(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := c.mal.GetClubPicture(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get club related anime/manga/character
// @tags club
// @accept json
// @produce json
// @param id path integer true "Club ID"
// @success 200 {object} utils.Response{data=model.ClubRelated}
// @router /club/{id}/related [get]
func (c *club) getClubRelated(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := c.mal.GetClubRelated(id)
	utils.ResponseWithJSON(w, code, data, err)
}
