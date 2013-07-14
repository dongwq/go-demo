/**
 * Comment: ??
 * User: DongWQ
 * Date: 13-3-12，下午9:23
 * @version 0.1
 */
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func Demo1() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/products/{id}", ProductsHandler)
	router.HandleFunc("/articles", ArticlesHandler)


	router.PathPrefix("/products/")
	http.Handle("/", router)

	http.ListenAndServe(":3452", nil)

//	r.ServeHTTP()
}


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler)
	router.HandleFunc("/{controller}", ProductsHandler)
	router.HandleFunc("/{controller}/{action}", ProductsHandler)
	router.HandleFunc("/{controller}/{action}/{id}", ProductsHandler)

//	router.PathPrefix("/products/")
	http.Handle("/", router)

	http.ListenAndServe(":3452", nil)

//	r.ServeHTTP()
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	 fmt.Fprintf(w, "HomeHandler %d",100)
}
func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ProductsHandler %d",100)
	vars := mux.Vars(r)
	controller := vars["controller"]
	action := vars["action"]
	id := vars["id"]
	fmt.Fprintf(w, "ProductsHandler %d, controller=%s,action=%s, id=%s",100,controller,action, id)
}
func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ArticlesHandler %d",100)
}
func GetAllLists(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home %d",100)
}

