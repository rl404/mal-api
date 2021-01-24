package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper"
	_utils "github.com/rl404/go-malscraper/pkg/utils"
	"github.com/rl404/mal-api/internal/pkg/utils"
)

type season struct {
	mal *malscraper.Malscraper
}

func registerSeason(r chi.Router, mal *malscraper.Malscraper) {
	s := season{mal: mal}
	r.Get("/season", s.getSeason1)
	r.Get("/season/{season}/{year}", s.getSeason2)
}

// @summary Get seasonal anime list (with query)
// @tags season
// @accept json
// @produce json
// @param season query string false "Season name" enums(winter,spring,summer,fall)
// @param year query integer false "Year"
// @success 200 {object} utils.Response{data=[]model.AnimeItem}
// @router /season [get]
func (c *season) getSeason1(w http.ResponseWriter, r *http.Request) {
	season := utils.GetQuery(r, "season", _utils.GetCurrentSeason())
	year, _ := strconv.Atoi(utils.GetQuery(r, "year", strconv.Itoa(time.Now().Year())))
	data, code, err := c.mal.GetSeason(season, year)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get seasonal anime list (with path)
// @tags season
// @accept json
// @produce json
// @param season path string true "Season name" enums(winter,spring,summer,fall)
// @param year path integer true "Year"
// @success 200 {object} utils.Response{data=[]model.AnimeItem}
// @router /season/{season}/{year} [get]
func (c *season) getSeason2(w http.ResponseWriter, r *http.Request) {
	season := chi.URLParam(r, "season")
	year, _ := strconv.Atoi(chi.URLParam(r, "year"))
	data, code, err := c.mal.GetSeason(season, year)
	utils.ResponseWithJSON(w, code, data, err)
}
