package http

import (
	"fmt"
	"net/http"
	"reblog-server/utils"
	"time"
)

func (c *APIServer) initDemoAPI() {
	c.Demo = c.Root.PathPrefix("/demo").Subrouter()

	c.Demo.HandleFunc("/ping", c.Pong)
	c.Demo.HandleFunc("/context/routine", c.DoSomethingWithAnotherRoutine)
}

// Ping
func (c *APIServer) Pong(w http.ResponseWriter, r *http.Request) {
	test := r.Context().Value("userid")
	w.Write([]byte("pong" + test.(string)))
}

// Routine
func (c *APIServer) DoSomethingWithAnotherRoutine(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	utils.Info("Handler started!")
	defer utils.Info("Handler ended!")

	select {

	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "Hello\n")

	case <-ctx.Done():
		err := ctx.Err()
		utils.Error("server:", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
	}
}
