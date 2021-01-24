package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper"
	"github.com/rl404/mal-api/internal/pkg/utils"
)

type manga struct {
	mal *malscraper.Malscraper
}

func registerManga(r chi.Router, mal *malscraper.Malscraper) {
	m := manga{mal: mal}
	r.Get("/manga/{id}", m.getManga)
	r.Get("/manga/{id}/reviews", m.getMangaReview)
	r.Get("/manga/{id}/recommendations", m.getMangaRecommendation)
	r.Get("/manga/{id}/stats", m.getMangaStats)
	r.Get("/manga/{id}/characters", m.getMangaCharacter)
	r.Get("/manga/{id}/news", m.getMangaNews)
	r.Get("/manga/{id}/article", m.getMangaArticle)
	r.Get("/manga/{id}/clubs", m.getMangaClub)
	r.Get("/manga/{id}/pictures", m.getMangaPicture)
	r.Get("/manga/{id}/more-info", m.getMangaMoreInfo)
}

// @summary Get manga details
// @tags manga
// @accept json
// @produce json
// @param id path integer true "Manga ID"
// @success 200 {object} utils.Response{data=model.Manga}
// @router /manga/{id} [get]
func (m *manga) getManga(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := m.mal.GetManga(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get manga review list
// @tags manga
// @accept json
// @produce json
// @param id path integer true "Manga ID"
// @param page query integer false "Page"
// @success 200 {object} utils.Response{data=[]model.Review}
// @router /manga/{id}/reviews [get]
func (m *manga) getMangaReview(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := m.mal.GetMangaReview(id, page)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get manga recommendation list
// @tags manga
// @accept json
// @produce json
// @param id path integer true "Manga ID"
// @success 200 {object} utils.Response{data=[]model.Recommendation}
// @router /manga/{id}/recommendations [get]
func (m *manga) getMangaRecommendation(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := m.mal.GetMangaRecommendation(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get manga score stats
// @tags manga
// @accept json
// @produce json
// @param id path integer true "Manga ID"
// @success 200 {object} utils.Response{data=model.Stats}
// @router /manga/{id}/stats [get]
func (m *manga) getMangaStats(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := m.mal.GetMangaStats(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get manga character list
// @tags manga
// @accept json
// @produce json
// @param id path integer true "Manga ID"
// @success 200 {object} utils.Response{data=[]model.Role}
// @router /manga/{id}/characters [get]
func (m *manga) getMangaCharacter(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := m.mal.GetMangaCharacter(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get manga news list
// @tags manga
// @accept json
// @produce json
// @param id path integer true "Manga ID"
// @success 200 {object} utils.Response{data=[]model.NewsItem}
// @router /manga/{id}/news [get]
func (m *manga) getMangaNews(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := m.mal.GetMangaNews(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get manga featured article list
// @tags manga
// @accept json
// @produce json
// @param id path integer true "Manga ID"
// @success 200 {object} utils.Response{data=[]model.ArticleItem}
// @router /manga/{id}/article [get]
func (m *manga) getMangaArticle(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := m.mal.GetMangaArticle(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get manga club list
// @tags manga
// @accept json
// @produce json
// @param id path integer true "Manga ID"
// @success 200 {object} utils.Response{data=[]model.ClubItem}
// @router /manga/{id}/clubs [get]
func (m *manga) getMangaClub(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := m.mal.GetMangaClub(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get manga picture list
// @tags manga
// @accept json
// @produce json
// @param id path integer true "Manga ID"
// @success 200 {object} utils.Response{data=[]string}
// @router /manga/{id}/pictures [get]
func (m *manga) getMangaPicture(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := m.mal.GetMangaPicture(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get manga more information
// @tags manga
// @accept json
// @produce json
// @param id path integer true "Manga ID"
// @success 200 {object} utils.Response{data=string}
// @router /manga/{id}/more-info [get]
func (m *manga) getMangaMoreInfo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := m.mal.GetMangaMoreInfo(id)
	utils.ResponseWithJSON(w, code, data, err)
}
