package delivery

import (
	api "SiOngkir/delivery/http"
	"html/template"
	"net/http"
)

func SiCepatHandler(w http.ResponseWriter, r *http.Request) {
	data, err := api.SiCepatOngkir("PBL", "PBL10014", "1")
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
	data, err := api.AnterAjaOngkir("35.13.12", "32.15.13")
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
	mux.HandleFunc("/sicepat", SiCepatHandler)
	mux.HandleFunc("/anteraja", AnterAjaHandler)

	server := http.Server{}
	server.Addr = ":8000"
	server.ListenAndServe()
}
