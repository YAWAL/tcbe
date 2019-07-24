package routers

import (
	"net/http"

	"github.com/gorilla/mux"
)

const (
	PatientRoute = "/patient/{id}"
	IllnessRoute = "/illness/{name}"
)

func Router(handlers map[string]func(resp http.ResponseWriter, req *http.Request)) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(PatientRoute, handlers[PatientRoute]).Methods(http.MethodGet)
	router.HandleFunc(IllnessRoute, handlers[IllnessRoute]).Methods(http.MethodGet)

	router.HandleFunc("/people", handlers[".GetPatient"]).Methods(http.MethodGet)
	return router
}
