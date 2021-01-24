package api

import (
	"errors"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper"
	"github.com/rl404/go-malscraper/model"
	"github.com/rl404/mal-api/internal/pkg/utils"
)

type search struct {
	mal *malscraper.Malscraper
}

func registerSearch(r chi.Router, mal *malscraper.Malscraper) {
	s := search{mal: mal}
	r.Get("/search/anime", s.searchAnime)
	r.Get("/search/manga", s.searchManga)
	r.Get("/search/character", s.searchCharacter)
	r.Get("/search/people", s.searchPeople)
	r.Get("/search/club", s.searchClub)
	r.Get("/search/user", s.searchUser)
}

// @summary Search anime
// @tags search
// @accept json
// @produce json
// @param query query string false "Anime title"
// @param page query integer false "Page"
// @param type query integer false "Anime type (1=TV, 2=OVA, 3=Movie, 4=Special, 5=ONA, 6=Music)" Enums(0,1,2,3,4,5,6)
// @param score query integer false "Score" enums(0,1,2,3,4,5,6,7,8,9,10)
// @param status query integer false "Anime airing status (1=on-going, 2=finished, 3=upcoming)" Enums(0,1,2,3)
// @param producer query integer false "Producer ID"
// @param rating query integer false "Anime rating (1=G, 2=PG, 3=PG13, 4=R17, 5=R, 6=RX)" enums(0,1,2,3,4,5,6)
// @param start query string false "Anime airing start date (yyyy-mm-dd)"
// @param end query string false "Anime airing end date (yyyy-mm-dd)"
// @param xgenre query boolean false "Exlude genre?"
// @param genre query []integer false "Genre ID"
// @param letter query string false "Anime first letter"
// @success 200 {object} utils.Response{data=[]model.AnimeSearch}
// @router /search/anime [get]
func (c *search) searchAnime(w http.ResponseWriter, r *http.Request) {
	var query model.Query
	query.Title = r.URL.Query().Get("query")
	query.Page, _ = strconv.Atoi(utils.GetQuery(r, "page", "1"))
	query.Type, _ = strconv.Atoi(r.URL.Query().Get("type"))
	query.Score, _ = strconv.Atoi(r.URL.Query().Get("score"))
	query.Status, _ = strconv.Atoi(r.URL.Query().Get("status"))
	query.ProducerID, _ = strconv.Atoi(r.URL.Query().Get("producer"))
	query.Rating, _ = strconv.Atoi(r.URL.Query().Get("rating"))
	query.ExcludeGenre, _ = strconv.ParseBool(r.URL.Query().Get("xgenre"))
	query.FirstLetter = r.URL.Query().Get("letter")

	if genre := r.URL.Query().Get("genre"); genre != "" {
		for _, g := range strings.Split(genre, ",") {
			id, _ := strconv.Atoi(g)
			query.GenreIDs = append(query.GenreIDs, id)
		}
	}

	var err error
	if start := r.URL.Query().Get("start"); start != "" {
		query.StartDate, err = time.Parse("2006-01-02", start)
		if err != nil {
			utils.ResponseWithJSON(w, http.StatusBadRequest, nil, errors.New("invalid date format (yyyy-mm-dd)"))
			return
		}
	}

	if end := r.URL.Query().Get("end"); end != "" {
		query.EndDate, err = time.Parse("2006-01-02", end)
		if err != nil {
			utils.ResponseWithJSON(w, http.StatusBadRequest, nil, errors.New("invalid date format (yyyy-mm-dd)"))
			return
		}
	}

	data, code, err := c.mal.AdvSearchAnime(query)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Search manga
// @tags search
// @accept json
// @produce json
// @param query query string false "Manga title"
// @param page query integer false "Page"
// @param type query integer false "Manga type (1=Manga, 2=Light Novel, 3=One-shot, 4=Doujinshi, 5=Manhwa, 6=Manhua, 8=Novel)" Enums(0,1,2,3,4,5,6,8)
// @param score query integer false "Score" enums(0,1,2,3,4,5,6,7,8,9,10)
// @param status query integer false "Manga publishing status (1=on-going, 2=finished, 3=upcoming, 4=hiatus, 5=discontinued)" Enums(0,1,2,3,4,5)
// @param magazine query integer false "Magazine ID"
// @param start query string false "Manga publishing start date (yyyy-mm-dd)"
// @param end query string false "Manga publishing end date (yyyy-mm-dd)"
// @param xgenre query boolean false "Exlude genre?"
// @param genre query []integer false "Genre ID"
// @param letter query string false "Manga first letter"
// @success 200 {object} utils.Response{data=[]model.MangaSearch}
// @router /search/manga [get]
func (c *search) searchManga(w http.ResponseWriter, r *http.Request) {
	var query model.Query
	query.Title = r.URL.Query().Get("query")
	query.Page, _ = strconv.Atoi(utils.GetQuery(r, "page", "1"))
	query.Type, _ = strconv.Atoi(r.URL.Query().Get("type"))
	query.Score, _ = strconv.Atoi(r.URL.Query().Get("score"))
	query.Status, _ = strconv.Atoi(r.URL.Query().Get("status"))
	query.MagazineID, _ = strconv.Atoi(r.URL.Query().Get("magazine"))
	query.ExcludeGenre, _ = strconv.ParseBool(r.URL.Query().Get("xgenre"))
	query.FirstLetter = r.URL.Query().Get("letter")

	if genre := r.URL.Query().Get("genre"); genre != "" {
		for _, g := range strings.Split(genre, ",") {
			id, _ := strconv.Atoi(g)
			query.GenreIDs = append(query.GenreIDs, id)
		}
	}

	var err error
	if start := r.URL.Query().Get("start"); start != "" {
		query.StartDate, err = time.Parse("2006-01-02", start)
		if err != nil {
			utils.ResponseWithJSON(w, http.StatusBadRequest, nil, errors.New("invalid date format (yyyy-mm-dd)"))
			return
		}
	}

	if end := r.URL.Query().Get("end"); end != "" {
		query.EndDate, err = time.Parse("2006-01-02", end)
		if err != nil {
			utils.ResponseWithJSON(w, http.StatusBadRequest, nil, errors.New("invalid date format (yyyy-mm-dd)"))
			return
		}
	}

	data, code, err := c.mal.AdvSearchManga(query)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Search character
// @tags search
// @accept json
// @produce json
// @param query query string false "Character name"
// @param page query integer false "Page"
// @success 200 {object} utils.Response{data=[]model.CharacterSearch}
// @router /search/character [get]
func (c *search) searchCharacter(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := c.mal.SearchCharacter(query, page)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Search people
// @tags search
// @accept json
// @produce json
// @param query query string false "People name"
// @param page query integer false "Page"
// @success 200 {object} utils.Response{data=[]model.PeopleSearch}
// @router /search/people [get]
func (c *search) searchPeople(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	page, _ := strconv.Atoi(utils.GetQuery(r, "page", "1"))
	data, code, err := c.mal.SearchPeople(query, page)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Search club
// @tags search
// @accept json
// @produce json
// @param query query string false "Club name"
// @param page query integer false "Page"
// @param category query integer false "Category (1=anime, 2=convention, 3=actor & artist, 4=character, 5=company, 6=game, 7=japan, 8=city, 9=music, 10=manga, 11=school, 12=other)" enums(0,1,2,3,4,5,6,7,8,9,10,11,12)
// @param sort query integer false "Sort (1=name, 2=comment, 3=post, 5=member)" enums(0,1,2,3,5)
// @success 200 {object} utils.Response{data=[]model.ClubSearch}
// @router /search/club [get]
func (c *search) searchClub(w http.ResponseWriter, r *http.Request) {
	var query model.ClubQuery
	query.Name = r.URL.Query().Get("query")
	query.Page, _ = strconv.Atoi(utils.GetQuery(r, "page", "1"))
	query.Category, _ = strconv.Atoi(r.URL.Query().Get("category"))
	query.Sort, _ = strconv.Atoi(r.URL.Query().Get("sort"))
	data, code, err := c.mal.AdvSearchClub(query)
	utils.ResponseWithJSON(w, code, data, err)
}

// @summary Search user
// @tags search
// @accept json
// @produce json
// @param query query string false "User name"
// @param page query integer false "Page"
// @param location query string false "Location"
// @param minAge query integer false "Minimum age"
// @param maxAge query integer false "Maximum age"
// @param gender query string false "Gender (1=male, 2=female, 3=non-binary)" enums(0,1,2,3)
// @success 200 {object} utils.Response{data=[]model.UserSearch}
// @router /search/user [get]
func (c *search) searchUser(w http.ResponseWriter, r *http.Request) {
	var query model.UserQuery
	query.Username = r.URL.Query().Get("query")
	query.Page, _ = strconv.Atoi(utils.GetQuery(r, "page", "1"))
	query.Location = r.URL.Query().Get("location")
	query.MinAge, _ = strconv.Atoi(r.URL.Query().Get("minAge"))
	query.MaxAge, _ = strconv.Atoi(r.URL.Query().Get("maxAge"))
	query.Gender, _ = strconv.Atoi(r.URL.Query().Get("gender"))
	data, code, err := c.mal.AdvSearchUser(query)
	utils.ResponseWithJSON(w, code, data, err)
}
