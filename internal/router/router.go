package router

import (
	"GymMembership-api/internal/handlers"
	"GymMembership-api/internal/middleware"

	"github.com/go-chi/chi/v5"
)

func NewRouter(h *handlers.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/register", h.Register)
	r.Post("/login", h.Login)
	r.Get("/classes", h.GetAllClasses)

	r.Group(func(r chi.Router) {
		r.Use(middleware.AuthMiddleware)

		r.Post("/bookings", h.BookClass)
		r.Delete("/bookings/{id}", h.CancelBooking)
		r.Get("/bookings", h.GetMyBookings)
		r.Post("/memberships", h.BuyMembership)
		r.Get("/memberships", h.GetMyMembership)
		r.Post("/classes", h.CreateClass)
	})

	return r
}
