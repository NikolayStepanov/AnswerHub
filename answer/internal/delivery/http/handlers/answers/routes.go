package answers

import "net/http"

func (h *Handler) RegisterRoutes() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc(http.MethodGet+" "+"/{id}", h.GetAnswer)
	r.HandleFunc(http.MethodDelete+" "+"/{id}", h.DeleteAnswer)
	return r
}
