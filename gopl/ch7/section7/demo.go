package section7

import (
	"fmt"
	"log"
	"net/http"
)

func DemoServeMux() {
	mux := http.NewServeMux()
	mux.Handle("/list", http.HandlerFunc(Db.List)) //http.HandlerFunc(db.list)其实是类型转换，将db.list转为 http.HandlerFunc
	mux.Handle("/price", http.HandlerFunc(Db.Price))
	log.Fatal(http.ListenAndServe(":8080", mux))
}

//DemoHttpHandler 手动实现HttpHandler
func DemoHttpHandler() {
	log.Fatal(http.ListenAndServe(":8080", Db))
}

type Dollars float64
type Goods map[string]Dollars

var Db = Goods{"banana": 12, "apple": 10}

func (g Goods) List(w http.ResponseWriter, r *http.Request) {
	for name, price := range g {
		fmt.Fprintf(w, "%s: %.2f\n", name, price)
	}
}
func (g Goods) Price(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("item")
	price, ok := g[key]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No such item:%s", key)
		return
	}
	fmt.Fprintf(w, "The price of %s is $%.2f\n", key, price)
}

func (g Goods) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/list":
		for name, price := range g {
			fmt.Fprintf(w, "%s: %.2f\n", name, price)
		}
		return
	case "/price":
		key := r.URL.Query().Get("item")
		price, ok := g[key]
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "No such item:%s", key)
			return
		}
		fmt.Fprintf(w, "The price of %s is $%.2f\n", key, price)
		return
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Page not found: %s", r.URL.Path)
		return
	}

}
