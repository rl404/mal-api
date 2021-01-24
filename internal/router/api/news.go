package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper"
	"github.com/rl404/mal-api/internal/pkg/utils"
)

type news struct {
	mal *malscraper.Malscraper
}

func registerNews(r chi.Router, mal *malscraper.Malscraper) {
	n := news{mal: mal}
	r.Get("/news", n.getNewsList)
	r.Get("/news/tags", n.getNewsTag)
	r.Get("/news/{id}", n.getNews)
}

// @summary Get news details
// @tags news
// @accept json
// @produce json
// @param id path integer true "News ID"
// @success 200 {object} utils.Response{data=model.News}
// @router /news/{id} [get]
func (c *news) getNews(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := c.mal.GetNews(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get news tag list
// @tags news
// @accept json
// @produce json
// @success 200 {object} utils.Response{data=model.NewsTag}
// @router /news/tags [get]
func (c *news) getNewsTag(w http.ResponseWriter, r *http.Request) {
	data, code, err := c.mal.GetNewsTag()
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get news list
// @tags news
// @accept json
// @produce json
// @param page query integer false "Page"
// @param tag query string false "News tag"
// @success 200 {object} utils.Response{data=[]model.NewsItem}
// @router /news [get]
func (c *news) getNewsList(w http.ResponseWriter, r *http.Request) {
	tag := r.URL.Query().Get("tag")
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := c.mal.GetNewsList(page, tag)
	utils.ResponseWithJSON(w, code, data, err)
}
