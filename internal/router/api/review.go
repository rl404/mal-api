package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper"
	"github.com/rl404/mal-api/internal/pkg/utils"
)

type review struct {
	mal *malscraper.Malscraper
}

func registerReview(r chi.Router, mal *malscraper.Malscraper) {
	re := review{mal: mal}
	r.Get("/review/{id}", re.getReview)
	r.Get("/reviews/anime", re.getAnimeReviews)
	r.Get("/reviews/manga", re.getMangaReviews)
	r.Get("/reviews/best", re.getBestReviews)
}

// @summary Get review details
// @tags review
// @accept json
// @produce json
// @param id path integer true "Review ID"
// @success 200 {object} utils.Response{data=model.Review}
// @router /review/{id} [get]
func (c *review) getReview(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := c.mal.GetReview(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get anime review list
// @tags review
// @accept json
// @produce json
// @param page query integer false "Page"
// @success 200 {object} utils.Response{data=[]model.Review}
// @router /reviews/anime [get]
func (c *review) getAnimeReviews(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := c.mal.GetAnimeReviews(page)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get manga review list
// @tags review
// @accept json
// @produce json
// @param page query integer false "Page"
// @success 200 {object} utils.Response{data=[]model.Review}
// @router /reviews/manga [get]
func (c *review) getMangaReviews(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := c.mal.GetMangaReviews(page)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get best review list
// @tags review
// @accept json
// @produce json
// @param page query integer false "Page"
// @success 200 {object} utils.Response{data=[]model.Review}
// @router /reviews/best [get]
func (c *review) getBestReviews(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := c.mal.GetBestReviews(page)
	utils.ResponseWithJSON(w, code, data, err)
}
