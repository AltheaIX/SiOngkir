package delivery

import (
	api "SiOngkir/delivery/http"
	"html/template"
	"net/http"
)

func SicepatHandler(w http.ResponseWriter, r *http.Request) {
	data, err := api.SiCepatOngkir()
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

func AnterAjaHandler(w http.ResponseWriter, r *http.Request) {
	data, err := api.AnterAjaOngkir()
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

func HandlerRun() {
	mux := http.DefaultServeMux
	mux.HandleFunc("/sicepat", SicepatHandler)
	mux.HandleFunc("/anteraja", AnterAjaHandler)

	server := http.Server{}
	server.Addr = ":8000"
	server.ListenAndServe()
}
