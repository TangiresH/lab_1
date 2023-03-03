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
}