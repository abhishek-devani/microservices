package handlers

import (
	"log"
	"net/http"

	"github.com/microservices/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	// p.l.Println("product handler")
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	// data, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	// rw.Write(data)

}
