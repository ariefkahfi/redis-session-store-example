package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"gopkg.in/boj/redistore.v1"
)



func main(){
	r := mux.NewRouter()
	rediStore , _ := redistore.NewRediStore(10 , "tcp",":6379","",[]byte("ASD"))

	r.Path("/").Methods("GET").HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w, r , "public/index.html")
	})

	r.Path("/f").Methods("GET").HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w,r , "public/pages/forbidden.html")
	})

	r.Path("/main").Methods("GET").HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		s , _ := rediStore.Get(r,"redis-session")

		_ , ok := s.Values["username"]

		if !ok {
			http.Redirect(w , r , "/f",302)
			return
		}

		http.ServeFile(w , r , "public/pages/main.html")
	})

	r.Path("/page2").Methods("GET").HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		s , _ := rediStore.Get(r,"redis-session")

		_ , ok := s.Values["username"]

		if !ok {
			http.Redirect(w , r , "/f",302)
			return
		}

		http.ServeFile(w,r,"public/pages/page2.html")
	})
	r.Path("/page1").Methods("GET").HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		s , _ := rediStore.Get(r,"redis-session")

		_ , ok := s.Values["username"]

		if !ok {
			http.Redirect(w , r , "/f",302)
			return
		}
		http.ServeFile(w , r , "public/pages/page1.html")			
	})
	
	r.Path("/error").Methods("GET").HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w,r , "public/pages/error-login.html")
	})

	
	r.Path("/login").Methods("POST").HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		username := r.FormValue("username")
		s , _ := rediStore.Get(r , "redis-session")

		s.Values["username"] = username
		s.Save(r,w)

		http.Redirect(w,r,"/main",302)
	})


	http.ListenAndServe(":8989",r)
}