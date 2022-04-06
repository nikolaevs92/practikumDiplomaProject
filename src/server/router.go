package server

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type DataBase interface {
}

func MakeRouter(db DataBase) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/api/user", func(r chi.Router) {
		r.Post("/register", RegisterPostHandler(db))
		r.Post("/login", LoginPostHandler(db))
		r.Post("/orders", OrdersPostPostHandler(db))
		r.Get("/orders", RegisterGetHandler(db))
		r.Route("/balance", func(r chi.Router) {
			r.Get("/", BalanceGetHandler(db))
			r.Post("/withdraw", WithdrawGetHandler(db))
			r.Get("/withdrawals", WithdrawalsGetHandler(db))
		})
	})

	return r
}

func RegisterPostHandler(db DataBase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}

func LoginPostHandler(db DataBase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}

func OrdersPostPostHandler(db DataBase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}

func RegisterGetHandler(db DataBase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}

func BalanceGetHandler(db DataBase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}

func WithdrawGetHandler(db DataBase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}

func WithdrawalsGetHandler(db DataBase) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
}
