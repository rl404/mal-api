package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper"
	"github.com/rl404/mal-api/internal/pkg/utils"
)

type recommendation struct {
	mal *malscraper.Malscraper
}

func registerRecommendation(r chi.Router, mal *malscraper.Malscraper) {
	re := recommendation{mal: mal}
	r.Get("/recommendation/anime/{id1}/{id2}", re.getAnimeRecommendation)
	r.Get("/recommendation/manga/{id1}/{id2}", re.getMangaRecommendation)
	r.Get("/recommendations/anime", re.getAnimeRecommendations)
	r.Get("/recommendations/manga", re.getMangaRecommendations)
}

// @summary Get anime recommendation from specific anime.
// @tags recommendation
// @accept json
// @produce json
// @param id1 path integer true "Anime ID"
// @param id2 path integer true "Anime ID"
// @success 200 {object} utils.Response{data=model.Recommendation}
// @router /recommendation/anime/{id1}/{id2} [get]
func (c *recommendation) getAnimeRecommendation(w http.ResponseWriter, r *http.Request) {
	id1, _ := strconv.Atoi(chi.URLParam(r, "id1"))
	id2, _ := strconv.Atoi(chi.URLParam(r, "id2"))
	data, code, err := c.mal.GetRecommendation(malscraper.AnimeType, id1, id2)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get manga recommendation from specific manga.
// @tags recommendation
// @accept json
// @produce json
// @param id1 path integer true "Manga ID"
// @param id2 path integer true "Manga ID"
// @success 200 {object} utils.Response{data=model.Recommendation}
// @router /recommendation/manga/{id1}/{id2} [get]
func (c *recommendation) getMangaRecommendation(w http.ResponseWriter, r *http.Request) {
	id1, _ := strconv.Atoi(chi.URLParam(r, "id1"))
	id2, _ := strconv.Atoi(chi.URLParam(r, "id2"))
	data, code, err := c.mal.GetRecommendation(malscraper.MangaType, id1, id2)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get anime recommendation list.
// @tags recommendation
// @accept json
// @produce json
// @param page query integer false "Page"
// @success 200 {object} utils.Response{data=[]model.Recommendation}
// @router /recommendations/anime [get]
func (c *recommendation) getAnimeRecommendations(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := c.mal.GetRecommendations(malscraper.AnimeType, page)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get manga recommendation list.
// @tags recommendation
// @accept json
// @produce json
// @param page query integer false "Page"
// @success 200 {object} utils.Response{data=[]model.Recommendation}
// @router /recommendations/manga [get]
func (c *recommendation) getMangaRecommendations(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := c.mal.GetRecommendations(malscraper.MangaType, page)
	utils.ResponseWithJSON(w, code, data, err)
}
