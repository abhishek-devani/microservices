package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/microservices/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// 	if r.Method == http.MethodGet {
// 		p.getProducts(rw, r)
// 		return
// 	}

// 	if r.Method == http.MethodPost {
// 		p.addProduct(rw, r)
// 		return
// 	}

// 	if r.Method == http.MethodPut {

// 		p.l.Println("PUT", r.URL.Path)

// 		reg := regexp.MustCompile(`/([0-9]+)`)
// 		g := reg.FindAllStringSubmatch(r.URL.Path, -1)

// 		if len(g) != 1 {
// 			http.Error(rw, "invalid URl", http.StatusNotFound)
// 			return
// 		}

// 		if len(g[0]) != 2 {
// 			http.Error(rw, "invalid URl", http.StatusNotFound)
// 			return
// 		}

// 		idStr := g[0][1]
// 		id, _ := strconv.Atoi(idStr)

// 		// p.l.Println("id:", id)

// 		p.updateProduct(id, rw, r)

// 	}

// 	// catch all
// 	rw.WriteHeader(http.StatusMethodNotAllowed)

// }

func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("updateProduct Handler")

	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err := data.UpdateProduct(&prod, id)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
	}

}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("addProduct Handler")

	prod := r.Context().Value(KeyProduct{}).(data.Product)

	// p.l.Printf("Prod: %#v", prod)
	data.AddProduct(&prod)

}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {

	p.l.Println("getProduct Handler")

	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	// data, err := json.Marshal(lp)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
	// rw.Write(data)

}

type KeyProduct struct{}

func (p Products) MiddlewareProductValidation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		prod := data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[Error] deserializing product", err)
			http.Error(rw, "Error reading product", http.StatusBadRequest)
			return
		}

		// validate the product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[Error] validating product", err)
			http.Error(rw, fmt.Sprintf("Error validating product: %s", err), http.StatusBadRequest)
		}

		// add the product to the context
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		// call the next handler which can be other middlerware in the chain, or the final handler
		next.ServeHTTP(rw, req)

	})
}
