package lab_1

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
}