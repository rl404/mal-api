package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/mal-api/internal/pkg/utils"
)

type user struct {
	mal *malscraper.Malscraper
}

func registerUser(r chi.Router, mal *malscraper.Malscraper) {
	u := user{mal: mal}
	r.Get("/user/{user}", u.getUser)
	r.Get("/user/{user}/stats", u.getUserStats)
	r.Get("/user/{user}/favorites", u.getUserFavorite)
	r.Get("/user/{user}/friends", u.getUserFriend)
	r.Get("/user/{user}/history", u.getUserHistory)
	r.Get("/user/{user}/history/anime", u.getUserHistoryAnime)
	r.Get("/user/{user}/history/manga", u.getUserHistoryManga)
	r.Get("/user/{user}/reviews", u.getUserReview)
	r.Get("/user/{user}/recommendations", u.getUserRecommendation)
	r.Get("/user/{user}/clubs", u.getUserClub)
	r.Get("/user/{user}/anime", u.getUserAnime)
	r.Get("/user/{user}/manga", u.getUserManga)
}

// @summary Get user details
// @tags user
// @accept json
// @produce json
// @param user path string true "User name"
// @success 200 {object} utils.Response{data=model.User}
// @router /user/{user} [get]
func (c *user) getUser(w http.ResponseWriter, r *http.Request) {
	user := chi.URLParam(r, "user")
	data, code, err := c.mal.GetUser(user)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get user stats
// @tags user
// @accept json
// @produce json
// @param user path string true "User name"
// @success 200 {object} utils.Response{data=model.UserStats}
// @router /user/{user}/stats [get]
func (c *user) getUserStats(w http.ResponseWriter, r *http.Request) {
	user := chi.URLParam(r, "user")
	data, code, err := c.mal.GetUserStats(user)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get user favorite list
// @tags user
// @accept json
// @produce json
// @param user path string true "User name"
// @success 200 {object} utils.Response{data=model.UserFavorite}
// @router /user/{user}/favorites [get]
func (c *user) getUserFavorite(w http.ResponseWriter, r *http.Request) {
	user := chi.URLParam(r, "user")
	data, code, err := c.mal.GetUserFavorite(user)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get user friend list
// @tags user
// @accept json
// @produce json
// @param user path string true "User name"
// @param page query integer false "Page"
// @success 200 {object} utils.Response{data=[]model.UserFriend}
// @router /user/{user}/friends [get]
func (c *user) getUserFriend(w http.ResponseWriter, r *http.Request) {
	user := chi.URLParam(r, "user")
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := c.mal.GetUserFriend(user, page)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get user anime & manga history
// @tags user
// @accept json
// @produce json
// @param user path string true "User name"
// @success 200 {object} utils.Response{data=[]model.UserHistory}
// @router /user/{user}/history [get]
func (c *user) getUserHistory(w http.ResponseWriter, r *http.Request) {
	user := chi.URLParam(r, "user")
	data, code, err := c.mal.GetUserHistory(user)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get user anime history
// @tags user
// @accept json
// @produce json
// @param user path string true "User name"
// @success 200 {object} utils.Response{data=[]model.UserHistory}
// @router /user/{user}/history/anime [get]
func (c *user) getUserHistoryAnime(w http.ResponseWriter, r *http.Request) {
	user := chi.URLParam(r, "user")
	data, code, err := c.mal.GetUserHistory(user, malscraper.AnimeType)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get user manga history
// @tags user
// @accept json
// @produce json
// @param user path string true "User name"
// @success 200 {object} utils.Response{data=[]model.UserHistory}
// @router /user/{user}/history/manga [get]
func (c *user) getUserHistoryManga(w http.ResponseWriter, r *http.Request) {
	user := chi.URLParam(r, "user")
	data, code, err := c.mal.GetUserHistory(user, malscraper.MangaType)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get user review list
// @tags user
// @accept json
// @produce json
// @param user path string true "User name"
// @param page query integer false "Page"
// @success 200 {object} utils.Response{data=[]model.Review}
// @router /user/{user}/reviews [get]
func (c *user) getUserReview(w http.ResponseWriter, r *http.Request) {
	user := chi.URLParam(r, "user")
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := c.mal.GetUserReview(user, page)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get user recommendation list
// @tags user
// @accept json
// @produce json
// @param user path string true "User name"
// @param page query integer false "Page"
// @success 200 {object} utils.Response{data=[]model.Recommendation}
// @router /user/{user}/recommendations [get]
func (c *user) getUserRecommendation(w http.ResponseWriter, r *http.Request) {
	user := chi.URLParam(r, "user")
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := c.mal.GetUserRecommendation(user, page)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get user club list
// @tags user
// @accept json
// @produce json
// @param user path string true "User name"
// @success 200 {object} utils.Response{data=[]model.Item}
// @router /user/{user}/clubs [get]
func (c *user) getUserClub(w http.ResponseWriter, r *http.Request) {
	user := chi.URLParam(r, "user")
	data, code, err := c.mal.GetUserClub(user)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get user anime list
// @tags user
// @accept json
// @produce json
// @param user path string true "User name"
// @param page query integer false "Page"
// @param status query integer false "Progress status (1=watching, 2=completed, 3=on-hold, 4=dropped, 6=planned, 7=all)" enums(0,1,2,3,4,6,7)
// @param order query integer false "Order (1=title, 2=finish date, 3=start date, 4=score, 6=type, 8=rating, 11=priority, 12=progress, 13=storage, 14=airing start, 15=airing end)" enums(0,1,2,3,4,6,8,11,12,13,14,15)
// @param tag query string false "Tag"
// @success 200 {object} utils.Response{data=[]model.UserAnime}
// @router /user/{user}/anime [get]
func (c *user) getUserAnime(w http.ResponseWriter, r *http.Request) {
	var query model.UserListQuery
	query.Username = chi.URLParam(r, "user")
	query.Page, _ = strconv.Atoi(utils.GetQuery(r, "page", "1"))
	query.Status, _ = strconv.Atoi(r.URL.Query().Get("status"))
	query.Order, _ = strconv.Atoi(r.URL.Query().Get("order"))
	query.Tag = r.URL.Query().Get("tag")
	data, code, err := c.mal.GetUserAnimeAdv(query)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get user manga list
// @tags user
// @accept json
// @produce json
// @param user path string true "User name"
// @param page query integer false "Page"
// @param status query integer false "Progress status (1=watching, 2=completed, 3=on-hold, 4=dropped, 6=planned, 7=all)" enums(0,1,2,3,4,6,7)
// @param order query integer false "Order (1=title, 2=finish date, 3=start date, 4=score, 7=priority, 8=chapter, 9=volume, 10=type, 11=publishing start, 12=publishing end)" enums(0,1,2,3,4,7,8,9,10,11,12)
// @param tag query string false "Tag"
// @success 200 {object} utils.Response{data=[]model.UserManga}
// @router /user/{user}/manga [get]
func (c *user) getUserManga(w http.ResponseWriter, r *http.Request) {
	var query model.UserListQuery
	query.Username = chi.URLParam(r, "user")
	query.Page, _ = strconv.Atoi(utils.GetQuery(r, "page", "1"))
	query.Status, _ = strconv.Atoi(r.URL.Query().Get("status"))
	query.Order, _ = strconv.Atoi(r.URL.Query().Get("order"))
	query.Tag = r.URL.Query().Get("tag")
	data, code, err := c.mal.GetUserMangaAdv(query)
	utils.ResponseWithJSON(w, code, data, err)
}
