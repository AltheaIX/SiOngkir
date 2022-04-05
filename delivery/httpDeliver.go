package delivery

import (
	api "SiOngkir/delivery/http"
	"SiOngkir/delivery/http/middleware"
	"SiOngkir/models"
	"html/template"
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
		tmpl := template.Must(template.New("sicepat.html").ParseFiles("delivery/template/sicepat.html"))
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
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
		tmpl := template.Must(template.New("anteraja.html").ParseFiles("delivery/template/anteraja.html"))
		if err := tmpl.Execute(w, data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func HandlerRun() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/sicepat", SiCepatHandler)
	mux.HandleFunc("/anteraja", AnterAjaHandler)

	var handler http.Handler = mux
	handler = middleware.GetOnly(handler)

	server := http.Server{}
	server.Addr = ":8000"
	server.Handler = handler
	server.ListenAndServe()
}
