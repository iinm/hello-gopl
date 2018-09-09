package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func main() {
	db := database{items: map[string]dollars{"shoes": 50, "socks": 5}}
	http.HandleFunc("/list", db.List)
	http.HandleFunc("/price", db.Price)
	http.HandleFunc("/create", db.Create)
	http.HandleFunc("/read", db.Price)
	http.HandleFunc("/update", db.Update)
	http.HandleFunc("/delete", db.Delete)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func (db *database) Create(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if len(item) == 0 {
		http.Error(w, "item name is empty\n", http.StatusBadRequest)
		return
	}
	price, err := strconv.ParseFloat(req.URL.Query().Get(`price`), 32)
	if err != nil || price < 0 {
		http.Error(
			w,
			fmt.Sprintf("invalid price: %q\n", req.URL.Query().Get(`price`)),
			http.StatusBadRequest,
		)
		return
	}

	db.Lock()
	defer db.Unlock()
	if _, ok := db.items[item]; ok {
		http.Error(w, fmt.Sprintf("item already exits: %q\n", item), http.StatusBadRequest)
		return
	}
	db.items[item] = dollars(price)
	fmt.Fprintf(w, "%s = %s", item, dollars(price))
}

func (db *database) Update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if len(item) == 0 {
		http.Error(w, "item name is empty\n", http.StatusBadRequest)
		return
	}
	price, err := strconv.ParseFloat(req.URL.Query().Get(`price`), 32)
	if err != nil || price < 0 {
		http.Error(
			w,
			fmt.Sprintf("invalid price: %q\n", req.URL.Query().Get(`price`)),
			http.StatusBadRequest,
		)
		return
	}

	db.Lock()
	defer db.Unlock()
	if _, ok := db.items[item]; !ok {
		http.Error(w, fmt.Sprintf("no such item: %q\n", item), http.StatusBadRequest)
		return
	}
	db.items[item] = dollars(price)
	fmt.Fprintf(w, "%s = %s", item, dollars(price))
}

func (db *database) Delete(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if len(item) == 0 {
		http.Error(w, "item name is empty\n", http.StatusBadRequest)
		return
	}

	db.Lock()
	defer db.Unlock()
	if _, ok := db.items[item]; !ok {
		http.Error(w, fmt.Sprintf("no such item: %q\n", item), http.StatusBadRequest)
		return
	}
	delete(db.items, item)
	fmt.Fprintf(w, "%s is deleted\n", item)
}

func (db *database) List(w http.ResponseWriter, req *http.Request) {
	db.RLock()
	defer db.RUnlock()
	for item, price := range db.items {
		fmt.Fprintf(w, "%s: %s\n", item, price)
	}
}

func (db *database) Price(w http.ResponseWriter, req *http.Request) {
	db.RLock()
	defer db.RUnlock()
	item := req.URL.Query().Get("item")
	if price, ok := db.items[item]; ok {
		fmt.Fprintf(w, "%s\n", price)
	} else {
		http.Error(w, fmt.Sprintf("no such item: %q\n", item), http.StatusNotFound)

	}
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database struct {
	sync.RWMutex
	items map[string]dollars
}
