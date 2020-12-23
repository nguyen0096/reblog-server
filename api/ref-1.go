package api

// import (
// 	"log"
// 	"net/http"
// 	"reblog-server/dependency"

// 	"github.com/gorilla/mux"
// )

// type API struct {
// 	Server dependency.IServer
// 	Routes *Routes
// }

// type Routes struct {
// 	Root *mux.Router

// 	ToDos *mux.Router
// 	Dummy *mux.Router
// }

// // Init ...
// func Init(srv dependency.IServer, r *mux.Router) *API {
// 	userMux := http.NewServeMux()
// 	initUser(userMux)
// 	r.Handle("/users", userMux)

// 	groupMux := http.NewServeMux()
// 	initGroup(groupMux)
// 	r.Handle("/groups", groupMux)

// 	r.Handle("/dummy", http.HandlerFunc(dummyHandler))

// 	api := &API{
// 		Server: srv,
// 		Routes: &Routes{},
// 	}

// 	api.Routes.Root = r
// 	api.Routes.ToDos = api.Routes.Root.PathPrefix("/todos").Subrouter()
// 	api.Routes.Dummy = api.Routes.Root.PathPrefix("/dummy").Subrouter()

// 	api.initTodos()
// 	api.initDummy()

// 	return api
// }

// // USERS
// func initUser(r *http.ServeMux) {
// 	r.Handle("/", http.HandlerFunc(getUsers))
// }

// func getUsers(w http.ResponseWriter, r *http.Request) {
// 	log.Println("getUsers")
// 	log.Println(r.Context().Value("username"))
// 	w.Write([]byte("getUsers"))
// }

// // GROUPS
// func initGroup(r *http.ServeMux) {
// 	r.Handle("/", http.HandlerFunc(getGroups))
// }

// func getGroups(w http.ResponseWriter, r *http.Request) {
// 	log.Println("getGroups")
// 	w.Write([]byte("getGroups"))
// }

// // Dummy
// func dummyHandler(w http.ResponseWriter, r *http.Request) {
// 	log.Printf("[dummyHandler] %v", r.URL.RawQuery)
// }

// // Todos
// func (api *API) initTodos() {
// 	api.Routes.ToDos.Handle("", APIHandler(getAllTodos)).Methods("GET")
// }

// func getAllTodos(w http.ResponseWriter, r *http.Request) (int, error) {
// 	log.Println("getAllTodos")
// 	http.Error(w, "Sorry!", http.StatusUnauthorized)
// 	return 0, nil
// }
