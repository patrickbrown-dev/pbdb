package http

import (
	"log"
	"net/http"
)

func unprocessableEntity(w http.ResponseWriter, r *http.Request) {
	handleStatus(http.StatusUnprocessableEntity, w, r)
}

func created(w http.ResponseWriter, r *http.Request) {
	handleStatus(http.StatusCreated, w, r)
}

func ok(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %d %s", r.Method, http.StatusOK, r.URL.Path)
}

func notFound(w http.ResponseWriter, r *http.Request) {
	handleStatus(http.StatusNotFound, w, r)
}

func internalServerError(w http.ResponseWriter, r *http.Request) {
	handleStatus(http.StatusInternalServerError, w, r)
}

func handleStatus(status int, w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(status)
	log.Printf("%s %d %s", r.Method, status, r.URL.Path)
}
