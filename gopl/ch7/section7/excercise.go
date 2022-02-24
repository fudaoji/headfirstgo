package section7

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

//Q711 excercise 7.11
func Q711() {
	http.HandleFunc("/list", Db.ListByHtml)
	http.HandleFunc("/price", Db.Price)
	http.HandleFunc("/add", Db.Add)
	http.HandleFunc("/update", Db.Update)
	http.HandleFunc("/delete", Db.Delete)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func (g Goods) ListByHtml(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./section7/list.html")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	// 利用给定数据渲染模板，并将结果写入w
	type assign struct {
		List map[string]Dollars
	}
	tmpl.Execute(w, assign{List: g})
}

//Delete 删除商品
func (g Goods) Delete(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("item")
	_, ok := g[key]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "No such item:%s", key)
		return
	}
	delete(g, key)
	fmt.Fprintf(w, "delete success!")
}

//update 更新商品
func (g Goods) Update(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("item")
	_, ok := g[key]
	if !ok {
		http.Error(w, fmt.Sprintf("no such item: %s !", key), http.StatusNotFound)
		return
	}
	price, err := strconv.ParseFloat(r.URL.Query().Get("price"), 64)
	if err != nil {
		http.Error(w, "param price is invalid!", http.StatusNotFound)
		return
	}
	if price < 0.0 {
		http.Error(w, "param price could not low than 0!", http.StatusNotFound)
		return
	}
	g[key] = Dollars(price)
	fmt.Fprintf(w, "update success!")
}

//Add 新增商品
func (g Goods) Add(w http.ResponseWriter, r *http.Request) {
	key := r.URL.Query().Get("item")
	_, ok := g[key]
	if ok {
		http.Error(w, fmt.Sprintf("item %s is exist!", key), http.StatusNotFound)
		return
	}
	price, err := strconv.ParseFloat(r.URL.Query().Get("price"), 64)
	if err != nil {
		http.Error(w, "param price is invalid!", http.StatusNotFound)
		return
	}
	if price < 0.0 {
		http.Error(w, "param price could not low than 0!", http.StatusNotFound)
		return
	}
	g[key] = Dollars(price)
	fmt.Fprintf(w, "add success!")
}
