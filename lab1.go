package lab_1

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

const (
	path = "/time"
	port = ":8795"
)

type handler struct{}

func (h *handler) ServeHTTP(res http.ResponseWrite, req *http.Request) {
	if req.Method != http.MethodGet || req.URL.Path != path {
		http.Error(res, "Not found", http.StatusNotFound)
		return
	}

	currentTime := time.Now().Format(time.RFC3339)
	response := map[string]string{"time": currentTime}

	res.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(res).Encode(response); err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}

	func main() {
		server := &http.Server{
			Addr:    port,
			Handler: &handler{},
		}
	}

}
