package models

type Request struct {
	Origin      string
	Destination string
}

type RequestSiCepat struct {
	Request Request
	Weight  string
}
