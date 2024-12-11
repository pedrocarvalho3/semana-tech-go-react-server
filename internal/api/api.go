package api

import (
	"net/http"

	"github.com/pedrocarvalho3/semana-tech-go-react-server/internal/store/pgstore"

	"github.com/go-chi/chi/v5"
)

type apiHandler struct {
	q *pgstore.Queries
	r *chi.Mux
}

func (h apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.r.ServeHTTP(w, r)
}

func NewHandler(q *pgstore.Queries) http.Handler {
	a := apiHandler{
		q: q,
		r: chi.NewRouter(),
	}
	a.r.Get("/rooms", a.GetRooms)
	a.r.Post("/rooms", a.InsertRoom)
	a.r.Get("/rooms/{id}", a.GetRoom)
	a.r.Post("/rooms/{id}/messages", a.InsertMessage)
	a.r.Get("/rooms/{id}/messages", a.GetRoomMessages)
	return a
}