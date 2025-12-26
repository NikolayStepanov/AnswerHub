package answers

import "net/http"

func (h *Handler) RegisterRoutes() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("GET /answers/{id}", h.GetAnswer)
	r.HandleFunc("DELETE /answers/{id}", h.DeleteAnswer)

	return r
}
