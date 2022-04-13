package delivery

import (
	api "SiOngkir/delivery/http"
	"SiOngkir/models"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func SiCepatHandler(w http.ResponseWriter, r *http.Request) {
	origin := r.URL.Query().Get("origin")
	dst := r.URL.Query().Get("destination")
	weight := r.URL.Query().Get("weight")
	if origin != "" && dst != "" {
		if weight == "" {
			weight = "1"
		}

		request := models.RequestSiCepat{models.Request{origin, dst}, weight}
		data, err := api.SiCepatOngkir(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(resp)
	}
}

func AnterAjaHandler(w http.ResponseWriter, r *http.Request) {
	origin := r.URL.Query().Get("origin")
	dst := r.URL.Query().Get("destination")
	if origin != "" && dst != "" {
		request := models.Request{origin, dst}
		data, err := api.AnterAjaOngkir(&request)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(resp)
	}
}

func SetContentType(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		middleware.SetHeader("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func HandlerRun() {
	r := chi.NewRouter()
	r.Route("/api", func(r chi.Router) {
		r.Use(SetContentType)
		r.Get("/sicepat", SiCepatHandler)
		r.Get("/anteraja", AnterAjaHandler)
	})
	http.ListenAndServe(":5000", r)
}
