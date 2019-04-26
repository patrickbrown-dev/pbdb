package http

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/viper"
	"github.com/kineticdial/pbdb/db"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	k := r.URL.Path[len("/get/"):]
	b, err := db.Get(k)
	if err != nil {
		log.Printf("Err: %s", err)
		internalServerError(w, r)
		return
	}

	if len(b) < 1 {
		notFound(w, r)
		return
	}

	fmt.Fprint(w, string(b))
	ok(w, r)
}

func setHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		unprocessableEntity(w, r)
		return
	}

	k := r.URL.Path[len("/set/"):]
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("Err: %s", err)
		internalServerError(w, r)
		return
	}

	err = db.Set(k, b)
	if err != nil {
		log.Printf("Err: %s", err)
		internalServerError(w, r)
		return
	}

	created(w, r)
}

// Serve ...
func Serve() {
	http.HandleFunc("/get/", getHandler)
	http.HandleFunc("/set/", setHandler)
	port := fmt.Sprintf(":%s", viper.GetString("port"))
	log.Printf("Starting pbdb server on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
