package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper"
	"github.com/rl404/mal-api/internal/pkg/utils"
)

type char struct {
	mal *malscraper.Malscraper
}

func registerCharacter(r chi.Router, mal *malscraper.Malscraper) {
	c := char{mal: mal}
	r.Get("/character/{id}", c.getCharacter)
	r.Get("/character/{id}/anime", c.getCharacterAnime)
	r.Get("/character/{id}/manga", c.getCharacterManga)
	r.Get("/character/{id}/article", c.getCharacterArticle)
	r.Get("/character/{id}/pictures", c.getCharacterPicture)
	r.Get("/character/{id}/clubs", c.getCharacterClub)
}

// @summary Get character details
// @tags character
// @accept json
// @produce json
// @param id path integer true "Character ID"
// @success 200 {object} utils.Response{data=model.Character}
// @router /character/{id} [get]
func (c *char) getCharacter(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := c.mal.GetCharacter(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get character animeography list
// @tags character
// @accept json
// @produce json
// @param id path integer true "Character ID"
// @success 200 {object} utils.Response{data=[]model.Role}
// @router /character/{id}/anime [get]
func (c *char) getCharacterAnime(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := c.mal.GetCharacterAnime(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get character mangaography list
// @tags character
// @accept json
// @produce json
// @param id path integer true "Character ID"
// @success 200 {object} utils.Response{data=[]model.Role}
// @router /character/{id}/manga [get]
func (c *char) getCharacterManga(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := c.mal.GetCharacterManga(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get character article featured list
// @tags character
// @accept json
// @produce json
// @param id path integer true "Character ID"
// @success 200 {object} utils.Response{data=[]model.ArticleItem}
// @router /character/{id}/article [get]
func (c *char) getCharacterArticle(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := c.mal.GetCharacterArticle(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get character picture list
// @tags character
// @accept json
// @produce json
// @param id path integer true "Character ID"
// @success 200 {object} utils.Response{data=[]string}
// @router /character/{id}/pictures [get]
func (c *char) getCharacterPicture(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := c.mal.GetCharacterPicture(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get character club list
// @tags character
// @accept json
// @produce json
// @param id path integer true "Character ID"
// @success 200 {object} utils.Response{data=[]model.ClubItem}
// @router /character/{id}/clubs [get]
func (c *char) getCharacterClub(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := c.mal.GetCharacterClub(id)
	utils.ResponseWithJSON(w, code, data, err)
}
