package apiv2

import (
	"log"
	"net/http"
)

// Init ...
func Init(r *http.ServeMux) {
	userMux := http.NewServeMux()
	initUser(userMux)
	r.Handle("/users", userMux)

	groupMux := http.NewServeMux()
	initGroup(groupMux)
	r.Handle("/groups", groupMux)

	r.Handle("/dummy", http.HandlerFunc(dummyHandler))
}

// USERS
func initUser(r *http.ServeMux) {
	r.Handle("/", http.HandlerFunc(getUsers))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("getUsers")
	w.Write([]byte("getUsers"))
}

// GROUPS
func initGroup(r *http.ServeMux) {
	r.Handle("/", http.HandlerFunc(getGroups))
}

func getGroups(w http.ResponseWriter, r *http.Request) {
	log.Println("getGroups")
	w.Write([]byte("getGroups"))
}

func dummyHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[dummyHandler] %v", r.URL.RawQuery)
}
