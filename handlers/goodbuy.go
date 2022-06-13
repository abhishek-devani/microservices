package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

type GoodBuy struct {
	l *log.Logger
}

func NewGoodbuy(l *log.Logger) *GoodBuy {
	return &GoodBuy{l}
}

func (g *GoodBuy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops..", http.StatusBadRequest)
	}

	rw.Write([]byte("Good Buy " + string(data) + "\n"))
}
