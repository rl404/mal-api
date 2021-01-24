package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper"
	"github.com/rl404/mal-api/internal/pkg/utils"
)

type anime struct {
	mal *malscraper.Malscraper
}

func registerAnime(r chi.Router, mal *malscraper.Malscraper) {
	a := anime{mal: mal}
	r.Get("/anime/{id}", a.getAnime)
	r.Get("/anime/{id}/videos", a.getAnimeVideo)
	r.Get("/anime/{id}/episodes", a.getAnimeEpisode)
	r.Get("/anime/{id}/reviews", a.getAnimeReview)
	r.Get("/anime/{id}/recommendations", a.getAnimeRecommendation)
	r.Get("/anime/{id}/stats", a.getAnimeStats)
	r.Get("/anime/{id}/characters", a.getAnimeCharacter)
	r.Get("/anime/{id}/staff", a.getAnimeStaff)
	r.Get("/anime/{id}/news", a.getAnimeNews)
	r.Get("/anime/{id}/article", a.getAnimeArticle)
	r.Get("/anime/{id}/clubs", a.getAnimeClub)
	r.Get("/anime/{id}/pictures", a.getAnimePicture)
	r.Get("/anime/{id}/more-info", a.getAnimeMoreInfo)
}

// @summary Get anime details
// @tags anime
// @accept json
// @produce json
// @param id path integer true "Anime ID"
// @success 200 {object} utils.Response{data=model.Anime}
// @router /anime/{id} [get]
func (a *anime) getAnime(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := a.mal.GetAnime(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get anime video list
// @tags anime
// @accept json
// @produce json
// @param id path integer true "Anime ID"
// @param page query integer false "Page"
// @success 200 {object} utils.Response{data=model.Video}
// @router /anime/{id}/videos [get]
func (a *anime) getAnimeVideo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := a.mal.GetAnimeVideo(id, page)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get anime episode list
// @tags anime
// @accept json
// @produce json
// @param id path integer true "Anime ID"
// @param page query integer false "Page"
// @success 200 {object} utils.Response{data=[]model.Episode}
// @router /anime/{id}/episodes [get]
func (a *anime) getAnimeEpisode(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := a.mal.GetAnimeEpisode(id, page)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get anime review list
// @tags anime
// @accept json
// @produce json
// @param id path integer true "Anime ID"
// @param page query integer false "Page"
// @success 200 {object} utils.Response{data=[]model.Review}
// @router /anime/{id}/reviews [get]
func (a *anime) getAnimeReview(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := a.mal.GetAnimeReview(id, page)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get anime recommendation list
// @tags anime
// @accept json
// @produce json
// @param id path integer true "Anime ID"
// @success 200 {object} utils.Response{data=[]model.Recommendation}
// @router /anime/{id}/recommendations [get]
func (a *anime) getAnimeRecommendation(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := a.mal.GetAnimeRecommendation(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get anime score stats
// @tags anime
// @accept json
// @produce json
// @param id path integer true "Anime ID"
// @success 200 {object} utils.Response{data=model.Stats}
// @router /anime/{id}/stats [get]
func (a *anime) getAnimeStats(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := a.mal.GetAnimeStats(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get anime character list
// @tags anime
// @accept json
// @produce json
// @param id path integer true "Anime ID"
// @success 200 {object} utils.Response{data=[]model.CharacterItem}
// @router /anime/{id}/characters [get]
func (a *anime) getAnimeCharacter(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := a.mal.GetAnimeCharacter(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get anime staff list
// @tags anime
// @accept json
// @produce json
// @param id path integer true "Anime ID"
// @success 200 {object} utils.Response{data=[]model.Role}
// @router /anime/{id}/staff [get]
func (a *anime) getAnimeStaff(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := a.mal.GetAnimeStaff(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get anime news list
// @tags anime
// @accept json
// @produce json
// @param id path integer true "Anime ID"
// @success 200 {object} utils.Response{data=[]model.NewsItem}
// @router /anime/{id}/news [get]
func (a *anime) getAnimeNews(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := a.mal.GetAnimeNews(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get anime featured article list
// @tags anime
// @accept json
// @produce json
// @param id path integer true "Anime ID"
// @success 200 {object} utils.Response{data=[]model.ArticleItem}
// @router /anime/{id}/article [get]
func (a *anime) getAnimeArticle(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := a.mal.GetAnimeArticle(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get anime club list
// @tags anime
// @accept json
// @produce json
// @param id path integer true "Anime ID"
// @success 200 {object} utils.Response{data=[]model.ClubItem}
// @router /anime/{id}/clubs [get]
func (a *anime) getAnimeClub(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := a.mal.GetAnimeClub(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get anime picture list
// @tags anime
// @accept json
// @produce json
// @param id path integer true "Anime ID"
// @success 200 {object} utils.Response{data=[]string}
// @router /anime/{id}/pictures [get]
func (a *anime) getAnimePicture(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := a.mal.GetAnimePicture(id)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get anime more information
// @tags anime
// @accept json
// @produce json
// @param id path integer true "Anime ID"
// @success 200 {object} utils.Response{data=string}
// @router /anime/{id}/more-info [get]
func (a *anime) getAnimeMoreInfo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	data, code, err := a.mal.GetAnimeMoreInfo(id)
	utils.ResponseWithJSON(w, code, data, err)
}
