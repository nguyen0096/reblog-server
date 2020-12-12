package apiv2

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"reblog-server/web"
)

type API struct {
	Routes *Routes
}

type Routes struct {
	Root *mux.Router

	ToDos *mux.Router
}

// Init ...
func Init(r *mux.Router) *API {
	userMux := http.NewServeMux()
	initUser(userMux)
	r.Handle("/users", userMux)

	groupMux := http.NewServeMux()
	initGroup(groupMux)
	r.Handle("/groups", groupMux)

	r.Handle("/dummy", http.HandlerFunc(dummyHandler))

	api := &API{
		Routes: &Routes{},
	}

	api.Routes.Root = r
	api.Routes.ToDos = api.Routes.Root.PathPrefix("/todos").Subrouter()

	api.initTodos()

	return api
}

// USERS
func initUser(r *http.ServeMux) {
	r.Handle("/", http.HandlerFunc(getUsers))
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("getUsers")
	log.Println(r.Context().Value("username"))
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

// Dummy
func dummyHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[dummyHandler] %v", r.URL.RawQuery)
}

// Todos
func (api *API) initTodos() {
	api.Routes.ToDos.Handle("", web.NewHandler(getAllTodos)).Methods("GET")
}

func getAllTodos(c *web.Context, w http.ResponseWriter, r *http.Request) (int, error) {
	log.Println("getAllTodos")
	http.Error(w, "Sorry!", http.StatusUnauthorized)
	return 0, nil
}
