package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper"
	"github.com/rl404/mal-api/internal/pkg/utils"
)

type top struct {
	mal *malscraper.Malscraper
}

func registerTop(r chi.Router, mal *malscraper.Malscraper) {
	t := top{mal: mal}
	r.Get("/top/anime", t.getTopAnime)
	r.Get("/top/manga", t.getTopManga)
	r.Get("/top/character", t.getTopCharacter)
	r.Get("/top/people", t.getTopPeople)
}

// @summary Get top anime list
// @tags top
// @accept json
// @produce json
// @param type query integer false "Top type (1=airing, 2=upcoming, 3=tv, 4=movie, 5=ova, 6=ona, 7=special, 8=popularity, 9=favorite" enums(0,1,2,3,4,5,6,7,8,9)
// @param page query integer false "Page"
// @success 200 {object} utils.Response{data=[]model.TopAnime}
// @router /top/anime [get]
func (c *top) getTopAnime(w http.ResponseWriter, r *http.Request) {
	t, _ := strconv.Atoi(r.URL.Query().Get("type"))
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := c.mal.GetTopAnime(t, page)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get top manga list
// @tags top
// @accept json
// @produce json
// @param type query integer false "Top type (1=manga, 2=novel, 3=one-shot, 4=doujin, 5=manhwa, 6=manhua, 7=popularity, 8=favorite" enums(0,1,2,3,4,5,6,7,8)
// @param page query integer false "Page"
// @success 200 {object} utils.Response{data=[]model.TopManga}
// @router /top/manga [get]
func (c *top) getTopManga(w http.ResponseWriter, r *http.Request) {
	t, _ := strconv.Atoi(r.URL.Query().Get("type"))
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := c.mal.GetTopManga(t, page)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get top character list
// @tags top
// @accept json
// @produce json
// @param page query integer false "Page"
// @success 200 {object} utils.Response{data=[]model.TopCharacter}
// @router /top/character [get]
func (c *top) getTopCharacter(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := c.mal.GetTopCharacter(page)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get top people list
// @tags top
// @accept json
// @produce json
// @param page query integer false "Page"
// @success 200 {object} utils.Response{data=[]model.TopPeople}
// @router /top/people [get]
func (c *top) getTopPeople(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := c.mal.GetTopPeople(page)
	utils.ResponseWithJSON(w, code, data, err)
}
