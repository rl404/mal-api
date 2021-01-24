package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper"
	"github.com/rl404/mal-api/internal/pkg/utils"
)

type article struct {
	mal *malscraper.Malscraper
}

func registerArticle(r chi.Router, mal *malscraper.Malscraper) {
	f := article{mal: mal}
	r.Get("/articles", f.getArticles)
	r.Get("/article/tags", f.getArticleTag)
	r.Get("/article/{id}", f.getArticle)
}

// @summary Get featured article details
// @tags featured article
// @accept json
// @produce json
// @param id path integer true "Featured article ID"
// @success 200 {object} utils.Response{data=model.Article}
// @router /article/{id} [get]
func (c *article) getArticle(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := c.mal.GetArticle(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get featured article tag list
// @tags featured article
// @accept json
// @produce json
// @success 200 {object} utils.Response{data=[]model.ArticleTagItem}
// @router /article/tags [get]
func (c *article) getArticleTag(w http.ResponseWriter, r *http.Request) {
	data, code, err := c.mal.GetArticleTag()
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get featured article list
// @tags featured article
// @accept json
// @produce json
// @param page query integer false "Page"
// @param tag query string false "Article tag"
// @success 200 {object} utils.Response{data=[]model.ArticleItem}
// @router /articles [get]
func (c *article) getArticles(w http.ResponseWriter, r *http.Request) {
	tag := r.URL.Query().Get("tag")
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := c.mal.GetArticles(page, tag)
	utils.ResponseWithJSON(w, code, data, err)
}
