package api

import (
	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper"
)

// API is for API routes.
type API struct {
	mal *malscraper.Malscraper
}

// New to create new Api routes.
func New(mal *malscraper.Malscraper) *API {
	return &API{
		mal: mal,
	}
}

// Register to register all Api routes.
func (v *API) Register(r chi.Router) {
	registerAnime(r, v.mal)
	registerManga(r, v.mal)
	registerCharacter(r, v.mal)
	registerPeople(r, v.mal)
	registerProducer(r, v.mal)
	registerGenre(r, v.mal)
	registerReview(r, v.mal)
	registerRecommendation(r, v.mal)
	registerSeason(r, v.mal)
	registerNews(r, v.mal)
	registerArticle(r, v.mal)
	registerClub(r, v.mal)
	registerTop(r, v.mal)
	registerUser(r, v.mal)
	registerSearch(r, v.mal)
}
