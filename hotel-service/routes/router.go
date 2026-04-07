package routes

import "github.com/go-chi/chi/v5"

type Router interface {
	Register(r chi.Router)
}

func SetupRouter(HotelRouter Router) *chi.Mux {
	r := chi.NewRouter()
	r.Route("/hotel", func(r chi.Router) {
		HotelRouter.Register(r)
	})
	return r
}
