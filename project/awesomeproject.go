package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	/*Kk
	c1 := make(chan string)

	 go pruebaAsync(c1)
	fmt.Println(<-c1)
	*/
	router := mux.NewRouter()
	router.HandleFunc("/api/user/new", controllers.CreateUser).Methods("POST")

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Print(err)
	}
}

/*
func main() {
	http.HandleFunc("/", indexHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}*/

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
