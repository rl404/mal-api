package api

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper"
	"github.com/rl404/mal-api/internal/pkg/utils"
)

type producer struct {
	mal *malscraper.Malscraper
}

func registerProducer(r chi.Router, mal *malscraper.Malscraper) {
	p := producer{mal: mal}
	r.Get("/producers", p.getProducers)
	r.Get("/producer/{id}", p.getProducer)
	r.Get("/magazines", p.getMagazines)
	r.Get("/magazine/{id}", p.getMagazine)
}

// @summary Get anime producer/studio/licensor list
// @tags producer
// @accept json
// @produce json
// @success 200 {object} utils.Response{data=[]model.ItemCount}
// @router /producers [get]
func (c *producer) getProducers(w http.ResponseWriter, r *http.Request) {
	data, code, err := c.mal.GetProducers()
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get anime list from specific producer/studio/licensor
// @tags producer
// @accept json
// @produce json
// @param id path integer true "Producer/studio/licensor ID"
// @param page query integer false "Page"
// @success 200 {object} utils.Response{data=[]model.AnimeItem}
// @router /producer/{id} [get]
func (c *producer) getProducer(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := c.mal.GetProducer(id, page)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get manga magazine/serialization list
// @tags magazine
// @accept json
// @produce json
// @success 200 {object} utils.Response{data=[]model.ItemCount}
// @router /magazines [get]
func (c *producer) getMagazines(w http.ResponseWriter, r *http.Request) {
	data, code, err := c.mal.GetMagazines()
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Get manga list from specific magazine/serialization
// @tags magazine
// @accept json
// @produce json
// @param id path integer true "Magazine/serialization ID"
// @param page query integer false "Page"
// @success 200 {object} utils.Response{data=[]model.MangaItem}
// @router /magazine/{id} [get]
func (c *producer) getMagazine(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := c.mal.GetMagazine(id, page)
	utils.ResponseWithJSON(w, code, data, err)
}
