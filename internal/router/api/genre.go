package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper"
	"github.com/rl404/mal-api/internal/pkg/utils"
)

type genre struct {
	mal *malscraper.Malscraper
}

func registerGenre(r chi.Router, mal *malscraper.Malscraper) {
	g := genre{mal: mal}
	r.Get("/genres/anime", g.getAnimeGenres)
	r.Get("/genres/manga", g.getMangaGenres)
	r.Get("/genre/{id}/anime", g.getAnimeGenre)
	r.Get("/genre/{id}/manga", g.getMangaGenre)
}

// @summary Get anime genre list
// @tags genre
// @accept json
// @produce json
// @success 200 {object} utils.Response{data=[]model.ItemCount}
// @router /genres/anime [get]
func (c *genre) getAnimeGenres(w http.ResponseWriter, r *http.Request) {
	data, code, err := c.mal.GetAnimeGenres()
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get anime with specific genre
// @tags genre
// @accept json
// @produce json
// @param id path integer true "Anime genre ID"
// @param page query integer false "Page"
// @success 200 {object} utils.Response{data=[]model.AnimeItem}
// @router /genre/{id}/anime [get]
func (c *genre) getAnimeGenre(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := c.mal.GetAnimeWithGenre(id, page)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get manga genre list
// @tags genre
// @accept json
// @produce json
// @success 200 {object} utils.Response{data=[]model.ItemCount}
// @router /genres/manga [get]
func (c *genre) getMangaGenres(w http.ResponseWriter, r *http.Request) {
	data, code, err := c.mal.GetMangaGenres()
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get manga with specific genre
// @tags genre
// @accept json
// @produce json
// @param id path integer true "Manga genre ID"
// @param page query integer false "Page"
// @success 200 {object} utils.Response{data=[]model.MangaItem}
// @router /genre/{id}/manga [get]
func (c *genre) getMangaGenre(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := c.mal.GetMangaWithGenre(id, page)
	utils.ResponseWithJSON(w, code, data, err)
}
